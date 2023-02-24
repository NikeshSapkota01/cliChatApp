package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/NikeshSapkota01/cliChatApp/db"
	"github.com/NikeshSapkota01/cliChatApp/pkg/user"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"

	"github.com/spf13/cobra"
)

var (
	host       string
	port       int
	namespaces []string
)

var (
	message string
)

func connectSocketIO(namespace string) {

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	green := "\033[32m"
	reset := "\033[0m"
	blue := "\033[34m"

	chats, err := user.GetAllChats(db)

	if err != nil {
		fmt.Printf("Failed to get chat: %v\n", err)
	}

	// Print all chats
	for _, chat := range chats {

		fmt.Printf("%sfrom: %s  message: %s%s\n", blue, chat.User, chat.Message, reset)
	}

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = "user"
	opts.Query["pwd"] = "pass"
	uri := "http://localhost:5000/socket.io/"

	client, err := socketio_client.NewClient(uri, opts)

	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}

	client.On("reply", func(msg string) {

		log.Printf("%s\nMessage: %s%s", green, msg, reset)
	})

	if err != nil {
		log.Printf("NewClient error: %s: %v\n", namespace, err)
		return
	}

	for {

		fmt.Print("Enter message: ")
		reader := bufio.NewReader(os.Stdin)
		message, _ = reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")

		usermail := &username
		user_id := &userId

		messageInfo := []string{*user_id, *usermail, message}
		parshedMessage, _ := json.Marshal(messageInfo)
		client.Emit("notice", string(parshedMessage))
	}
}

func runWebSocketCLI(cmd *cobra.Command, args []string) {

	for _, namespace := range namespaces {
		connectSocketIO(namespace)
	}

	select {}
}

var socketCmd = &cobra.Command{
	Use:   "socketio",
	Short: "Connect to a Socket.IO server",
	Long:  "Connect to a Socket.IO server and send/receive messages",
	Run:   runWebSocketCLI,
}

func init() {

	socketCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Socket.IO server hostname")
	socketCmd.Flags().IntVarP(&port, "port", "p", 8000, "Socket.IO server port")
	socketCmd.Flags().StringSliceVarP(&namespaces, "namespace", "n", []string{"/"}, "Socket.IO namespaces to connect to")

	rootCmd.AddCommand(socketCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
