package storage


type Person struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
	Age int `json: "age"`
}

var Storage []Person

func init() {
	Storage = []Person{
		{Id: 1, Name: "Huy", Age: 23},
		{Id: 2, Name: "Hoa", Age: 24},
		{Id: 3, Name: "Hung", Age: 25},
		{Id: 4, Name: "Hau"},

	}
}
func GenerateId() int{
	var maxId int
	for _, person := range Storage{
		if person.Id > maxId {
			maxId = person.Id
		}
	}
	return maxId+1
}