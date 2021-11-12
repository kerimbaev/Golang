package store

import (
	"context"
	"hw8/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Categories() CategoriesRepository
	Reciepts() RecieptsRepository
}

type CategoriesRepository interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
}

type RecieptsRepository interface {
	Create(ctx context.Context, good *models.Reciept) error
	All(ctx context.Context) ([]*models.Reciept, error)
	ByID(ctx context.Context, id int) (*models.Reciept, error)
	Update(ctx context.Context, good *models.Reciept) error
	Delete(ctx context.Context, id int) error
}
