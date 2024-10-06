package exepciton

import (
	"encoding/json"
	"net/http"
	"simple_restful_api_golang/model/api"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validateError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validateError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exeception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		standartResponse := api.StandartResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exeception.Error(),
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(standartResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exeception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		standartResponse := api.StandartResponse{
			Code:   http.StatusNotFound,
			Status: "Internal Server Error",
			Data:   exeception.Error,
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(standartResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	standartResponse := api.StandartResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(standartResponse)
	if err != nil {
		panic(err)
	}
}
