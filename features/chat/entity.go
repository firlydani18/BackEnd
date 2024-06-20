package chat

import (
	"KosKita/features/user"
	"time"
)

type Core struct {
	ID         uint
	Message    string
	RoomID     string
	ReceiverID uint
	SenderID   uint
	User       user.Core
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// interface untuk Data Layer
type ChatDataInterface interface {
	CreateRoom(roomID string, receiverID int, senderID int) error
	CreateMessage(receiverID int, senderID int, input Core) (Core, error)
	GetMessage(roomId string) ([]Core, error)
	GetRoom(userIdlogin int) ([]Core, error)
}

// interface untuk Service Layer
type ChatServiceInterface interface {
	CreateRoom(roomID string, receiverID int, senderID int) error
	CreateChat(receiverID int, senderID int, input Core) (Core, error)
	GetMessage(roomId string) ([]Core, error)
	GetRoom(userIdlogin int) ([]Core, error)
}