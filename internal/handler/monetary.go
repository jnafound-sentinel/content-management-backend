package handler

import (
	"jna-manager/internal/config"
	payments "jna-manager/internal/domain/models/payments"
	"jna-manager/internal/domain/schemas"
	"jna-manager/internal/middleware"
	"jna-manager/internal/service"
	"jna-manager/pkg/common/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DonationHandler struct {
	service *service.DonationService
	cfg     *config.Config
}

type PaymentHandler struct {
	service         *service.PaymentService
	paystackService *service.PaystackService
	donation        *service.DonationService
	cfg             *config.Config
}

func NewPaymentHandler(service *service.PaymentService, paystackService *service.PaystackService, donationService *service.DonationService, cfg *config.Config) *PaymentHandler {
	return &PaymentHandler{
		service:         service,
		paystackService: paystackService,
		donation:        donationService,
		cfg:             cfg,
	}
}

func NewDonationHandler(service *service.DonationService, cfg *config.Config) *DonationHandler {
	return &DonationHandler{
		service: service,
		cfg:     cfg,
	}
}

func (h *DonationHandler) RegisterRoutes(r *gin.RouterGroup) {
	donations := r.Group("/donations")

	admins := donations.Group("")
	admins.Use(middleware.JWTMiddleware(h.cfg))

	admins = admins.Group("", middleware.RequireRoles("admin"))
	{
		admins.DELETE("/:id", h.DeleteDonation)
	}

	{
		donations.POST("/", h.CreateDonation)
		donations.GET("/", h.ListDonations)
		donations.GET("/:id", h.GetDonation)
	}
}

func (h *PaymentHandler) RegisterRoutes(r *gin.RouterGroup) {
	payment := r.Group("/payments")

	admins := payment.Group("")
	admins.Use(middleware.JWTMiddleware(h.cfg))

	admins = admins.Group("", middleware.RequireRoles("admin"))
	{
		admins.GET("/", h.ListPayments)
		admins.GET("/:id", h.GetPayment)
	}

	{
		payment.POST("/", h.CreatePayment)
		payment.POST("/verify", h.VerifyPayment)
	}
}

func (h *DonationHandler) CreateDonation(c *gin.Context) {
	var req schemas.CreateDonationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	donation := &payments.Donation{
		TagName:     req.TagName,
		Purpose:     req.Purpose,
		Amount:      req.Amount,
		Description: req.Description,
	}

	if err := h.service.CreateDonation(donation); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(donation))
}

func (h *DonationHandler) ListDonations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	donations, total, err := h.service.ListDonations(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(
		http.StatusOK,
		response.NewSuccessResponse(gin.H{
			"donations": donations,
			"meta": gin.H{
				"total":     total,
				"page":      page,
				"page_size": pageSize,
			},
		}),
	)
}

func (h *DonationHandler) GetDonation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	donation, err := h.service.GetDonation(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewFailureResponse("Donation not found"))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(donation))
}

func (h *DonationHandler) DeleteDonation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is invalid"))
		return
	}

	if err := h.service.DeleteDonation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{"message": "Donation deleted successfully"}))
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req schemas.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	donation, err := h.donation.GetDonation(req.DonationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	payment := &payments.Payment{
		DonationID: uuid.MustParse(req.DonationID),
		Amount:     req.Amount,
		Reference:  req.Reference,
		Status:     "pending",
		Donation:   *donation,
	}

	ref, authURL, err := h.paystackService.InitiateTransaction(req.Email, req.CallbackUrl, float64(req.Amount))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewFailureResponse(err.Error()))
		return
	}

	if err := h.service.CreatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewFailureResponse(err.Error()))
		return
	}

	responseData := schemas.PaymentResponse{
		Reference:  ref,
		AuthURL:    authURL,
		DonationID: req.DonationID,
		Amount:     req.Amount,
		Status:     "processing",
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(responseData))
}

func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	var req schemas.VerifyPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	payment, err := h.service.GetPayment(req.PaymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewServerResponse(err.Error()))
		return
	}

	transaction, err := h.paystackService.VerifyTransaction(req.Reference)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewServerResponse("Failed to verify payment transaction!"))
		return
	}

	if !transaction.Status {
		payment.Status = "failed"
	} else if transaction.Data.Status == "success" {
		payment.Status = "verified"
	}

	if err := h.service.UpdatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse("Failed to update payment transaction status!"))
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(payment))
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	payment, err := h.service.GetPayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewFailureResponse("Payment not found"))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(payment))
}

func (h *PaymentHandler) ListPayments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	payments, total, err := h.service.ListPayments(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(
		http.StatusOK,
		response.NewSuccessResponse(gin.H{
			"payments": payments,
			"meta": gin.H{
				"total":     total,
				"page":      page,
				"page_size": pageSize,
			},
		}),
	)
}
