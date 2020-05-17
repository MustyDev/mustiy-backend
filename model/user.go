package model

import (
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	All() []User
	Save(*User) error
	Find(int) *User
	Update(*User) error
	Delete(user *User) error
	Login(string) *User
	FindEmail(string) *User
}
type UserStoreMySQL struct {
	DB *sql.DB
}
type User struct {
	ID       int
	Username string
	Email    string
	Phone    string
	Password string
	Role     int
	Token    string
}

func NewUserMySQL() UserStore {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true&clientFoundRows=true"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	return &UserStoreMySQL{DB: db}
}

func (store *UserStoreMySQL) All() []User {
	users := []User{}
	rows, err := store.DB.Query("SELECT * FROM users")

	if err != nil {
		return users
	}
	user := User{}
	for rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Token,
		)
		users = append(users, user)
	}
	return users
}
func CreateUser(username, email, phone, password string, role int, token string) (*User, error) {
	return &User{
		Username: username,
		Email:    email,
		Phone:    phone,
		Password: password,
		Role:     role,
		Token:    token,
	}, nil
}
func (store *UserStoreMySQL) Save(user *User) error {
	result, err := store.DB.Exec(`
		INSERT INTO users(Username, Email, Phone, Password, Role, Token) VALUES(?,?,?,?,?,?)`,
		user.Username,
		user.Email,
		user.Phone,
		user.Password,
		user.Role,
		user.Token,
	)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return err
	}
	user.ID = int(lastID)

	return nil
}

func (store *UserStoreMySQL) Find(id int) *User {
	user := User{}

	err := store.DB.
		QueryRow(`SELECT * FROM users WHERE id=?`, id).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Token,
		)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &user
}

func (store *UserStoreMySQL) Update(user *User) error {
	result, err := store.DB.Exec(`
		UPDATE users SET Username= ?, Email = ?, Phone = ?, Password = ?, Role = ? WHERE id =?`,
		user.Username,
		user.Email,
		user.Phone,
		user.Password,
		user.Role,
		user.ID,
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

func (store *UserStoreMySQL) Delete(user *User) error {
	result, err := store.DB.Exec(`
	DELETE FROM users WHERE id = ?`,
		user.ID,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil
	}
	return nil
}

func (store *UserStoreMySQL) Login(email string) *User {
	user := User{}

	err := store.DB.
		QueryRow(`SELECT * FROM users WHERE Email=?`, email).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Token,
		)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &user
}

func (store *UserStoreMySQL) FindEmail(email string) *User {
	user := User{}

	err := store.DB.
		QueryRow(`SELECT * FROM users WHERE Email=?`, email).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Token,
		)

	if err != nil {
		return nil
	}

	return &user
}

func Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hashed), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
