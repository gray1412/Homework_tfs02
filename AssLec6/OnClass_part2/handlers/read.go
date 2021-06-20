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
	var ListStudent []database.Student
	db.Find(&ListStudent)

	//response
	responseWithJson(w, http.StatusOK, ListStudent)

	fmt.Println(ListStudent)
}

func ReadbyID(w http.ResponseWriter, r *http.Request) {
	//get ID
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//connect database
	db := *database.ConnectToDatabase()
	defer db.Close()

	//find by ID (Primary key)
	var student database.Student
	db.First(&student, id)

	//Check result
	if (database.Student{} == student) {
		fmt.Println("Not Found !")
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Not Found!"})
		return
	}

	//response
	responseWithJson(w, http.StatusOK, student)

	fmt.Println(student)
}
