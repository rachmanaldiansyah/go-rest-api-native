package book_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// Detail will return the detail of a book by ID
// @Summary Get book by ID
// @Description Get book by ID
// @ID get-book-by-id
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /books/{id}/detail [get]
func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var books models.Book
	var booksResponse models.BookResponse

	if err := config.DB.Joins("Author").First(&books, id).First(&booksResponse, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Book not found", nil)
			return
		}

		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Detail Book", booksResponse)
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

// Update will update an existing book in the database based on the provided ID.
// @Summary Update book
// @Description Update book
// @ID update-book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book data"
// @Success 200 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /books/{id}/update [put]
func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var books models.Book
	var booksResponse models.BookResponse

	if err := config.DB.First(&books, id).First(&booksResponse, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Book not found", nil)
			return
		}

		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var booksPayload models.Book
	if err := json.NewDecoder(r.Body).Decode(&booksPayload); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author
	if booksPayload.AuthorID != 0 {
		if err := config.DB.First(&author, booksPayload.AuthorID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helpers.Response(w, 404, "Author not found", nil)
				return
			}

			helpers.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&booksPayload).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Update book successfully", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var books models.Book
	res := config.DB.Delete(&books, id)
	if res.Error != nil {
		helpers.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, 404, "Book not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully deleted book", nil)
}
