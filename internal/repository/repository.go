package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Item - структура данных
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// --- НУЖНОЕ ДОБАВЛЕНИЕ ---
// Repository - это интерфейс (контракт).
// Он говорит: "Любой, кто хочет быть репозиторием, должен уметь делать GetAll".
// Мы будем передавать этот интерфейс в хендлер, чтобы не зависеть от конкретной базы данных.
type Repository interface {
	GetAll(ctx context.Context) ([]Item, error)
}
// --- КОНЕЦ ДОБАВЛЕНИЯ ---

// PostgresRepo - реальная реализация для PostgreSQL
type PostgresRepo struct {
	DB *pgxpool.Pool
}

// NewPostgresRepo - конструктор
func NewPostgresRepo(pool *pgxpool.Pool) Repository {
	return &PostgresRepo{DB: pool}
}

// GetAll - реализация метода
func (r *PostgresRepo) GetAll(ctx context.Context) ([]Item, error) {
	query := `SELECT id, name FROM items ORDER BY id ASC`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying items: %w", err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, fmt.Errorf("error scanning item: %w", err)
		}
		items = append(items, i)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return items, nil
}