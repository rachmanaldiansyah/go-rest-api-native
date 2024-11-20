package book_controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"

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
	var booksResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&books).Find(&booksResponse).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "List Books", booksResponse)
}

// Create will create a new book in the database
// @Summary Create book
// @Description Create book
// @ID create-book
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book data"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /books [post]
func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// Check Author's
	var author models.Author
	if err := config.DB.First(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Author not found", nil)
			return
		}

		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Create book successfully", nil)
}
