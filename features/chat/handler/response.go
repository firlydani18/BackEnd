package handler

import (
	"KosKita/features/chat"
	"KosKita/features/user/handler"
)

type RoomRes struct {
	ID string `json:"room_id"`
}

type GetRoomRespon struct {
	ID           string `json:"room_id"`
	SenderID     uint   `json:"sender_id"`
	ReceiverID   uint   `json:"receiver_id"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
}

type ChatRes struct {
	ID         string `json:"room_id"`
	Message    string `json:"message"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
}

func CoreToGetChat(chat chat.Core) ChatRes {
	return ChatRes{
		ID:         chat.RoomID,
		Message:    chat.Message,
		SenderID:   chat.SenderID,
		ReceiverID: chat.ReceiverID,
	}
}

func CoreToGetChats(chats []chat.Core) []ChatRes {
	res := make([]ChatRes, 0)
	for _, chat := range chats {
		res = append(res, ChatRes{
			ID:         chat.RoomID,
			Message:    chat.Message,
			SenderID:   chat.SenderID,
			ReceiverID: chat.ReceiverID,
		})
	}
	return res
}

func CoreToGetUser(room chat.Core) GetRoomRespon {
	user := handler.CoreToResponse(&room.User)
	return GetRoomRespon{
		ID:           room.RoomID,
		SenderID:     room.SenderID,
		ReceiverID:   room.ReceiverID,
		Name:         user.UserName,
		PhotoProfile: user.PhotoProfile,
	}
}

