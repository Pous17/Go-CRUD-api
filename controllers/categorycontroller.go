package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var category entities.Category
	json.NewDecoder(r.Body).Decode(&category)
	database.Instance.Create(&category)
	json.NewEncoder(w).Encode(category)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["id"]
	if !CheckIfCategoryExists(categoryId) {
		json.NewEncoder(w).Encode("Category not found")
	}
	w.Header().Set("Category-type", "application/json")
	var category entities.Category
	database.Instance.First(&category, categoryId)
	json.NewEncoder(w).Encode(category)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var categories []entities.Category
	database.Instance.Find(&categories)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["id"]
	if !CheckIfCategoryExists(categoryId) {
		json.NewEncoder(w).Encode("Category not found")
	}
	w.Header().Set("Category-type", "application/json")
	var category entities.Category
	database.Instance.First(&category, categoryId)
	json.NewDecoder(r.Body).Decode(&category)
	database.Instance.Save(&category)
	json.NewEncoder(w).Encode(category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(r)["id"]
	if !checkIfProductExists(categoryId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Category Not Found!")
		return
	}
	var category entities.Category
	database.Instance.Delete(&category, categoryId)
	json.NewEncoder(w).Encode("Category Deleted Successfully!")
}

func CheckIfCategoryExists(categoryID string) bool {
	var category entities.Category
	database.Instance.First(&category, categoryID)
	return category.ID != 0
}
