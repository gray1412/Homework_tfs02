package request

import (
	"course/validator"	
	"errors"
)


type ReqCreateUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	RoleId    int    `json:"roleid"`
}

func CheckReqCreateUser(req *ReqCreateUser) error {
	if !validator.CheckLength(req.FirstName,20){
		return errors.New("FirstName invalid")
	}
	if !validator.CheckLength(req.LastName,20){
		return errors.New("LastName invalid")
	}
	if !validator.CheckPhone(req.Phone){
		return errors.New("Phone invalid")
	}
	if !validator.CheckMail(req.Email){
		return errors.New("Email invalid")
	}
	if !validator.CheckLength(req.Address,50){
		return errors.New("Address invalid")
	}
	if !validator.CheckLength(req.Username,20){
		return errors.New("Username invalid")
	}
	if !validator.CheckLength(req.Password,20){
		return errors.New("Password invalid")
	}
	return nil
}
type ReqUpdateUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Address   string `json:"address"`
	RoleId    int    `json:"roleid"`
}

func CheckReqUpdateUser(req *ReqUpdateUser) error {
	if !validator.CheckLength(req.FirstName,20){
		return errors.New("FirstName invalid")
	}
	if !validator.CheckLength(req.LastName,20){
		return errors.New("LastName invalid")
	}
	if !validator.CheckLength(req.Address,50){
		return errors.New("Address invalid")
	}
	return nil
}