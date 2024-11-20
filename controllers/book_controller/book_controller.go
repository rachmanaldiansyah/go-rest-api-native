package book_controller

import (
	"net/http"

	"go-api-native/config"
	"go-api-native/helpers"
	"go-api-native/models"
)

// Index will return all books from the database
// @Summary Get all books
// @Description Get all books
// @ID get-all-books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /books [get]
func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "List Books", books)
}
