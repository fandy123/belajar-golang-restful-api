package main

import (
	"fandy123/belajar-golang-restful-api/app"
	"fandy123/belajar-golang-restful-api/controller"
	"fandy123/belajar-golang-restful-api/exception"
	"fandy123/belajar-golang-restful-api/helper"
	"fandy123/belajar-golang-restful-api/repository"
	"fandy123/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryContoller := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryContoller.FindAll)
	router.GET("/api/categories/:categoryId", categoryContoller.FindById)
	router.POST("/api/categories", categoryContoller.Create)
	router.PUT("/api/categories/:categoryId", categoryContoller.Update)
	router.DELETE("/api/categories/:categoryId", categoryContoller.Delete)

	// router.PanicHandler = exception.ErrorHandler()

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
