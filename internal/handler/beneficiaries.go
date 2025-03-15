package handler

import (
	"net/http"
	"strconv"

	"jna-manager/internal/config"
	org "jna-manager/internal/domain/models/org"
	"jna-manager/internal/domain/schemas"
	"jna-manager/internal/middleware"
	"jna-manager/internal/service"
	"jna-manager/pkg/common/response"

	"github.com/gin-gonic/gin"
)

type BeneficiaryHandler struct {
	service *service.BeneficiaryService
	cfg     *config.Config
}

func NewBeneficiaryHandler(service *service.BeneficiaryService, cfg *config.Config) *BeneficiaryHandler {
	return &BeneficiaryHandler{service: service, cfg: cfg}
}

func (h *BeneficiaryHandler) RegisterRoutes(r *gin.RouterGroup) {
	beneficiarys := r.Group("/beneficiaries")

	admins := beneficiarys.Group("")
	admins.Use(middleware.JWTMiddleware(h.cfg))

	admins = admins.Group("", middleware.RequireRoles("admin"))
	{
		admins.POST("/", h.CreateBeneficiary)
		admins.PUT("/:id", h.UpdateBeneficiary)
		admins.DELETE("/:id", h.DeleteBeneficiary)
	}

	{
		beneficiarys.GET("/", h.ListBeneficiarys)
		beneficiarys.GET("/:id", h.GetBeneficiary)
	}
}

func (h *BeneficiaryHandler) CreateBeneficiary(c *gin.Context) {
	var req schemas.CreateBeneficiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	beneficiary := &org.Beneficiary{
		RecipientName: req.RecipientName,
		Image:         req.Image,
		ProgramType:   req.ProgramType,
		ShortBio:      req.ShortBio,
		FullBio:       req.FullBio,
		Quote:         req.Quote,
		Featured:      req.Featured,
	}

	if err := h.service.CreateBeneficiary(beneficiary); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(beneficiary))
}

func (h *BeneficiaryHandler) GetBeneficiary(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	beneficiary, err := h.service.GetBeneficiary(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewFailureResponse("Beneficiary not found"))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(beneficiary))
}

func (h *BeneficiaryHandler) UpdateBeneficiary(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	var req schemas.UpdateBeneficiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	beneficiary, err := h.service.GetBeneficiary(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewFailureResponse("Beneficiary not found"))
		return
	}

	if req.RecipientName != "" {
		beneficiary.RecipientName = req.RecipientName
	}
	if req.ProgramType != "" {
		beneficiary.ProgramType = req.ProgramType
	}
	if req.ShortBio != "" {
		beneficiary.ShortBio = req.ShortBio
	}
	if req.FullBio != "" {
		beneficiary.FullBio = req.FullBio
	}
	if req.Quote != "" {
		beneficiary.Quote = req.Quote
	}
	if req.Featured != beneficiary.Featured {
		beneficiary.Featured = req.Featured
	}
	if req.Image != "" {
		beneficiary.Image = req.Image
	}

	if err := h.service.UpdateBeneficiary(beneficiary); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(beneficiary))
}

func (h *BeneficiaryHandler) DeleteBeneficiary(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	if err := h.service.DeleteBeneficiary(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{"message": "Beneficiary deleted successfully"}))
}

func (h *BeneficiaryHandler) ListBeneficiarys(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	beneficiaries, total, err := h.service.ListBeneficiaries(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(
		http.StatusOK,
		response.NewSuccessResponse(gin.H{
			"beneficiaries": beneficiaries,
			"meta": gin.H{
				"total":     total,
				"page":      page,
				"page_size": pageSize,
			},
		}),
	)
}
