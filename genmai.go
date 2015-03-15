package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

type sample struct {
	Id   int `db:"pk"`
	Misc string
}

// This implements `genmai.TableNamer` interface, in order to specify table name.
func (s *sample) TableName() string {
	return "samples"
}

func main() {
	/* open db */
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		fmt.Printf("Failed to open db: %s\n", err)
		return
	}
	defer db.Close()

	/* create table */
	if err = db.CreateTable(&sample{}); err != nil {
		fmt.Printf("Failed to create table", err)
	}

	/* insert */
	err = db.Begin()
	if err != nil {
		fmt.Printf("Failed to begin transaction: %s\n", err)
		return
	}

	sample1 := &sample{
		Id:   1,
		Misc: "sample 1",
	}
	sample2 := &sample{
		Id:   2,
		Misc: "sample 2",
	}
	samples := []*sample{sample1, sample2}

	_, err = db.Insert(samples)
	if err != nil {
		fmt.Printf("Failed to insert: %s\nTry to rollback transaction.", err)
		err = db.Rollback() // if an error occurs, try to rollback transction.
		if err != nil {
			fmt.Printf("Failed to rollback transaction: %s\n", err)
		}
		return
	}

	err = db.Commit()
	if err != nil {
		fmt.Printf("Failed to commit transaction: %s\n", err)
		return
	}

	/* select all */
	var res1 []sample
	err = db.Select(&res1) // must pass pointer to Select method
	if err != nil {
		fmt.Printf("Failed to select: %s\n", err)
		return
	}
	for _, sample := range res1 {
		fmt.Printf("id: %d, misc: %s\n", sample.Id, sample.Misc)
	}

	/* select one */
	var res2 []sample
	err = db.Select(&res2, db.Where("id", "=", 1))
	if err != nil {
		fmt.Printf("Failed to select by id: %s\n", err)
		return
	}
	for _, sample := range res2 {
		fmt.Printf("id: %d, misc: %s\n", sample.Id, sample.Misc)
	}
}
