package db
//5
//Conexão com o banco de dados PostgreSQL

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "click.db")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	db.Exec("INSERT INTO click (name,value) VALUES ('click',0) where not exists (select 1 from click where id = 1)")
	return db, nil
}