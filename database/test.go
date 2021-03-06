package main

import (
	"github.com/Shan4Kites/GoShanPersonnal/orm"
	"fmt"
	"github.com/Shan4Kites/GoShanPersonnal/models"
	"github.com/Shan4Kites/GoShanPersonnal/initializer"
)

func main() {
	initializer.InitDb()
	defer initializer.DB.Close()
	//user := orm.Users{DB: initializer.DB, Id: 4}
	////user.Create()
	////user.UserName = "fk_updated"
	////user.Update()
	//res := user.Where("id=$1", 4)
	//fmt.Println("users : " , res)

	//user := models.User{Id: 1, Name: "shanmugam", Place: "chennai"}
	//fmt.Println("F user", user)
	//activeRecord := orm.ActiveRecord{DB: initializer.DB, ObjectMapper: &user}
	//activeRecord.Create()
	//user.Name = "shan-updated"
	//activeRecord.Update()

	user := models.User{}
	activeRecord := orm.ActiveRecord{DB: initializer.DB, ObjectMapper: &user}
	//Reading the existing record
	res := activeRecord.Where("username=$1", "shan")
	fmt.Println("S user", res)
}