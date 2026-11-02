package handlers

import (
	"net/http"

	"my-go-app/internal/service" // <-- Новый импорт

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemService *service.ItemService // <-- Зависим от сервиса, а не от репозитория
}

// Конструктор тоже меняется
func NewItemHandler(s *service.ItemService) *ItemHandler {
	return &ItemHandler{itemService: s}
}

func (h *ItemHandler) ListItems(c *gin.Context) {
	// Handler просто делегирует работу
	items, err := h.itemService.GetAvailableItems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}