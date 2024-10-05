package service

import (
	"context"
	"simple_restful_api_golang/model/api"
)

type CategoryServices interface {
	// biasanya jumlah service nya ituh mengikuti jumlah api nya
	CreateCategory(ctx context.Context, request api.CategoryCreateRequest) api.CategoryCreateResponse
	UpdateCategory(ctx context.Context, request api.CategoryUpdateRequest) api.CategoryCreateResponse
	DeleteCategory(ctx context.Context, categoryId int)
	FindByIdCategory(ctx context.Context, categoryId int) api.CategoryCreateResponse
	FindAllCategory(ctx context.Context) []api.CategoryCreateResponse
}
