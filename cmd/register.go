package cmd

import (
	"fmt"
	"strings"

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

		for {
			fmt.Println("Trying to login into the system...")
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
