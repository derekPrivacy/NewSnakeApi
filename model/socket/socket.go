package model

import (
	"sync"

	"github.com/gorilla/websocket"
)

type CReq struct {
	Type string

	// for addPlayer
	Data   string
	RoomID int

	// for updateAvatar
	AvatarID  int
	PositionX float64
	PositionY float64

	BodyLength int
	Direction  string
	//
	FoodPositionX float64
	FoodPositionY float64
}

type CRes struct {
	Room Room
}

type Connection struct {
	ID         string
	Player     string
	Connection *websocket.Conn
	Mu         sync.Mutex
}

type Room struct {
	ID            int
	ConnectionArr []Connection
	Data          []string
	Avatar        []Avatar
	Food          Food
}

// array for data
var RoomArr []Room

var Players []string
