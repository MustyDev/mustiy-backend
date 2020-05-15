package model

import (
	"database/sql"
	"log"
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
		rows.Scan(
			&dana.ID,
			&dana.Judul,
			&dana.Kategori,
			&dana.Nama,
			&dana.Organisasi,
			&dana.Email,
			&dana.Nominal,
			&dana.Deskripsi,
			&dana.Waktu_start,
			&dana.Waktu_end,
			&dana.Url,
			&dana.Status,
		)
		danas = append(danas, dana)
	}

	return danas
}

func (store *DanaMysql) Save(dana *Dana) error {
	result, err := store.DB.Exec(`INSERT INTO dana(judul,kategori,nama,organisasi,email,nominal,deskripsi,waktu_start,waktu_end,url,status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`, dana.Judul, dana.Kategori, dana.Nama, dana.Organisasi, dana.Email, dana.Nominal, dana.Deskripsi, dana.Waktu_start, dana.Waktu_end, dana.Url, dana.Status)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	// new id
	lastID, err := result.LastInsertId()

	dana.ID = int(lastID)

	return nil
}

func (store *DanaMysql) Find(id int) *Dana {
	dana := Dana{}

	err := store.DB.QueryRow(`SELECT * FROM dana WHERE id=?`, id).Scan(
		&dana.ID,
		&dana.Judul,
		&dana.Kategori,
		&dana.Nama,
		&dana.Organisasi,
		&dana.Email,
		&dana.Nominal,
		&dana.Deskripsi,
		&dana.Waktu_start,
		&dana.Waktu_end,
		&dana.Url,
		&dana.Status,
	)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &dana
}

func (store *DanaMysql) Found(kategori int) *Dana {
	dana := Dana{}

	err := store.DB.QueryRow(`SELECT * FROM dana WHERE kategori=?`, kategori).Scan(
		&dana.ID,
		&dana.Judul,
		&dana.Kategori,
		&dana.Nama,
		&dana.Organisasi,
		&dana.Email,
		&dana.Nominal,
		&dana.Deskripsi,
		&dana.Waktu_start,
		&dana.Waktu_end,
		&dana.Url,
		&dana.Status,
	)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &dana
}

func (store *DanaMysql) Update(dana *Dana) error {
	result, err := store.DB.Exec(`
    UPDATE dana SET judul= ?, kategori= ?, nama = ?, organisasi = ?,email = ?, nominal = ?, deskripsi = ?, waktu_start = ?,waktu_end = ? WHERE id = ?`,
		dana.Judul,
		dana.Kategori,
		dana.Nama,
		dana.Organisasi,
		dana.Email,
		dana.Nominal,
		dana.Deskripsi,
		dana.Waktu_start,
		dana.Waktu_end,
		dana.ID,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (store *DanaMysql) Delete(dana *Dana) error {
	result, err := store.DB.Exec(`DELETE FROM dana WHERE id = ?`, dana.ID)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
