package handler

import (
	"encoding/json"
	"fmt"
	"homework/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const JsonContentType = "application/json"

// middleware checking
func ContenTypeCheckingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// print content header type
func ContentHeaderType(w http.ResponseWriter, r *http.Request) {
	contentHeaderType := r.Header.Get("Content-Type")
	fmt.Fprintf(w, "Hello with content type %v", contentHeaderType)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// get data from json body
	product := storage.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}
	// connect database and add new student to it
	db := *storage.ConnectDatabase()
	defer db.Close()

	// check if id has available
	_, productCheck := storage.Products[product.Id]
	if productCheck {
		fmt.Fprintln(w, "Id has been used")
		return
	}
	// save to database and memory
	storage.Products[product.Id] = product
	db.Create(&product)
	fmt.Println("Created succesful !")
}

func ReadProduc(w http.ResponseWriter, r *http.Request) {
	//connect database
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := *storage.ConnectDatabase()
	defer db.Close()

	//find list item in database
	var ListProducts []storage.Product
	db.Find(&ListProducts)

	//response
	json.NewEncoder(w).Encode(ListProducts)
	fmt.Println(ListProducts)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get infor form body
	product := storage.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}

	// connect database to update product
	db := *storage.ConnectDatabase()
	defer db.Close()

	// find product by ID
	productUpdate := storage.Product{}
	db.Find(&productUpdate, product.Id)

	// check if the same student in storage
	if productUpdate == product {
		fmt.Fprintln(w, "Same as the student in the storage")
		fmt.Println(w, "There is no update product")
		return
	}

	// save
	productUpdate = product
	db.Save(&productUpdate)
	storage.Products[product.Id] = product

	// response
	json.NewEncoder(w).Encode(productUpdate)
	fmt.Println(productUpdate)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id of student will be delete
	vars := mux.Vars(r)
	idDelete, _ := strconv.Atoi(vars["id"])

	// check if hasn't product to delete
	_, productCheck := storage.Products[idDelete]
	if !productCheck {
		fmt.Fprintf(w, "Don't have student with id: %v", idDelete)
		fmt.Printf("Don't have student with id: %v", idDelete)
		return
	}

	// connect database to delete product
	db := *storage.ConnectDatabase()
	defer db.Close()

	// find product will be delete
	productDelete := storage.Product{}
	db.Find(&productDelete, idDelete)

	// delete student has idDelete
	delete(storage.Products, idDelete)
	db.Delete(&storage.Product{}, idDelete)
	fmt.Fprintf(w, "Student with id %v has been delete", idDelete)
	fmt.Printf("Student with id %v has been delete", idDelete)
}
