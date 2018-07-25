package main

import "fmt"
import (
	"github.com/Shan4Kites/GoShanPersonnal/orm"
	"github.com/Shan4Kites/GoShanPersonnal/initializer"
)


func main() {
	initializer.InitDb()
	defer initializer.DB.Close()
	user := orm.Users{DB: initializer.DB, Id: 4}
	//user.Create()
	//user.UserName = "fk_updated"
	//user.Update()
	res := user.Where("id=$1", 4)
	fmt.Println("users : " , res)
}