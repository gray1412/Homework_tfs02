package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (u *UserInfo) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(b)
}

type User struct {
	gorm.Model
	FirstName      string
	LastName       string
	Email          string `gorm:"unique_index:user_email_index"`
	Password       string
	Token          string
	TokenExpiresAt uint
}

func main() {
	// db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/crawl")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// fmt.Println(1)
	// _, err = db.Exec("CREATE TABLE tfs_users(id int NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(100));")
	// if err != nil {
	// 	fmt.Println("Cannot create table: ", err)
	// 	return
	// }
	// _, err = db.Exec(`INSERT INTO tfs_users (name) VALUES ("Truong");`)
	// if err != nil {
	// 	fmt.Println("Cannot insert into table: ", err)
	// 	return
	// }
	// rows, err := db.Query("SELECT * FROM tfs_users LIMIT 1")
	// if err != nil {
	// 	fmt.Println("Cannot get from users")
	// 	return
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	id := int64(0)
	// 	name := ""
	// 	err = rows.Scan()
	// 	if err == nil {
	// 		fmt.Println(id, name)
	// 	}
	// }

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}

	if db.HasTable(&User{}) == false {
		db.CreateTable(&User{})
	}

	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "insecurepassword",
	}

	db.Create(&user)

	var users []User
	db.Find(&users)

	fmt.Println("There are", len(users), "user records in the table.")

}
