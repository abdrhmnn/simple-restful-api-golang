package controller

import (
	"encoding/json"
	"net/http"
	"simple_restful_api_golang/model/api"
	"simple_restful_api_golang/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryServices
}

func NewCategoryController(categoryService service.CategoryServices) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := api.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	if err != nil {
		panic(err)
	}

	categoryResponse := controller.CategoryService.CreateCategory(request.Context(), categoryCreateRequest)
	standartResponse := api.StandartResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(standartResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := api.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	if err != nil {
		panic(err)
	}

	categoryId := params.ByName("categoryId")
	id, _ := strconv.Atoi(categoryId)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.UpdateCategory(request.Context(), categoryUpdateRequest)
	standartResponse := api.StandartResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(standartResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		panic(err)
	}

	controller.CategoryService.DeleteCategory(request.Context(), id)
	standartResponse := api.StandartResponse{
		Code:   200,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(standartResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		panic(err)
	}

	categoryResponse := controller.CategoryService.FindByIdCategory(request.Context(), id)
	standartResponse := api.StandartResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(standartResponse)
	if errEncode != nil {
		panic(errEncode)
	}
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAllCategory(request.Context())
	standartResponse := api.StandartResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(standartResponse)
	if err != nil {
		panic(err)
	}
}
