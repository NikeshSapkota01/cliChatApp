package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var loginCmd = &cobra.Command{
	Use:   "Login",
	Short: "Create a new Login",
	Run: func(cmd *cobra.Command, args []string) {

		for {
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
			fmt.Printf("Email: %s\n", email)
			fmt.Printf("Password: %s\n", strings.Repeat("*", len(password)))
			break
		}

		// Your code to create a new Login goes here
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
