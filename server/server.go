package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

type MsgServerInfo struct {
	TYPE           string `json:"type"`
	MPV_STATUS     int    `json:"mpv_status"`
	YTDLP_STATUS   int    `json:"ytdlp_status"`
	SERVER_VERSION string `json:"server_version"`
	MPV_VERSION    string `json:"mpv_version"`
	YTDLP_VERSION  string `json:"ytdlp_version"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendJsonMsg(conn *websocket.Conn, msg any) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	SendJsonMsg(conn, MsgServerInfo{
		TYPE:           "server_info",
		MPV_STATUS:     MPV_STATUS,
		YTDLP_STATUS:   YTDLP_STATUS,
		SERVER_VERSION: VERSION,
		MPV_VERSION:    MPV_VERSION,
		YTDLP_VERSION:  YTDLP_VERSION,
	})

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// parse message to JSON
		var msg map[string]interface{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			continue
		}

		msgType, ok := msg["type"].(string)

		if !ok {
			log.Println("Invalid message format")
			continue
		}

		switch msgType {
		case "shutdown":
			log.Println("Byebye")
			os.Exit(0)
		}

	}
}

func RunWSServer() {
	http.HandleFunc("/ytb2mpv", wsHandler)
	http.ListenAndServe("127.0.0.1:"+SERVER_PORT, nil)
}
