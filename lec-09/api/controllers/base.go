package controllers

import (
	"fmt"
	"log"
	"net/http"
	"tfs-02/lec-09/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter().StrictSlash(true)

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port: ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
