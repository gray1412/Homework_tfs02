package handlers

import (
	"asslec6p2/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := database.ConnectToDatabase()
	defer db.Close()

	//save to database
	db.Create(&newProduct)

	//response
	responseWithJson(w, http.StatusCreated, map[string]string{"message": "Created !"})

	fmt.Println("Created !")
}
