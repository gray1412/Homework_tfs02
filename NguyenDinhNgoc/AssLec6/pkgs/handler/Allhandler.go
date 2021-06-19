package handler

import (
	"asslec6/pkgs/data"
	"asslec6/pkgs/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type object struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ParseRequestBodyToObject(w http.ResponseWriter, r *http.Request) {
	newObject := object{}                             //Tạo sẵn một đối tượng rỗng để Parse Json trong r.body vào
	err := json.NewDecoder(r.Body).Decode(&newObject) //Parse to newObject
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	//nếu thành công thì có thể
	fmt.Fprintf(w, "hello %v", newObject.Name)

	//nếu muốn chuyển từ object sang json và nạp vào w
	json.NewEncoder(w).Encode(newObject)
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
func ReadbyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id := vars["id"]

	// //convert to int
	// idint, _ := strconv.Atoi(id)

	idint, _ := strconv.Atoi(vars["id"])
	fmt.Println(idint)
	//find
	student, err := data.Students[idint]
	if !err {
		fmt.Println(err)
		fmt.Fprintf(w, "No rerult !")
		return
	}
	//convert student to json
	// ret, _ := json.Marshal(&student)
	//đóng vào reponse
	//cách 1
	// fmt.Fprintf(w, string(ret))
	//cách 2
	json.NewEncoder(w).Encode(student)
	fmt.Println(data.Students)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	// for _, student := range data.Students {
	// 	studentJson, _ := json.Marshal(&student)
	// 	fmt.Fprintf(w, string(studentJson)+"\n")
	// }
	json.NewEncoder(w).Encode(data.Students)
	fmt.Println(data.Students)
}
func Create(w http.ResponseWriter, r *http.Request) {
	newStudent := storage.Student{}
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	idCreate := newStudent.ID
	//check
	_, err2 := data.Students[idCreate]
	if err2 {
		fmt.Fprintf(w, "ID has been used !")
		return
	}

	//create
	data.Students[idCreate] = newStudent

	//response
	responseWithJson(w, http.StatusCreated, newStudent)

	//print
	fmt.Println(data.Students)
}
func Update(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
	// 	return
	// }

	var updateStudent storage.Student
	err := json.NewDecoder(r.Body).Decode(&updateStudent)

	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	idNeedUpdate := updateStudent.ID

	//check
	_, err2 := data.Students[idNeedUpdate]
	if err2 {
		fmt.Fprintf(w, "There is no such student !")
		return
	}

	//update
	data.Students[idNeedUpdate] = updateStudent

	//response
	responseWithJson(w, http.StatusOK, updateStudent)

	//print
	fmt.Println(data.Students)

	// for i := range data.Students {
	// 	if i == 1 {
	// 		data.Students[i] = updateStudent
	// 		responseWithJson(w, http.StatusOK, updateStudent)
	// 		fmt.Println(data.Students)
	// 		return
	// 	}
	// }
	// fmt.Println(data.Students)
	// responseWithJson(w, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

const JsonContentType = "application/json"

func ContentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type: application/json")
			return
		}

		next.ServeHTTP(w, r)
	})

}
