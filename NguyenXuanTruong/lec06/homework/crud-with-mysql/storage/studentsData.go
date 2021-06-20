package storage

// Struct Student bao gồm các thông tin cơ bản của học sinh
type Student struct {
	Id       uint   `gorm:"primaryKey"` // Dữ liệu Id được đánh số tự động tăng dần khi đưa vào db
	FullName string `db:"fullname"`
	Age      int    `db:"age"`
	Phone    string `db:"phone"`
}

// Tạo một slice để lưu dữ liệu truyền vào vào memory
var Students []Student

// func init() {
// 	Students = []Student{
// 		{FullName: "Truong", Age: 20, Phone: "22222"},
// 		{FullName: "haha", Age: 20, Phone: "22222"},
// 		{FullName: "Trduong", Age: 20, Phone: "22222"},
// 		{FullName: "Trufsfadfong", Age: 20, Phone: "22222"},
// 	}
// }

type MemoryStorageDataStudent struct {
	M map[int64]Student
}

//var mapOfStudents map[uint]Student

// func init() {
// 	Students = map[int]Student{
// 		1: Student{
// 			"Truong", 21, "2222299",
// 		},
// 		2: Student{
// 			"Duy Anh", 22, "22255522",
// 		},
// 		3: Student{
// 			"haha", 50, "229992",
// 		},
// 	}
// }
