package handler

import (
	"go_tweets/internal/dto"
	"go_tweets/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validator   *validator.Validate
	userService service.UserService
}

func NewHandler(api *gin.Engine, validator *validator.Validate, userService service.UserService) *Handler {
	return &Handler{
		api:         api,
		validator:   validator,
		userService: userService,
	}
}

func (h *Handler) RouterList() {
	authorized := h.api.Group("/authorized")

	authorized.POST("/register", h.Register)
	authorized.POST("/login", h.Login)
	authorized.POST("/refresh", h.RefreshToken)
}

func (h *Handler) Register(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
	)

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	register, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, dto.RegisterResponse{
		UserId: register,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.LoginRequest
	)

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, refreshToken, statusCode, err := h.userService.Login(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RefreshTokenRequest
	)

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, statusCode, err := h.userService.RefreshToken(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
