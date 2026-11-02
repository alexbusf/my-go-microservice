package main

import (
	"log"
	"fmt"
	"context" // <-- Импорт контекста

	"my-go-app/internal/config"
	"my-go-app/internal/handlers"
	"my-go-app/internal/repository"
	"my-go-app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.InitConfig()

	// Собираем строку из данных конфига
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbConfig.User,
		cfg.DbConfig.Password,
		cfg.DbConfig.Host,
		cfg.DbConfig.Port,
		cfg.DbConfig.Name,
	)
	
	// --- ИСПРАВЛЕНИЕ ЗДЕСЬ ---
	// Используем context.Background() вместо несуществующего os.Context
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer pool.Close()

	// Передаем пул в репозиторий
	repo := repository.NewPostgresRepo(pool)

	// Передаем репозиторий в обработчик
	itemService := service.NewItemService(repo)

	// 3. Создаем обработчик и отдаем ему сервис 
	handler := handlers.NewItemHandler(itemService)

	// Запуск сервера
	r := gin.Default()
	r.GET("/items", handler.ListItems)

	log.Printf("Server starting on :%s", cfg.Port)
	r.Run(":" + cfg.Port)
}