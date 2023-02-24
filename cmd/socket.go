package cmd

import (
	"fmt"
	"log"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"

	"github.com/spf13/cobra"
)

var (
	host       string
	port       int
	namespaces []string
)

func connectSocketIO(namespace string) {

	opts := &socketio_client.Options{
		Transport: "websocket",
	}
	client, err := socketio_client.NewClient(fmt.Sprintf("http://localhost:5000/%s", namespace), opts)

	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}

	client.On("connect", func() {
		log.Printf("Connected to WebSocket namespace %s", namespace)
	})

	client.On("disconnection", func() {
		log.Printf("Disconnected from WebSocket namespace %s", namespace)
	})

	client.On("message", func(msg string) {
		log.Printf("Received message from WebSocket namespace %s: %s", namespace, msg)
	})

	if err != nil {
		log.Printf("NewClient error: %s: %v\n", namespace, err)
		return
	}

	// Send a message to all connected clients
	if err := client.Emit("message", "Hello, World!"); err != nil {
		log.Printf("Failed to send message to WebSocket namespace %s: %v", namespace, err)
		return
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
