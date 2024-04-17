package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/maratb3k/go-bookstore/pkg/models"
	"github.com/maratb3k/go-bookstore/pkg/utils"
)

var NewAuthor models.Author

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	newAuthors := models.GetAllAuthors()
	res, _ := json.Marshal(newAuthors)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("authorId")
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	authorDetails, _ := models.GetAuthorById(ID)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	CreateAuthor := &models.Author{}
	utils.ParseBody(r, CreateAuthor)
	a := CreateAuthor.CreateAuthor()
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("authorId")
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	author := models.DeleteAuthorById(ID)
	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "pgklication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var updateAuthor = &models.Author{}
	utils.ParseBody(r, updateAuthor)
	authorId := r.URL.Query().Get("authorId")
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	authorDetails, db := models.GetAuthorById(ID)
	if updateAuthor.FirstName != "" {
		authorDetails.FirstName = updateAuthor.FirstName
	}
	if updateAuthor.LastName != "" {
		authorDetails.LastName = updateAuthor.LastName
	}
	if updateAuthor.MiddleName != "" {
		authorDetails.MiddleName = updateAuthor.MiddleName
	}
	db.Save(&authorDetails)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "pgklication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
