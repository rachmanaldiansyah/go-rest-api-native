package controllers

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

// Index will return all authors from the database
// @Summary Get all authors
// @Description Get all authors
// @ID get-all-authors
// @Tags authors
// @Accept  json
// @Produce  json
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /authors [get]
func Index(w http.ResponseWriter, r *http.Request) {
	var author []models.Author

	if err := config.DB.Find(&author).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "List Author's", author)
}

// Create will create a new author in the database
// @Summary Create author
// @Description Create author
// @ID create-author
// @Tags authors
// @Accept  json
// @Produce  json
// @Param author body models.Author true "Author data"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /authors [post]
func Create(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	if err := config.DB.Create(&author).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Successfully created author", nil)
}

// Detail will return the detail of an author by ID
// @Summary Get author by ID
// @Description Get author by ID
// @ID get-author-by-id
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /authors/{id}/detail [get]
func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Author not found", nil)
			return
		}

		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Detail Author", author)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Author not found", nil)
			return
		}

		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&author).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Successfully updated author", nil)
}
