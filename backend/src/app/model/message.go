package model

import (
	"time"
)

type CreateMessageRequest struct {
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
	ParentID int    `json:"parent_id"`
}

type EditRequest struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Message struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	ParentID  *int   `json:"parent_id"`
	CreatedAt time.Time
}

func NewMessage(content, userID string, parentID *int) *Message {

	return &Message{
		UserID:    userID,
		Content:   content,
		ParentID:  parentID,
		CreatedAt: time.Now().UTC(),
	}
}
