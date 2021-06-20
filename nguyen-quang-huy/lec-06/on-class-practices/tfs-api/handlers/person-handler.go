package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tfs/tfs-api/storage"

	"github.com/gorilla/mux"
)

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	ResponseWithJson(w, storage.Storage)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ResponseWithJson(w, map[string]string{"message": "Invalid param"})
		return
	}
	var updatePerson storage.Person
	if err = json.NewDecoder(r.Body).Decode(&updatePerson); err != nil {
		ResponseWithJson(w, map[string]string{"message": "Invalid body"})
		return
	}
	updatePerson.Id = id

	for i, person := range storage.Storage {
		if person.Id == id {
			storage.Storage[i] = updatePerson
			ResponseWithJson(w, map[string]string{"message": "Succes"})
			return
		}
	}
	ResponseWithJson(w, map[string]string{"message": "Person not found"})

}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	req := storage.Person{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseWithJson(w, map[string]string{"message": "Invalid body"})
		return
	}
	req.Id = storage.GenerateId()
	storage.Storage = append(storage.Storage, req)
	ResponseWithJson(w, map[string]string{"message": "Succes"})
}
func ResponseWithJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "applacation/json")
	json.NewEncoder(w).Encode(obj)
}
