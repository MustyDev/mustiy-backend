package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/EnggarSe/mustiy-backend/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func appUsers(e *echo.Echo, storeUsers model.UserStore) {

	// curl http://localhost:9001/users
	e.GET("/users", func(c echo.Context) error {
		// Process
		users := storeUsers.All()

		// Response
		return c.JSON(http.StatusOK, users)
	})

	// curl -d "Username=EnggarSe &Email=enggarseptrinas@yahoo.com &Phone=0812670053234 &Password=asdzxc" http://localhost:9001/users
	e.POST("/users", func(c echo.Context) error {
		// Given
		var role int
		name := c.FormValue("Username")
		email := c.FormValue("Email")
		phone := c.FormValue("Phone")
		password := c.FormValue("Password")
		token := "Belum Diketahui"
		if phone != "" {
			role = 1
		} else {
			role = 2
		}
		//Hashing password
		passwordHash, _ := model.Hash(password)

		// Create instabce
		user, _ := model.CreateUser(name, email, phone, passwordHash, role, token)

		// Persist
		storeUsers.Save(user)

		// Response
		return c.JSON(http.StatusOK, user)
	})

	//curl -X DELETE http://localhost:9001/users/19

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		user := storeUsers.Find(id)

		storeUsers.Delete(user)

		return c.JSON(http.StatusOK, user)
	})

	// curl -X PUT -d "Username=GentaKamsa &Email=enggarseptrinas@yahoo.com &Phone=0812670053234 &Password=asdzxc" http://localhost:9001/users/22

	e.PUT("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		user := storeUsers.Find(id)
		user.Username = c.FormValue("Username")
		user.Email = c.FormValue("Email")
		user.Phone = c.FormValue("Phone")
		password := c.FormValue("password")

		user.Password, _ = model.Hash(password)

		storeUsers.Update(user)

		return c.JSON(http.StatusOK, user)
	})

}

func app(e *echo.Echo, store model.DanaStore) {
	e.GET("/donasi", func(c echo.Context) error {
		danas := store.All()
		return c.JSON(http.StatusOK, danas)
	})

	e.GET("/donasi/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		danas := store.Find(id)
		return c.JSON(http.StatusOK, danas)
	})

	e.GET("/donasi/kategori/:kategori", func(c echo.Context) error {

		kategori, _ := strconv.Atoi(c.Param("kategori"))
		danas := store.Found(kategori)
		return c.JSON(http.StatusOK, danas)
	})

	e.POST("/donasi", func(c echo.Context) error {
		judul := c.FormValue("judul")
		kategori, _ := strconv.Atoi(c.Param("kategori"))
		nama := c.FormValue("nama")
		organisasi := c.FormValue("organisasi")
		email := c.FormValue("email")
		nominal, _ := strconv.Atoi(c.Param("nominal"))
		deskripsi := c.FormValue("deskripsi")
		waktu_start := c.FormValue("waktu_start")
		waktu_end := c.FormValue("waktu_end")
		url := c.FormValue("url")
		status := c.FormValue("status")

		danas, _ := model.CreateDana(judul, kategori, nama, organisasi, email, nominal, deskripsi, waktu_start, waktu_end, url, status)
		store.Save(danas)

		return c.JSON(http.StatusOK, danas)
	})

	e.PUT("/donasi/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		dana := store.Find(id)
		dana.Judul = c.FormValue("judul")
		dana.Kategori, _ = strconv.Atoi(c.Param("kategori"))
		dana.Nama = c.FormValue("nama")
		dana.Organisasi = c.FormValue("organisasi")
		dana.Email = c.FormValue("email")
		dana.Nominal, _ = strconv.Atoi(c.Param("nominal"))
		dana.Deskripsi = c.FormValue("deskripsi")
		dana.Waktu_start = c.FormValue("waktu_start")
		dana.Waktu_end = c.FormValue("waktu_end")
		dana.Url = c.FormValue("url")
		dana.Status = c.FormValue("status")

		store.Update(dana)

		return c.JSON(http.StatusOK, dana)
	})

	e.DELETE("/donasi/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		dana := store.Find(id)
		store.Delete(dana)
		return c.JSON(http.StatusOK, dana)
	})

}

func main() {
	godotenv.Load()
	var storeUsers model.UserStore
	storeUsers = model.NewUserMySQL()
	store := model.NewDanaStoreMysql()
	e := echo.New()
	appUsers(e, storeUsers)
	app(e, store)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
