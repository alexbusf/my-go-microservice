package service

import (
	"context"

	"my-go-app/internal/repository"
)

// ItemService - наш слой бизнес-логики
type ItemService struct {
	repo repository.Repository
}

// Конструктор
func NewItemService(r repository.Repository) *ItemService {
	return &ItemService{repo: r}
}

// GetAvailableItems - метод, содержащий правило бизнеса
func (s *ItemService) GetAvailableItems(ctx context.Context) ([]repository.Item, error) {
	// 1. Сначала просто получаем всё из базы (через репозиторий)
	items, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// 2. Применяем ПРАВИЛО БИЗНЕСА
	// Фильтруем список: оставляем только те товары, где ID нечетный
	var availableItems []repository.Item
	for _, item := range items {
		if item.ID%2 != 0 { // Если остаток от деления не 0 (нечетное)
			availableItems = append(availableItems, item)
		}
	}

	// 3. Возвращаем очищенный список
	return availableItems, nil
}