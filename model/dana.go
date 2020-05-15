package model

import "time"

type Dana struct {
	ID, Kategori, Nominal                                  int
	Nama, Judul, Status, Url, Organisasi, Deskripsi, Email string
	Waktu_start, Waktu_end                                 time.Time
}

func CreateSiswa(judul string, kategori int, nama, organisasi, email string, nominal int, deskripsi string, waktu_start, waktu_end time.Time, url, status string) (*Dana, error) {
	return &Dana{
		Judul:       judul,
		Kategori:    kategori,
		Nama:        nama,
		Organisasi:  organisasi,
		Email:       email,
		Nominal:     nominal,
		Deskripsi:   deskripsi,
		Waktu_start: waktu_start,
		Waktu_end:   waktu_end,
		Url:         url,
		Status:      status,
	}, nil
}
