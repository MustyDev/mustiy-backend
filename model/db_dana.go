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
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true&clientFoundRows=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return &DanaMysql{DB: db}
}

func (store *DanaMysql) All() []Dana {
	danas := []Dana{}
	rows, err := store.DB.Query("SELECT * FROM donasi")
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
	result, err := store.DB.Exec(`INSERT INTO donasi(judul,kategori,nama,organisasi,email,nominal,deskripsi,waktu_start,waktu_end,url,status) VALUES(?,?,?,?,?,?,?,?,?,?,?)`, dana.Judul, dana.Kategori, dana.Nama, dana.Organisasi, dana.Email, dana.Nominal, dana.Deskripsi, dana.Waktu_start, dana.Waktu_end, dana.Url, dana.Status)
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

	err := store.DB.QueryRow(`SELECT * FROM donasi WHERE id=?`, id).Scan(
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

	err := store.DB.QueryRow(`SELECT * FROM donasi WHERE kategori=?`, kategori).Scan(
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
    UPDATE donasi SET judul= ?, kategori= ?, nama = ?, organisasi = ?,email = ?, nominal = ?, deskripsi = ?, waktu_start = ?,waktu_end = ?, url = ?  WHERE id = ?`,
		dana.Judul,
		dana.Kategori,
		dana.Nama,
		dana.Organisasi,
		dana.Email,
		dana.Nominal,
		dana.Deskripsi,
		dana.Waktu_start,
		dana.Waktu_end,
		dana.Url,
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

func (store *DanaMysql) Status(dana *Dana) error {
	result, err := store.DB.Exec(`
    UPDATE donasi SET status= ? WHERE id = ?`,
		dana.Status,
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
	result, err := store.DB.Exec(`DELETE FROM donasi WHERE id = ?`, dana.ID)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
