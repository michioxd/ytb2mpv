package main

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

type MsgServerInfo struct {
	TYPE           string `json:"type"`
	MPV_STATUS     int    `json:"mpv_status"`
	YTDLP_STATUS   int    `json:"ytdlp_status"`
	SERVER_VERSION string `json:"server_version"`
	MPV_VERSION    string `json:"mpv_version"`
	YTDLP_VERSION  string `json:"yt-dlp_version"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendJsonMsg(conn *websocket.Conn, msg any) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

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
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

	}
}

func RunWSServer() {
	http.HandleFunc("/ytb2mpv", wsHandler)
	fmt.Println("Server started on :" + SERVER_PORT)
	http.ListenAndServe(":"+SERVER_PORT, nil)
}
