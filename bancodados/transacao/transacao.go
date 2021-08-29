package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	smt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")
	smt.Exec(2000, "Bia")
	smt.Exec(2000, "Carlos")

	_, err = smt.Exec(1, "Tiago")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
