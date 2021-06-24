package storage

func Migrate() {
	CreateRole()
	CreateCourse()
	CreateLesstion()
	CreateRoom()
	CreateSlot()
	CreateClass()
	CreatePerson()
	CreateRegistration()
}
