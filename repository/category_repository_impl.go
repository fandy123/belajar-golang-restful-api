package repository

import (
	"context"
	"database/sql"
	"errors"
	"fandy123/belajar-golang-restful-api/helper"
	"fandy123/belajar-golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Catagory) domain.Catagory {
	SQL := "insert into category(name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Catagory) domain.Catagory {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Catagory) {
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Catagory, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Catagory{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Catagory {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Catagory
	for rows.Next() {
		category := domain.Catagory{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
