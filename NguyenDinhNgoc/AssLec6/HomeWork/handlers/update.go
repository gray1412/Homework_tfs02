package handlers

import (
	"asslec6p2/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var updateProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := *database.ConnectToDatabase()
	defer db.Close()

	//find by ID (Primary key)
	var product database.Product
	db.First(&product, updateProduct.ID)

	//check
	if (database.Product{} == product) {
		fmt.Println("There is no such product !")
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such product !"})
		return
	}

	//update and save
	product = updateProduct
	db.Save(&product)

	//reponse
	responseWithJson(w, http.StatusOK, map[string]string{"message": "Updated !"})

	fmt.Println(product)
}
