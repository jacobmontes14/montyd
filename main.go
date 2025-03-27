package main

import (
	"fmt"
	"strconv"

	sqlite "github.com/jacobmontes14/montyd/internal/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlite.InitializeDb("./datastore")
	//db.PrepareDb()

	//db.AddToDb("Golang")

	row := db.SelectFromDb(1)
	var id int
	var value string

	row.Scan(&id, &value)
	fmt.Println(strconv.Itoa(id) + ": " + value)
}
