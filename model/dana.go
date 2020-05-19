package model

type Dana struct {
	ID, Kategori, Nominal, Jumlah                          int
	Nama, Judul, Status, Url, Organisasi, Deskripsi, Email string
	Waktu_start, Waktu_end                                 string
}

func CreateDana(judul string, kategori int, nama, organisasi, email string, nominal int, deskripsi string, waktu_start, waktu_end, url, status string, jumlah int) (*Dana, error) {
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
		Jumlah:      jumlah,
	}, nil
}
