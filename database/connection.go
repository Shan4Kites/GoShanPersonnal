package main


import (
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/Shan4Kites/GoShanPersonnal/initializer"
)

func main() {
	initializer.InitDb()
	defer initializer.DB.Close()
	http.HandleFunc("/create", createUserHandler)
	http.HandleFunc("/insert", insertUserHandler)
	http.HandleFunc("/update", updateUserHandler)
	http.HandleFunc("/delete", deleteUserHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `drop table users`
	_, err := initializer.DB.Exec(sqlStatement)
	sqlStatement = `CREATE TABLE users (id integer, username VARCHAR (50))`
	//sqlStatement := `INSERT INTO public.account (username) VALUES('shanmugam')`
	_, err = initializer.DB.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func insertUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `INSERT INTO users (id, username) VALUES ($1, $2)`
	_, err := initializer.DB.Exec(sqlStatement, 1, "shanmugam")
	if err != nil {
		panic(err)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `UPDATE users SET username = $2 WHERE id = $1`
	_, err := initializer.DB.Exec(sqlStatement, 1, "shanmugam-updated")
	if err != nil {
		panic(err)
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := initializer.DB.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
