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

type BlogHandler struct {
	service *service.BlogService
	cfg     *config.Config
}

func NewBlogHandler(service *service.BlogService, cfg *config.Config) *BlogHandler {
	return &BlogHandler{service: service, cfg: cfg}
}

func (h *BlogHandler) RegisterRoutes(r *gin.RouterGroup) {
	blogs := r.Group("/blogs")

	admins := blogs.Group("")
	admins.Use(middleware.JWTMiddleware(h.cfg))

	admins = admins.Group("", middleware.RequireRoles("admin"))
	{
		admins.POST("/", h.CreateBlog)
		admins.PUT("/:id", h.UpdateBlog)
		admins.DELETE("/:id", h.DeleteBlog)
	}

	{
		blogs.GET("/", h.ListBlogs)
		blogs.GET("/:id", h.GetBlog)
	}
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var req schemas.CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	blog := &org.Blog{
		Title:   req.Title,
		Content: req.Content,
		Excerpt: req.Excerpt,
		Date:    req.Date,
		Tag:     req.Tag,
		Image:   req.Image,
	}

	if err := h.service.CreateBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(blog))
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	blog, err := h.service.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewFailureResponse("Blog not found"))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(blog))
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	var req schemas.UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse(err.Error()))
		return
	}

	blog, err := h.service.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewFailureResponse("Blog not found"))
		return
	}

	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Content != "" {
		blog.Content = req.Content
	}
	if req.Excerpt != "" {
		blog.Excerpt = req.Excerpt
	}
	if req.Tag != "" {
		blog.Tag = req.Tag
	}
	if req.Image != "" {
		blog.Image = req.Image
	}
	if !req.Date.IsZero() {
		blog.Date = req.Date
	}

	if err := h.service.UpdateBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(blog))
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewFailureResponse("ID is required"))
		return
	}

	if err := h.service.DeleteBlog(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{"message": "Blog deleted successfully"}))
}

func (h *BlogHandler) ListBlogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	blogs, total, err := h.service.ListBlogs(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewServerResponse(err.Error()))
		return
	}

	c.JSON(
		http.StatusOK,
		response.NewSuccessResponse(gin.H{
			"blogs": blogs,
			"meta": gin.H{
				"total":     total,
				"page":      page,
				"page_size": pageSize,
			},
		}),
	)
}
