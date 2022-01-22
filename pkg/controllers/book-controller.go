package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/baijupadmanabhan/golang-bookstore/pkg/models"
	"github.com/baijupadmanabhan/golang-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
)

//var NewBook models.Book

type bookData struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	books := []bookData{}
	res, _ := json.Marshal(newBook)
	copier.Copy(&books, &newBook)
	println("Struct value : ", books)
	resp, _ := json.Marshal(books)
	println("Value of response : ", resp)
	println("Value of res : ", res)
	w.Header().Set("Content.Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	books := []bookData{}
	copier.Copy(&books, &bookDetails)
	println("Get book by id db : ", db)
	res, _ := json.Marshal(books)
	w.Header().Set("Content.Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	books := []bookData{}
	copier.Copy(&books, &b)
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	books := []bookData{}
	copier.Copy(&books, &book)
	res, _ := json.Marshal(books)
	w.Header().Set("Content.Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
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
	books := []bookData{}
	copier.Copy(&books, &bookDetails)
	res, _ := json.Marshal(books)
	w.Header().Set("Content.Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
