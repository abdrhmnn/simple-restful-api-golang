package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple_restful_api_golang/model/entity"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		panic(err)
	}

	Id, _ := result.LastInsertId()

	category.Id = int(Id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	result, err := tx.QueryContext(ctx, SQL, category.Name, category.Id)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	if err != nil {
		panic(err)
	}
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	result, err := tx.QueryContext(ctx, SQL, categoryId)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	category := entity.Category{}

	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	SQL := "SELECT id, name FROM category"
	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	var categories []entity.Category
	for result.Next() {
		category := entity.Category{}
		err := result.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}
