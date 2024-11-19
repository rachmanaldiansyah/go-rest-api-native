package controllers

import (
	"encoding/json"
	"net/http"

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

func Detail(w http.ResponseWriter, r *http.Request) {

}
