package controllers

import (
	"net/http"
	"tfs-02/lec-09/api/utils"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
