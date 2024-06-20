package service

import (
	"KosKita/features/chat"
)

type chatService struct {
	chatData chat.ChatDataInterface
}

func New(repo chat.ChatDataInterface) chat.ChatServiceInterface {
	return &chatService{
		chatData: repo,
	}
}

// GetRoom implements chat.ChatServiceInterface.
func (cs *chatService) GetRoom(userIdLogin int) ([]chat.Core, error) {
	rooms, err := cs.chatData.GetRoom(userIdLogin)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}


// CreateRoom implements chat.ChatServiceInterface.
func (cs *chatService) CreateRoom(roomID string, receiverID int, senderID int) error {
	return cs.chatData.CreateRoom(roomID, receiverID, senderID)
}

// CreateRoom implements chat.ChatServiceInterface.
func (cs *chatService) CreateChat(receiverID int, senderID int, input chat.Core) (chat.Core, error) {
	chatOutput, errInsert := cs.chatData.CreateMessage(receiverID, senderID, input)
	if errInsert != nil {
		return chat.Core{}, errInsert
	}

	return chatOutput, nil
}

// GetMessage implements chat.ChatServiceInterface.
func (cs *chatService) GetMessage(roomId string) ([]chat.Core, error) {
	chats, errGet := cs.chatData.GetMessage(roomId)
	if errGet != nil {
		return nil, errGet
	}

	return chats, nil
}
