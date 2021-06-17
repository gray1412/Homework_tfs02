package storage

type Memory_Storage_Film struct {
	M map[int64]Film
}
type Memory_Storage_Item struct {
	M map[int64]Item //kieu Item duoc khai bao trong mysql_storage.go
}

type Memory_Storage_YouTube struct {
	M map[int64]YouTube
}
