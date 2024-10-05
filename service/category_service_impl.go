package service

import (
	"context"
	"database/sql"
	"simple_restful_api_golang/model/api"
	"simple_restful_api_golang/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, request api.CategoryCreateRequest) (_ api.CategoryCreateResponse) {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, request api.CategoryUpdateRequest) (_ api.CategoryCreateResponse) {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) FindByIdCategory(ctx context.Context, categoryId int) (_ api.CategoryCreateResponse) {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) FindAllCategory(ctx context.Context) (_ []api.CategoryCreateResponse) {
	panic("not implemented") // TODO: Implement
}
