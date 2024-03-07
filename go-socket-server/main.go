package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"math/rand"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

type Message struct {
	ID int 			`json:"id"`
	Time int64		`json:"time"`
	Data string		`json:"data"`

}

func handleWebSocket(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected")

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		log.Println("Received message: %s\n", p)
		if err := conn.WriteMessage(messageType,p); err != nil {
			log.Println(err)
			return
		}
	}

}
var idCounter = 1
func sendDataToClient(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	for{
		
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 99
		randomInt := rand.Intn(100)

		message := Message {
			ID: idCounter,
			Time: int64(randomInt),
			Data: "Xin chao",
		}
		idCounter++
		fmt.Println(randomInt)
	 // Marshal the Message struct to JSON
		jsonData, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		// Send the message to the client
		if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Println(err)
			return
		}
	}
}


func main(){
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/sendData", sendDataToClient)
	log.Fatal(http.ListenAndServe(":8000", nil))
}