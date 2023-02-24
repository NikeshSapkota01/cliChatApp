package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	username   string
	email      string
	password   string
	isLoggedIn bool
	message    string
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "cliChatApp",
		Long:  `Create a CLI chat application using golang.`,

		Run: func(cmd *cobra.Command, args []string) {
			for {

				if isLoggedIn {
					break
				}

				fmt.Println("Create a login or registration form")
				var choice int
				for {
					fmt.Println("1: Login")
					fmt.Println("2: Register")
					fmt.Print("Enter your choice: ")
					if _, err := fmt.Scanln(&choice); err != nil {
						fmt.Println("Invalid input. Please enter a number.")
						continue
					}
					if choice != 1 && choice != 2 {
						fmt.Println("Invalid choice. Please enter 1 or 2.")
						continue
					}
					break
				}
				if choice == 1 {
					fmt.Println("Login process will now begin...")

					loginCmd.Run(cmd, args)
				} else {
					fmt.Println("Registration process will now begin...")

					registerCmd.Run(cmd, args)
				}
			}

			if isLoggedIn {
				fmt.Println("Chat app will now begin...")

				socketCmd.Run(cmd, args)
			}

			username = ""
			email = ""
			password = ""

		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
