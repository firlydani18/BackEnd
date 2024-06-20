package handler

import (
	"fmt"
	"math/rand"
	"time"
)

type CreateRoomReq struct {
	ReceiverID int `json:"receiver_id"`
	SenderID   int `json:"sender_id"`
}

func generateRoomID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%05d", rand.Intn(100000))
}

