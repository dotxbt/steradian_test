package config

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "steradian.db")
	if err != nil {
		fmt.Println("error open DB")
	}

	schema := `
	 CREATE TABLE IF NOT EXISTS cars (
	 	car_id INTEGER PRIMARY KEY AUTOINCREMENT,
		car_name TEXT NOT NULL,
		day_rate REAL NOT NULL,
		month_rate REAL NOT NULL,
		image TEXT NOT NULL
	 );

	 CREATE TABLE IF NOT EXISTS orders (
	 	order_id INTEGER PRIMARY KEY AUTOINCREMENT,
		car_id INTEGER NOT NULL,
		order_date TEXT DEFAULT (datetime('now')),
		pickup_date TEXT NOT NULL,
		dropoff_date TEXT NOT NULL,
		pickup_location TEXT NOT NULL,
		dropoff_location TEXT NOT NULL,
		FOREIGN KEY(car_id) REFERENCES cars(car_id) ON DELETE CASCADE
	 );
	`
	_, er := db.Exec(schema)
	if er != nil {
		panic(fiber.NewError(fiber.StatusInternalServerError, "Upd! something happen with my service"))
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(fiber.NewError(fiber.StatusInternalServerError, "Upd! something happen with my service"))
	}
	return db
}
