package data

import (
	"KosKita/features/chat"

	"gorm.io/gorm"
)

type chatQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) chat.ChatDataInterface {
	return &chatQuery{
		db: db,
	}
}

// GetRoom implements chat.ChatDataInterface.
func (repo *chatQuery) GetRoom(userIdLogin int) ([]chat.Core, error) {
	var rooms []Chat
	tx := repo.db.Where("receiver_id = ? OR sender_id = ?", userIdLogin, userIdLogin).Preload("UserReceiver").Preload("UserSender").Find(&rooms)
	if tx.Error != nil {
		return nil, tx.Error
	}
	roomMap := make(map[string]chat.Core)
	for _, room := range rooms {
		roomMap[room.RoomID] = room.ModelToCoreRoom(uint(userIdLogin))
	}
	var result []chat.Core
	for _, room := range roomMap {
		result = append(result, room)
	}
	return result, nil
}

// CreateRoom implements chat.ChatDataInterface.
func (repo *chatQuery) CreateRoom(roomID string, receiverID int, senderID int) error {
	room := Chat{
		RoomID:     roomID,
		ReceiverID: uint(receiverID),
		SenderID:   uint(senderID),
	}

	tx := repo.db.Create(&room)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// CreateRoom implements chat.ChatDataInterface.
func (repo *chatQuery) CreateMessage(receiverID int, senderID int, input chat.Core) (chat.Core, error) {
	chatInput := CoreToModelChat(input)
	chatInput.ReceiverID = uint(receiverID)
	chatInput.SenderID = uint(senderID)

	tx := repo.db.Create(&chatInput)
	if tx.Error != nil {
		return chat.Core{}, tx.Error
	}

	return chat.Core{
		Message:    chatInput.Message,
		RoomID:     chatInput.RoomID,
		ReceiverID: chatInput.ReceiverID,
		SenderID:   chatInput.SenderID,
	}, nil
}

// GetMessage implements chat.ChatDataInterface.
func (repo *chatQuery) GetMessage(roomId string) ([]chat.Core, error) {
	var chats []Chat
	tx := repo.db.Where("room_id = ?", roomId).Order("created_at desc").Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var cores []chat.Core
	for _, c := range chats {
		cores = append(cores, c.ModelToCoreChat())
	}

	return cores, nil
}
