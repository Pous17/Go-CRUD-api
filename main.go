package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)
	RegisterCategoryRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}

func RegisterCategoryRoutes(router *mux.Router) {
	router.HandleFunc("/api/category", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/api/category/{id}", controllers.GetCategoryById).Methods("GET")
	router.HandleFunc("/api/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/category/{id}", controllers.DeleteCategory).Methods("DELETE")
}
