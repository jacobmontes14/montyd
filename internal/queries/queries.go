package queries

const (
	SelectUserByID = "SELECT key, value FROM data WHERE key = ?"
	DeleteUserByID = "DELETE FROM data WHERE key = ?"
	InsertProduct  = "INSERT INTO data (value) VALUES (?)"
	CreateTable    = "CREATE TABLE IF NOT EXISTS data (key INTEGER PRIMARY KEY, value TEXT)"
)
