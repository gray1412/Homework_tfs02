package respone

import "course/storage"

type GetUsersRes struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname" `
	RoleId    int    `json:"roleid"`
	// UserClassList []UserClass `json:"userclasslist" gorm:"foreignKey:UserId"`
}

func MapGetUserRes(sourse *storage.User) GetUsersRes {
	return GetUsersRes{
		Id : sourse.Id,
		FirstName: sourse.FirstName,
		LastName: sourse.LastName,
		RoleId: sourse.RoleId,

	}
	
}
type GetUserByIdRes struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname" `
	Phone     string `json:"phone"`
	Email     string `json:"email" `
	Address   string `json:"address" `
	RoleId    int    `json:"roleid"`
}

func MapGetUserByIdRes(sourse *storage.User) GetUserByIdRes {
	return GetUserByIdRes{
		Id : sourse.Id,
		FirstName: sourse.FirstName,
		LastName: sourse.LastName,
		Phone: sourse.Phone,
		Email: sourse.Email,
		Address: sourse.Address,
		RoleId: sourse.RoleId,
	}
	
}
