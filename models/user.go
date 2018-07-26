package models

type User struct {
	Id int
	Name string
}

func (self User) GetMapping() (mapper map[string]interface{}) {
	mapper = make(map[string]interface{})
	mapper["id"] = self.Id
	mapper["username"] = self.Name
	return mapper
}

func (self User) GetTableName() string {
	return "users"
}