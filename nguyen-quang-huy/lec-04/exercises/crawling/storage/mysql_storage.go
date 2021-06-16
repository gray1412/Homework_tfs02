package storage

type Film struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Rate  string `db:"rate"`
	Url   string `db:url`
}
