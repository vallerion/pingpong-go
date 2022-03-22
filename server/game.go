package server

import "github.com/google/uuid"

type roomStatus int

const (
	PENDING roomStatus = iota
	RUNNING
	COMPLETED
)

type Game struct {
	//Rooms map[uuid.UUID]*Room
	Status                  roomStatus
	LeftPlayer, RightPlayer *Player
}

type Room struct {
	status                  roomStatus
	leftPlayer, rightPlayer *Player
}

type Player struct {
	Id   uuid.UUID
	X, Y int64
}
