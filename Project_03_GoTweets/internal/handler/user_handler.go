package handler

import (
	"go_tweets/internal/dto"
	"go_tweets/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	api         *gin.Engine
	userService service.UserService
}

func NewHandler(api *gin.Engine, userService service.UserService) *Handler {
	return &Handler{
		api:         api,
		userService: userService,
	}
}

func (h *Handler) RouterList() {
	authorized := h.api.Group("/authorized")

	authorized.POST("/register", h.Register)
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

	register, i, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(i, dto.RegisterResponse{
		UserId: register,
	})
}
