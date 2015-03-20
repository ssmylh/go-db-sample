package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ssmylh/go-db-sample/argen/gen"
)

// Preparation:
// - cd gen
// - argen sample.go
func main() {
	/* open db */
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Printf("Failed to open db: %s\n", err)
		return
	}
	defer db.Close()

	/* create table */
	_, err = db.Exec("create table samples (id integer not null primary key, misc text)")
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		return
	}

	gen.Use(db)

	/* insert */
	/* In the present(2015/03/20), argen does not support transaction */
	sample1 := &gen.Sample{Misc: "sample 1"}
	_, arErr := sample1.Save()
	if arErr != nil {
		fmt.Println("Failed to insert record")
		for k, es := range arErr.Messages {
			for _, e := range es {
				fmt.Printf("%s: %s\n", k, e)
			}
		}
		return
	}

	sample2 := &gen.Sample{Misc: "sample 2"}
	_, arErr = sample2.Save()
	if arErr != nil {
		fmt.Println("Failed to insert record")
		for k, es := range arErr.Messages {
			for _, e := range es {
				fmt.Printf("%s: %s\n", k, e)
			}
		}
		return
	}

	/* select all */
	samples, err := gen.Sample{}.All().Query()
	if err != nil {
		fmt.Printf("Failed to select all: %s\n", err)
		return
	}
	for _, sample := range samples {
		fmt.Printf("id: %d, misc: %s\n", sample.Id, sample.Misc)
	}

	/* select one */
	sample, err := gen.Sample{}.Find(sample1.Id)
	if err != nil {
		fmt.Printf("Failed to select one: %s\n", err)
		return
	}
	fmt.Printf("id: %d, misc: %s\n", sample.Id, sample.Misc)
}
