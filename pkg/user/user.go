package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"github.com/NikeshSapkota01/cliChatApp/db"
	"github.com/NikeshSapkota01/cliChatApp/mystructs"
	"github.com/NikeshSapkota01/cliChatApp/util"
)

func CreateUser(db *db.Database, u mystructs.UserInfo) error {
	id := uuid.New()

	hashed_pass, err := util.HashPassword(u.Password)

	if err != nil {
		return err
	}

	query := "INSERT INTO users (id, username, email, hashed_password) VALUES ($1, $2, $3, $4)"
	_, err = db.GetDB().Exec(query, id, u.Username, u.Email, hashed_pass)

	if err != nil {
		return fmt.Errorf("failed to insert user into database: %v", err)
	}

	fmt.Printf("User created successfully... \n")

	return nil
}

func IdentifyUser(db *db.Database, username string, password string) (*mystructs.User, error) {
	var user mystructs.User

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

	fmt.Printf("User logged in successfully... \n")
	return &user, nil
}

func UserChats(db *db.Database, userId string, message string) error {
	id := uuid.New()

	query := "INSERT INTO chats (id, user_id, message) VALUES ($1, $2, $3)"
	_, err := db.GetDB().Exec(query, id, userId, message)

	if err != nil {
		return fmt.Errorf("failed to insert user into database: %v", err)
	}

	fmt.Println("Chat added successfully...")

	return nil
}

func GetAllChats(db *db.Database) ([]mystructs.Chat, error) {
	var chats []mystructs.Chat

	query := "SELECT users.username, chats.message FROM chats JOIN users ON chats.user_id = users.id"
	rows, err := db.GetDB().Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to get chats: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chat mystructs.Chat
		if err := rows.Scan(&chat.User, &chat.Message); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %v", err)
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating chats: %v", err)
	}

	fmt.Printf("chats information %v\n", chats)
	fmt.Println("Chats retrieved successfully...")
	return chats, nil
}
