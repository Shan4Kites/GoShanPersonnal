package orm

import (
	"database/sql"
)

type Users struct {
	DB *sql.DB
	Id int
	UserName string
}

/*
This creates a new record and stores it
 */
func (self Users) Create()  {
	sqlStatement := `INSERT INTO users (id, username) VALUES ($1, $2)`
	_, err := self.DB.Exec(sqlStatement, self.Id, self.UserName)
	if err != nil {
		panic(err)
	}
}

/*
This updates the record. The record needs to be fetched for updating.
 */
func (self Users) Update()  {
	sqlStatement := `UPDATE users SET username = $2 WHERE id = $1`
	_, err := self.DB.Exec(sqlStatement, self.Id, self.UserName)
	if err != nil {
		panic(err)
	}
}

/*
This deletes the record. The record needs to be fetched for deleting.
 */
func (self Users) Delete()  {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := self.DB.Exec(sqlStatement, self.Id)
	if err != nil {
		panic(err)
	}
}

/*
This gets all the records matches the condition
this expects query in this format : "id = $1" followed by list of values
 */
func (self Users) Where(query string, args ...interface{})  (usersList []Users) {
	sqlStatement := "select username, id from users where " + query
	rows, err := self.DB.Query(sqlStatement, args...)
	if err != nil {
		panic(err)
	}
	usersList = []Users{}
	for rows.Next() {
		var id int
		var username string
		err = rows.Scan(&username, &id)
		user:= Users{Id: id, UserName: username}
		usersList = append(usersList, user)
	}
	return
}
