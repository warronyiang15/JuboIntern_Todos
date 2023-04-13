package db

import (
	"database/sql"
	"fmt"
	"todos/utility"

	_ "github.com/go-sql-driver/mysql"
)

// Database Data
const (
	UserName string = "root"
	Password string = ""
	Addr     string = "127.0.0.1"
	Port     int    = 3306
	Database string = "jubotodos"
)

// Establish a connection to the database
func EstablishConnection() (db *sql.DB) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	db, err := sql.Open("mysql", conn)

	ret := initializedDatabase(db)

	// Something went wrong while initialize
	if ret == 0 || err != nil {
		return nil
	}

	// Otherwise successfully established connection
	return db
}

// Initialized action of database
func initializedDatabase(db *sql.DB) (ret int) {
	ret = CreateTable(db)
	return ret
}

// Create Table in Database
func CreateTable(db *sql.DB) (ret int) {
	sql := `CREATE TABLE IF NOT EXISTS todos(
		ID bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		Title VARCHAR(255),
		Description VARCHAR(512),
		Completed INT(2) NOT NULL DEFAULT 0,
		CreatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(sql); err != nil {
		return 0
	}
	return 1
}

// GetAllTodos
func GetAllTodosFromDB(db *sql.DB) (ret int, todoArray []utility.Todos) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return 0, nil
	}
	defer rows.Close()

	var todosArray []utility.Todos

	// Scan Through the query rows and append it to return Array
	for rows.Next() {
		var newtodos utility.Todos
		err = rows.Scan(&newtodos.Id, &newtodos.Title, &newtodos.Description, &newtodos.Completed, &newtodos.CreatedAt)
		if err != nil {
			return 0, nil
		}
		todosArray = append(todosArray, newtodos)
	}

	return 1, todosArray
}

// Get A todos
func GetTodosWithIDFromDB(db *sql.DB, id int) (ret int, todos utility.Todos) {
	var newtodos utility.Todos
	row := db.QueryRow("SELECT * FROM todos where id=?", id)

	// Scan and check existences
	if err := row.Scan(&newtodos.Id, &newtodos.Title, &newtodos.Description, &newtodos.Completed, &newtodos.CreatedAt); err != nil {
		return 0, newtodos
	}

	return 1, newtodos
}

// Create A todos to database, return non-zero ID if successfuly
func CreateTodosToDB(db *sql.DB, todos utility.Todos) (retID int64) {
	sql := `INSERT INTO todos(Title, Description) VALUES (?, ?)`
	result, err := db.Exec(sql, todos.Title, todos.Description)
	if err != nil {
		return 0
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0
	}

	return lastInsertID
}

// Update a todos to database with given ID, return non-zero ID if successful
func UpdateTodosToDB(db *sql.DB, todos utility.Todos) (retID int) {
	sql := "UPDATE todos SET Title=?, Description=?, Completed=?, CreatedAt=? where id=?"
	_, err := db.Exec(sql, todos.Title, todos.Description, todos.Completed, todos.CreatedAt, todos.Id)
	if err != nil {
		return 0
	}
	return todos.Id
}

// Delete a todos in database with given ID, return 1 if successful, otherwise 0
func DeleteTodosFromDB(db *sql.DB, id int) (ret int) {
	sql := "DELETE FROM todos WHERE id=?"
	_, err := db.Exec(sql, id)
	if err != nil {
		return 0
	}

	return 1
}
