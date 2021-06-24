package students

import (
	Api "Api/Storage"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func PathStudent() {
	r := mux.NewRouter().StrictSlash(true)
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/students").HandlerFunc(viewStudent)
	get.Use(contentTypeJson)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/students").HandlerFunc(addStudent)
	r.Methods(http.MethodPost).Path("/upload").HandlerFunc(uploadFile)
	post.Use(contentTypeJson)
	put := r.Methods(http.MethodPut).Subrouter()
	put.Path("/students").HandlerFunc(updateStudent)
	put.Use(contentTypeJson)
	delete := r.Methods(http.MethodDelete).Subrouter()
	delete.Path("/students").HandlerFunc(deleteStudent)
	delete.Use(contentTypeJson)
	http.Handle("/", r)
	handler := cors.New(cors.Options{
        AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS","PUT"},
    }).Handler(r)
    // log.Fatal(http.ListenAndServe(":8082", handler))
	http.ListenAndServe(":8000", handler)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}
func viewStudent(w http.ResponseWriter, r *http.Request) {
	db := *Api.Connect()
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	var s Api.Students
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := Api.Connect()
	var Studens = Api.Students{Name: s.Name, Address: s.Address, Phone: s.Phone, Age: s.Age}
	db.Create(&Studens)
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	var s Api.Students
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(s)
	db := Api.Connect()
	// var Studens = Api.Students{Id:1}
	// db.Model(&Studens).Update(s.Name, s.Address, s.Phone, s.Age)
	var Students = Api.Students{Id: s.Id}
	db.First(&Students)
	Students.Name = s.Name
	Students.Age = s.Age
	Students.Address = s.Address
	Students.Phone = s.Phone
	db.Save(&Students)
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
}
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	var s Api.Students
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := Api.Connect()
	var Students = Api.Students{Id: s.Id}
	db.Delete(&Students)
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
}
func contentTypeJson(next http.Handler) http.Handler {
	const contentType = "application/json"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != contentType {
			fmt.Fprintf(w, "Only allow request with content type %v", contentType)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func contentTextPlan(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
