package sqlite

import (
	"database/sql"

	queries "github.com/jacobmontes14/montyd/internal/queries"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	db *sql.DB
}

func InitializeDb(dbName string) *Db {
	d, _ := sql.Open("sqlite3", dbName)

	return &Db{db: d}
}

func (d *Db) PrepareDb() {
	statement, _ := d.db.Prepare(queries.CreateTable)
	statement.Exec()
}

func (d *Db) AddToDb(value string) {
	statement, _ := d.db.Prepare(queries.InsertProduct)
	statement.Exec(value)
}

func (d *Db) DeleteFromDb(id int) {
	statement, _ := d.db.Prepare(queries.DeleteUserByID)
	statement.Exec(id)
}

func (d *Db) SelectFromDb(id int) *sql.Row {
	statement := d.db.QueryRow(queries.SelectUserByID, id)
	return statement
}
