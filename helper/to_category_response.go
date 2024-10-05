package helper

import (
	"simple_restful_api_golang/model/api"
	"simple_restful_api_golang/model/entity"
)

func ToCategoryResponse(category entity.Category) api.CategoryCreateResponse {
	return api.CategoryCreateResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
