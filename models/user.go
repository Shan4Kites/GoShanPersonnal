package models

type User struct {
	Id int
	Name string
}

func (self User) GetMapping() (mapper map[string]interface{}) {
	mapper = make(map[string]interface{})
	mapper["id"] = self.Id
	mapper["name"] = self.Name
	return mapper
}