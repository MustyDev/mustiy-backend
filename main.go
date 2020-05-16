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

}

func main() {
	godotenv.Load()
	var storeUsers model.UserStore
	storeUsers = model.NewUserMySQL()
	e := echo.New()
	appUsers(e, storeUsers)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
