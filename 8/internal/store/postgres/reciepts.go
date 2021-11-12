package postgres

import (
	"context"
	"hw8/internal/models"
	"hw8/internal/store"

	"github.com/jmoiron/sqlx"
)

func (db *DB) Reciepts() store.RecieptsRepository {
	if db.Reciepts() == nil {
		db.reciepts = NewRecieptsRepository(db.conn)
	}
	return db.reciepts
}

type RecieptsRepository struct {
	conn *sqlx.DB
}

func NewRecieptsRepository(conn *sqlx.DB) store.RecieptsRepository {
	return &RecieptsRepository{conn: conn}
}

func (c RecieptsRepository) Create(ctx context.Context, reciept *models.Reciept) error {
	_, err := c.conn.Exec("INSERT INTO reciepts(name, category_id, ingridient, description) VALUES ($1, $2, $3, $4)", reciept.Name, reciept.CategoryID, reciept.Ingridient, reciept.Description)
	if err != nil {
		return err
	}

	return nil
}

func (c RecieptsRepository) All(ctx context.Context) ([]*models.Reciept, error) {
	reciepts := make([]*models.Reciept, 0)
	if err := c.conn.Select(&reciepts, "SELECT * FROM reciepts"); err != nil {
		return nil, err
	}

	return reciepts, nil
}

func (c RecieptsRepository) ByID(ctx context.Context, id int) (*models.Reciept, error) {
	reciept := new(models.Reciept)
	if err := c.conn.Get(reciept, "SELECT * FROM categories WHERE id=$1", id); err != nil {
		return nil, err
	}

	return reciept, nil
}

func (c RecieptsRepository) Update(ctx context.Context, reciept *models.Reciept) error {
	_, err := c.conn.Exec("UPDATE jobs SET name = $1, ingridient = $2, description = $3  WHERE id = $4", reciept.Name, reciept.Ingridient, reciept.Description, reciept.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c RecieptsRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM reciepts WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
