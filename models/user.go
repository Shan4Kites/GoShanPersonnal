package models

import (
	"fmt"
	"strconv"
)

type User struct {
	Id int
	Name string
	Place string
}

func (self User) GetMapping() (mapper map[string]interface{}) {
	mapper = make(map[string]interface{})
	mapper["id"] = self.Id
	mapper["username"] = self.Name
	mapper["place"] = self.Place
	return mapper
}

func (self User) GetObject(mapper map[string][]byte) (object interface{}) {
	user := User{}

	if mapper["id"] != nil {
		Id, error := strconv.Atoi(string(mapper["id"]))
		if error != nil {
			fmt.Println("error is ", error)
		} else {
			user.Id = Id
		}
	}

	if mapper["username"] != nil {
		user.Name = string(mapper["username"])
	}

	if mapper["place"] != nil {
		user.Place = string(mapper["place"])
	}

	return user
}

func (self User) GetTableName() string {
	return "users"
}