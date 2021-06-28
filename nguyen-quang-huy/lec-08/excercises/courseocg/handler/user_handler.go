package handler

import (
	"course/database"
	"course/storage"
	"course/storage/request"
	"course/storage/respone"
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

var db = *database.ConnectDB()

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var listUsers []storage.User
	var listUsersRes []respone.GetUsersRes

	db.Find(&listUsers)
	for _, user := range listUsers {
		listUsersRes = append(listUsersRes, respone.MapGetUserRes(&user))
	}
	respone.ResponseWithJson(w, 200, listUsersRes)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": "Invalid param"})
		return
	}
	var user storage.User
	findPerson := db.First(&user, id)
	if findPerson.Error != nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "User not found"})
		return
	}
	respone.ResponseWithJson(w, 200, respone.MapGetUserByIdRes(&user))
}
func DeleteUser (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": "Invalid param"})
		return
	}
	var user storage.User
	result := db.First(&user, id)
	if result.Error != nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "User not found"})
		return
	}

	result = db.Delete(&storage.User{}, id)
	if result.Error != nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "Delete fail"})
		return
	}
	respone.ResponseWithJson(w, 201, map[string]string{"message": "Deleted"})


}
func UpdateUser (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": "Invalid param"})
		return
	}
	var user storage.User
	findPerson := db.First(&user, id)
	if findPerson.Error != nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "User not found"})
		return
	}

	req := request.ReqUpdateUser{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": "Invalid body"})
		return
	}
	err = request.CheckReqUpdateUser(&req)
	if err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": err.Error()})
		return
	}

	user.FirstName =req.FirstName
	user.LastName =req.LastName
	user.Address =req.Address
	user.RoleId =req.RoleId

	fmt.Println(user)
	result := db.Save(&user)
	if result.Error != nil {
		respone.ResponseWithJson(w, 503, map[string]string{"message": "Update Error"})
		return
	}
	respone.ResponseWithJson(w, 201, map[string]string{"message": "Succes"})
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := request.ReqCreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": "Invalid body"})
		return
	}
	err := request.CheckReqCreateUser(&req)
	if err != nil {
		respone.ResponseWithJson(w, 400, map[string]string{"message": err.Error()})
		return
	}
	user := storage.User{}
	result := db.Where("email = ? OR phone = ? OR username = ?", req.Email, req.Phone,req.Username).First(&user)
	if result.Error == nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "email,phone or user exist"})
		return
	}
	role := storage.Role{}
	result = db.First(&role, req.RoleId)
	if result.Error != nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "role not exist"})
		return
	}
	user = storage.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Phone: req.Phone,
		Address: req.Address,
		Email: req.Email,
		Username: req.Username,
		Password: req.Password,
		RoleId: req.RoleId,
	}
	result = db.Create(&user)
	if result.Error != nil {
		respone.ResponseWithJson(w, 404, map[string]string{"message": "Create fail"})
		return
	}
	
	respone.ResponseWithJson(w, 201, map[string]string{"message": "Succes"})
}

