package main

import (
	"./handlers"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()

	e.Static("/", "public")
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/task", handlers.PostTask(db))
	e.PUT("/task", handlers.PutTask(db))
	e.DELETE("/task/:id", handlers.DeleteTask(db))

	// Start as a web server
	e.Logger.Fatal(e.Start(":1323"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			done INTEGER NOT NULL
		);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
