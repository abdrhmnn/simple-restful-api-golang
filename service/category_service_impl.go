package service

import (
	"context"
	"database/sql"
	"simple_restful_api_golang/helper"
	"simple_restful_api_golang/model/api"
	"simple_restful_api_golang/model/entity"
	"simple_restful_api_golang/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryServices {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, request api.CategoryCreateRequest) api.CategoryCreateResponse {
	errValidasi := service.Validate.Struct(request)
	if errValidasi != nil {
		panic(errValidasi)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolbak := tx.Rollback()
			if errRolbak != nil {
				panic(errRolbak)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	category := entity.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, request api.CategoryUpdateRequest) api.CategoryCreateResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolbak := tx.Rollback()
			if errRolbak != nil {
				panic(errRolbak)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	category, errGetId := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if errGetId != nil {
		panic(errGetId)
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolbak := tx.Rollback()
			if errRolbak != nil {
				panic(errRolbak)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(err)
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindByIdCategory(ctx context.Context, categoryId int) api.CategoryCreateResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolbak := tx.Rollback()
			if errRolbak != nil {
				panic(errRolbak)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(err)
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAllCategory(ctx context.Context) []api.CategoryCreateResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolbak := tx.Rollback()
			if errRolbak != nil {
				panic(errRolbak)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	categories := service.CategoryRepository.FindAll(ctx, tx)

	var categoryResponse []api.CategoryCreateResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, helper.ToCategoryResponse(category))
	}

	return categoryResponse
}
