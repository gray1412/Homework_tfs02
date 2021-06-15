package Storage

func MigrateFilm() {
	db := Connect()
	if db.HasTable(&Imdb{}) == false {
		db.CreateTable(&Imdb{})
	}
}
func MigrateHomedecor() {
	db := Connect()
	if db.HasTable(&Product{}) == false {
		db.CreateTable(&Product{})
	}
}
