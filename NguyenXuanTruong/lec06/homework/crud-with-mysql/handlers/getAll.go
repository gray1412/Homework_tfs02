package handlers

import (
	p "CallPkgs/pkg"
	s "CallPkgs/storage"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	// Connect với database
	db := *p.ConnectStudentDB()
	defer db.Close()

	// Kiểm tra nếu trong db chưa có bảng thì tạo bảng
	if (!db.HasTable(&s.Student{})) {
		db.CreateTable(&s.Student{})
	}

	var students []s.Student
	db.Find(&students) // get all data of students table

	// Hiển thị kết quả trong Body
	p.ResponseWithJson(w, http.StatusOK, students)
}
