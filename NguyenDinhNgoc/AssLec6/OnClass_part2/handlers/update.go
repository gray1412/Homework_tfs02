package handlers

import (
	"asslec6p2/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var updateStudent database.Student
	err := json.NewDecoder(r.Body).Decode(&updateStudent)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := *database.ConnectToDatabase()
	defer db.Close()

	//find by ID (Primary key)
	var student database.Student
	db.First(&student, updateStudent.ID)

	//check
	if (database.Student{} == student) {
		fmt.Println("There is no such student !")
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such student !"})
		return
	}

	//update and save
	student = updateStudent
	db.Save(&student)

	//reponse
	responseWithJson(w, http.StatusOK, map[string]string{"message": "Updated !"})

	fmt.Println(student)
}
