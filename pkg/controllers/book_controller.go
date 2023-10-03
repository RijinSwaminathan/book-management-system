package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RijinSwaminathan/book-management-system/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error occured", err)
	}
	getBooks, _ := models.GetBookById(ID)
	res, _ := json.Marshal(getBooks)
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
