package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()
	defer db.Close()
	http.HandleFunc("/create", createUserHandler)
	http.HandleFunc("/insert", insertUserHandler)
	http.HandleFunc("/update", updateUserHandler)
	http.HandleFunc("/delete", deleteUserHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `drop table users`
	_, err := db.Exec(sqlStatement)
	sqlStatement = `CREATE TABLE users (id integer, username VARCHAR (50))`
	//sqlStatement := `INSERT INTO public.account (username) VALUES('shanmugam')`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func insertUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `INSERT INTO users (id, username) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, 1, "shanmugam")
	if err != nil {
		panic(err)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `UPDATE users SET username = $2 WHERE id = $1`
	_, err := db.Exec(sqlStatement, 1, "shanmugam-updated")
	if err != nil {
		panic(err)
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])
	fmt.Println(psqlInfo)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	//host, ok := os.LookupEnv(dbhost)
	//if !ok {
	//	panic("DBHOST environment variable required but not set")
	//}
	//port, ok := os.LookupEnv(dbport)
	//if !ok {
	//	panic("DBPORT environment variable required but not set")
	//}
	//user, ok := os.LookupEnv(dbuser)
	//if !ok {
	//	panic("DBUSER environment variable required but not set")
	//}
	//password, ok := os.LookupEnv(dbpass)
	//if !ok {
	//	panic("DBPASS environment variable required but not set")
	//}
	//name, ok := os.LookupEnv(dbname)
	//if !ok {
	//	panic("DBNAME environment variable required but not set")
	//}
	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "subashni"
	conf[dbpass] = "suba"
	conf[dbname] = "online_appt_development"
	return conf
}