package sqlite

import (
	"database/sql"
	"log"

	"github.com/YerzatCode/learing_kz_backend/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

func InitSqlite(url string) *sql.DB {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		log.Fatal("failed to connect database: %w", err)
	}
	defer db.Close()

	createSuperAdmin(db)

	return db
}

func createSuperAdmin(db *sql.DB) {
	password, err := utils.HashPassword("password")
	if err != nil {
		log.Fatal("failed to hash password: %w", err)
	}
	result, err := db.Exec("insert into User (name, email,password,role) values ('super', $1, $2, $3)",
		"admin@example.com", password, "super")
	if err != nil {
		panic(err)
	}
	_ = result
}
