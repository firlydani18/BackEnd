package service

import (
	cc "KosKita/features/chat"
	cd "KosKita/features/chat/data"
	cu "KosKita/features/user"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn       *websocket.Conn
	Message    chan *cd.Chat
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Name       string `json:"name"`
	RoomID     string `json:"roomId"`
}

type ChatRes struct {
	Message    string `json:"message"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
	RoomID     string `json:"room_id"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		result := ChatRes{
			Message:    message.Message,
			RoomID:     message.RoomID,
			ReceiverID: message.ReceiverID,
			SenderID:   message.SenderID,
		}

		c.Conn.WriteJSON(result)
	}
}

func (c *Client) ReadMessage(hub *Hub, chatService cc.ChatServiceInterface, cu cu.UserServiceInterface) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		senderID, err := strconv.Atoi(c.SenderID)
		if err != nil {
			log.Printf("Error converting SenderID to integer: %v", err)
			continue
		}

		receiverID, err := strconv.Atoi(c.ReceiverID)
		if err != nil {
			log.Printf("Error converting ReceiverID to integer: %v", err)
			continue
		}

		msg := &cd.Chat{
			Message:    string(m),
			SenderID:   uint(senderID),
			ReceiverID: uint(receiverID),
			RoomID:     c.RoomID,
		}

		coreMsg := cc.Core{
			Message:    msg.Message,
			RoomID:     msg.RoomID,
			ReceiverID: msg.ReceiverID,
			SenderID:   msg.SenderID,
		}

		_, err = chatService.CreateChat(senderID, receiverID, coreMsg)
		if err != nil {
			log.Printf("Error saving message: %v", err)
			continue
		}
		hub.Broadcast <- msg
	}
}
