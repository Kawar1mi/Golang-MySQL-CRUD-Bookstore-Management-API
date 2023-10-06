package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kawar1mi/go-bookstore/pkg/models"
	"github.com/Kawar1mi/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	bookId, ok := utils.ParseParamBookId(mux.Vars(r))
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "error while parsing param ID")
		return
	}

	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}
	err := utils.ParseBody(r, newBook)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "error while parsing ID")

		return
	}

	b := newBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	bookId, ok := utils.ParseParamBookId(mux.Vars(r))
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "error while parsing param ID")
		return
	}

	var updateBook = &models.Book{}
	err := utils.ParseBody(r, updateBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "error while parsing ID")
		return
	}

	bookDetails, db := models.GetBookById(bookId)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	bookId, ok := utils.ParseParamBookId(mux.Vars(r))
	if !ok {
		fmt.Fprintln(w, "error while parsing ID")
		return
	}

	b := models.DeleteBook(bookId)
	res, _ := json.Marshal(b)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
