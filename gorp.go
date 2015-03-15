package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

type sample struct {
	Id   int
	Misc string
}

func main() {
	/* open db */
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Printf("Failed to open db: %s\n", err)
		return
	}

	/* initialize DbMap */
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	defer dbmap.Db.Close()

	dbmap.AddTableWithName(sample{}, "samples").SetKeys(false, "Id") // don't use autoincrement

	/* create table */
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		return
	}

	/* insert */
	tx, err := dbmap.Begin()
	if err != nil {
		fmt.Printf("Failed to begin transaction: %s\n", err)
		return
	}

	sample1 := sample{
		Id:   1,
		Misc: "sample 1",
	}
	sample2 := sample{
		Id:   2,
		Misc: "sample 2",
	}
	samples := []sample{sample1, sample2}
	for _, sample := range samples {
		err = tx.Insert(&sample)
		if err != nil {
			fmt.Printf("Failed to insert: %s\nTry to rollback transaction.", err)
			err = tx.Rollback() // if an error occurs, try to rollback transction.
			if err != nil {
				fmt.Printf("Failed to rollback transaction: %s\n", err)
			}
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Failed to commit transaction: %s\n", err)
		return
	}

	/* select all */
	var res1 []sample
	_, err = dbmap.Select(&res1, "select * from samples")
	if err != nil {
		fmt.Printf("Failed to select: %s\n", err)
		return
	}
	for _, sample := range res1 {
		fmt.Printf("id: %d, misc: %s\n", sample.Id, sample.Misc)
	}

	/* select one */
	obj, err := dbmap.Get(sample{}, 1)
	if err != nil {
		fmt.Printf("Failed to select by id: %s\n", err)
	}
	sample := obj.(*sample)
	fmt.Printf("id: %d, misc: %s\n", sample.Id, sample.Misc)
}
