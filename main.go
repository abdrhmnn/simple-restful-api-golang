package main

import (
	"fmt"
	"net/http"
	"simple_restful_api_golang/controller"
	"simple_restful_api_golang/database"
	"simple_restful_api_golang/exepciton"
	"simple_restful_api_golang/middleware"
	"simple_restful_api_golang/repository"
	"simple_restful_api_golang/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := database.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.POST("/api/categories", categoryController.Create)

	router.PanicHandler = exepciton.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("server is running on host: http://localhost:3000")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
