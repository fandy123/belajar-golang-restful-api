package repository

import (
	"context"
	"database/sql"
	"fandy123/belajar-golang-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Catagory) domain.Catagory
	Update(ctx context.Context, tx *sql.Tx, category domain.Catagory) domain.Catagory
	Delete(ctx context.Context, tx *sql.Tx, category domain.Catagory)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Catagory, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Catagory
}
