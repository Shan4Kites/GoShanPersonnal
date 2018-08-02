package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

type User struct {
	Id int
	Name string
}
//
//func constructObject(queryResult map[string]interface{}, object interface{}) interface{} {
//	object{id}
//}

func printFieldsOfAnObject(user User) {
	s := reflect.ValueOf(&user).Elem()
	fmt.Println("s", s)
	typeOfT := s.Type()
	fmt.Println("typeOfT", typeOfT)
	fmt.Println("NumField is : ", s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err == nil {
		fmt.Println("Output is : ", b)
	}

	//user := User{Id: 1, Name: "shanmugam"}
	//outputType := reflect.TypeOf(user)
	//object := reflect.New(outputType)
	//fmt.Println("Output is : ", object)

	//printFieldsOfAnObject(object)
	//user := User{}
	//queryResult := make(map[string]interface{})
	//queryResult["Id"] = 1
	//queryResult["Name"] = "shanmugam"
	//constructObject(queryResult, user)



	//fmt.Println(int(reflect.ValueOf(user).Field(0).Int()))
	//reflect.ValueOf(&user).Elem().Field(0).SetInt(321)
	//
	//fmt.Println(int(reflect.ValueOf(user).Field(0).Int()))

}