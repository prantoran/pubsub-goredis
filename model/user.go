package model

import "github.com/gorilla/websocket"

type User struct {
	ID   string
	conn *websocket.Conn
}
