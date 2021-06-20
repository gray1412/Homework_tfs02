package main

import (
	storage "Api/Storage"
	apiStudents "Api/students"
)

func main() {
	storage.Connect()
	storage.Migrate()
	apiStudents.PathStudent()
}
