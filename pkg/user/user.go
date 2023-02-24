package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"github.com/NikeshSapkota01/cliChatApp/db"
	"github.com/NikeshSapkota01/cliChatApp/util"
)

type User struct {
	Id       string
	Username string
	Password string
}

func CreateUser(db *db.Database, username string, email string, password string) error {
	id := uuid.New()

	hashed_pass, err := util.HashPassword(password)

	if err != nil {
		return err
	}

	query := "INSERT INTO users (id, username, email, hashed_password) VALUES ($1, $2, $3, $4)"
	_, err = db.GetDB().Exec(query, id, username, email, hashed_pass)

	if err != nil {
		return fmt.Errorf("failed to insert user into database: %v", err)
	}

	fmt.Println("User created successfully...")

	return nil
}

func IdentifyUser(db *db.Database, username string, password string) (*User, error) {
	var user User

	query := "SELECT id, username, hashed_password FROM users WHERE username=$1"
	err := db.GetDB().QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	err = util.CheckPassword(password, user.Password)

	if err != nil {
		return nil, err
	}

	fmt.Print("User logged in successfully...")
	return &user, nil
}
