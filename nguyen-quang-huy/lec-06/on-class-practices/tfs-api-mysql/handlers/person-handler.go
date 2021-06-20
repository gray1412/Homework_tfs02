package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tfs/tfs-api-mysql/database"
	"tfs/tfs-api-mysql/storage"

	"github.com/gorilla/mux"
)

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	db := *database.ConnectDB()
	defer db.Close()
	var listStudents []storage.Person
	db.Find(&listStudents)
	ResponseWithJson(w, 200, listStudents)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ResponseWithJson(w, 400, map[string]string{"message": "Invalid param"})
		return
	}
	db := *database.ConnectDB()
	var person storage.Person
	defer db.Close()
	findPerson := db.First(&person, id)
	if findPerson.Error != nil {
		ResponseWithJson(w, 404, map[string]string{"message": "Person not found"})
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&person); err != nil {
		ResponseWithJson(w, 400, map[string]string{"message": "Invalid body"})
		return
	}
	result := db.Save(&person)
	if result.Error != nil {
		ResponseWithJson(w, 503, map[string]string{"message": "Update Error"})
		return
	}
	ResponseWithJson(w, 201, map[string]string{"message": "Succes"})

}
func CreatePerson(w http.ResponseWriter, r *http.Request) {

	req := storage.Person{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseWithJson(w, 400, map[string]string{"message": "Invalid body"})
		return
	}
	db := *database.ConnectDB()
	defer db.Close()
	db.Create(&req)
	ResponseWithJson(w, 201, map[string]string{"message": "Succes"})
}
func ResponseWithJson(w http.ResponseWriter, status int, obj interface{}) {
	w.Header().Set("Content-Type", "applacation/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(obj)
}
