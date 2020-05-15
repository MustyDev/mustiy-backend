package model

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

type DanaMysql struct {
	DB *sql.DB
}

func NewDanaStoreMysql() DanaStore {
	godotenv.Load()
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true&clientFoundRows=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return &DanaMysql{DB: db}
}

func (store *DanaMysql) All() []Dana {
	danas := []Dana{}
	rows, err := store.DB.Query("SELECT * FROM dana")
	if err != nil {
		return danas
	}

	dana := Dana{}
	for rows.Next() {
		// rows.Scan(&siswa.ID, &siswa.Nama, &siswa.Nisn, &siswa.Pendidikan)
		danas = append(danas, dana)
	}

	return danas
}
