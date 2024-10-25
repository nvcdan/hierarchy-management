package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")
	loc := os.Getenv("DB_LOC")

	if charset == "" {
		charset = "utf8mb4"
	}

	if loc == "" {
		loc = "Local"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		user,
		pass,
		host,
		port,
		name,
		charset,
		loc,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for {
		err = db.Ping()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	fmt.Println("Connected to the database successfully!")
	return db, nil
}
