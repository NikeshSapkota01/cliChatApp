package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NikeshSapkota01/cliChatApp/db"
	"github.com/NikeshSapkota01/cliChatApp/pkg/user"

	socketio "github.com/googollee/go-socket.io"
)

var numClients int

func main() {
	server := socketio.NewServer(nil)

	type Database struct {
		emailName string
		message   string
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		// fmt.Println("connected:", s.ID())
		numClients++

		fmt.Println("connected:", numClients)
		s.Join("bcast")

		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		var messageInfo []string
		err := json.Unmarshal([]byte(msg), &messageInfo)
		if err != nil {
			fmt.Println("Error parsing message:", err)
			return
		}

		if len(messageInfo) < 3 {
			fmt.Println("Invalid message format")
			return
		}

		userId := messageInfo[0]
		email := messageInfo[1]
		message := messageInfo[2]

		newMessage := map[string]interface{}{
			"from":    email,
			"message": message,
		}

		parshedMessage, _ := json.Marshal(newMessage)
		fmt.Println(string(parshedMessage))

		db, err := db.NewDatabase()

		if err != nil {
			fmt.Printf("Failed to connect to database: %v\n", err)
			return
		}
		defer db.Close()

		if err := user.UserChats(db, userId, message); err != nil {
			fmt.Printf("Failed to add chat: %v\n", err)
			return
		}

		server.BroadcastToRoom("/", "bcast", "reply", string(parshedMessage))
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		numClients--

	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
