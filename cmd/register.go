package cmd

import (
	"fmt"
	"strings"

	"github.com/NikeshSapkota01/cliChatApp/db"
	"github.com/NikeshSapkota01/cliChatApp/pkg/user"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	username string
	email    string
	password string
)

var registerCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a new service",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			if username == "" {
				fmt.Print("Enter username: ")
				fmt.Scanln(&username)
				continue
			}
			if email == "" {
				fmt.Print("Enter email: ")
				fmt.Scanln(&email)
				continue
			}
			if password == "" {
				fmt.Print("Enter password: ")
				passwordBytes, err := terminal.ReadPassword(0)
				if err != nil {
					fmt.Println(err)
					return
				}
				password = string(passwordBytes)
				fmt.Println()
				continue
			}
			break
		}

		// create a new user in the database
		db, err := db.NewDatabase()
		if err != nil {
			fmt.Printf("Failed to connect to database: %v\n", err)
			return
		}
		defer db.Close()
		if err := user.CreateUser(db, username, email, password); err != nil {
			fmt.Printf("Failed to create user: %v\n", err)
			return
		}

		for {
			fmt.Println("Created user with...")

			fmt.Printf("Username: %s\n", username)
			fmt.Printf("Email: %s\n", email)
			fmt.Printf("Password: %s\n", strings.Repeat("*", len(password)))
			break
		}
	},
}

func init() {

	rootCmd.AddCommand(registerCmd)
}
