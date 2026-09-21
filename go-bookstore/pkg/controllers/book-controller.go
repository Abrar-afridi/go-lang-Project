package controller

import (
	"fmt"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create Book")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get Books")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get Book By ID")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update Book")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Book")
}
