package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/* open db */
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Printf("Failed to open db: %s\n", err)
		return
	}
	defer db.Close()

	/* create table */
	_, err = db.Exec("create table sample (id integer not null primary key, misc text)")
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		return
	}

	/* insert */
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Failed to begin transaction: %s\n", err)
		return
	}

	stmt1, err := tx.Prepare("insert into sample (id, misc) values(?, ?)")
	if err != nil {
		fmt.Printf("Failed to prepare insert query: %s\n", err)
		return
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(1, "sample 1")
	if err != nil {
		fmt.Printf("Failed to exec stmt: %s\n", err)
		return
	}
	_, err = stmt1.Exec(2, "sample 2")
	if err != nil {
		fmt.Printf("Failed to exec stmt: %s\n", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Failed to commit transaction: %s\n", err)
	}

	/* select all */
	rows, err := db.Query("select id, misc from sample")
	if err != nil {
		fmt.Printf("Failed to execute query: %s\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var misc string
		if err = rows.Scan(&id, &misc); err != nil {
			fmt.Printf("Failed to scan row: %s\n", err)
			return
		}
		fmt.Printf("id: %d, misc: %s\n", id, misc)
	}

	/* select one */
	stmt2, err := db.Prepare("select misc from sample where id = ?")
	if err != nil {
		fmt.Printf("Failed to prepare select query: %s\n", err)
		return
	}
	defer stmt2.Close()

	id := 1
	var misc string
	err = stmt2.QueryRow(id).Scan(&misc)
	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No record with this id: %d, error: %s\n", id, err)
		return
	case err != nil:
		fmt.Printf("Failed to execute query: %s\n", err)
		return
	default:
		fmt.Printf("id: %d, misc: %s\n", id, misc)
	}
}
