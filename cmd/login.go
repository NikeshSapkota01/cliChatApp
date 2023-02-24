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

var loginCmd = &cobra.Command{
	Use:   "Login",
	Short: "Create a new Login",
	Run: func(cmd *cobra.Command, args []string) {
		username = ""
		email = ""
		password = ""

		for {

			if username == "" {
				fmt.Print("Enter username: ")
				fmt.Scanln(&username)
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

		db, err := db.NewDatabase()

		if err != nil {
			fmt.Printf("Failed to connect to database: %v\n", err)
			return
		}
		defer db.Close()
		user, err := user.IdentifyUser(db, username, password)

		if err != nil {
			fmt.Printf("Failed to create user: %v\n", err)
			return
		}

		for {
			fmt.Println("Trying to login into the system...")
			fmt.Printf("Email: %s\n", user)
			fmt.Printf("Password: %s\n", strings.Repeat("*", len(password)))
			break
		}

		if user != nil {
			isLoggedIn = true
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
