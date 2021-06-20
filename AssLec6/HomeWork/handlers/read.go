package handlers

import (
	"asslec6p2/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ReadAll(w http.ResponseWriter, r *http.Request) {
	//connect database
	db := *database.ConnectToDatabase()
	defer db.Close()

	//find
	var ListProducts []database.Product
	db.Find(&ListProducts)

	//response
	responseWithJson(w, http.StatusOK, ListProducts)

	fmt.Println(ListProducts)
}

func ReadbyID(w http.ResponseWriter, r *http.Request) {
	//get ID
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//connect database
	db := *database.ConnectToDatabase()
	defer db.Close()

	//find by ID (Primary key)
	var product database.Product
	db.First(&product, id)

	//Check result
	if (database.Product{} == product) {
		fmt.Println("Not Found !")
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Not Found!"})
		return
	}

	//response
	responseWithJson(w, http.StatusOK, product)

	fmt.Println(product)
}
