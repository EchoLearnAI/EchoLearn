package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Message struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (customize this in production)
	},
}

func initDB() {
	var err error
	connStr := "postgres://postgres:postgres@db:5432/chatdb?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error verifying database connection: %v", err)
	}
	fmt.Println("Connected to the database!")
}

func fetchMessagesFromDB() ([]Message, error) {
	rows, err := db.Query("SELECT id, username, message, timestamp FROM messages ORDER BY timestamp ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Username, &msg.Message, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	// Extract the username query parameter
	username := r.URL.Query().Get("username")

	// Query messages specific to the username
	var rows *sql.Rows
	var err error
	if username != "" {
		rows, err = db.Query("SELECT id, username, message, timestamp FROM messages WHERE username = $1 ORDER BY timestamp ASC", username)
	} else {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Username, &msg.Message, &msg.Timestamp)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		messages = append(messages, msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		msg.Timestamp = time.Now()

		// Save the message to the database
		if err := saveMessageToDB(msg); err != nil {
			log.Println("Failed to save message:", err)
			continue
		}

		// Echo the message back to the client
		conn.WriteJSON(msg)
	}
}

func saveMessageToDB(msg Message) error {
	query := "INSERT INTO messages (username, message, timestamp) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, msg.Username, msg.Message, msg.Timestamp)
	return err
}

func main() {
	initDB()

	http.HandleFunc("/messages", handleMessages) // Endpoint to fetch all messages
	http.HandleFunc("/ws", handleWebSocket)      // WebSocket for real-time messaging

	fmt.Println("Server running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
