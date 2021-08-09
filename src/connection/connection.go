package connection

import (
	"database/sql"
	"log"
)

var db *sql.DB

// GetConnection use one time
func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "todo.db")

	if err != nil {
		panic(err)
	}

	log.Println("DB open")
	return db
}

// Migrations create table/s
func Migrations() error {

	db := GetConnection()

	taskTable := `
    CREATE TABLE task (
      "id" INTEGER PRIMARY KEY AUTOINCREMENT,
      "title" varchar(64) NOT NULL,
      "description" varchar(200) NOT NULL,
      "status" varchar(1) NOT NULL,
      "created_at" TIMESTAMP DEFAULT DATETIME,
      "updated_at" TIMESTAMP NOT NULL
    );
  `

	_, err := db.Exec(taskTable)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
