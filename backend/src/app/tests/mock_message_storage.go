package tests

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sedi/message-board/model"
)

type MockMessageStorage struct{}

func (s *MockMessageStorage) GetMessages() ([]*model.Message, error) {
	messages := []*model.Message{
		{ID: 1, UserID: "user1", Content: "Message 1", CreatedAt: time.Now()},
		{ID: 2, UserID: "user2", Content: "Message 2", CreatedAt: time.Now()},
	}
	return messages, nil
}

func (s *MockMessageStorage) CreateMessage(message *model.Message) error {
	return nil // return no error and hanlded create message
}

func (s *MockMessageStorage) DeleteMessage(id int) error {
	return nil  // return no error and hanlded delete message
}

func (s *MockMessageStorage) UpdateMessage(message *model.Message) error {
	return nil // return no error and hanlded update message
}

func (s *MockMessageStorage) GetMainMessages() ([]*model.Message, error) {
	mainMessages := []*model.Message{
		{ID: 3, UserID: "user3", Content: "Main Message 1", CreatedAt: time.Now(), ParentID: nil},
		{ID: 4, UserID: "user4", Content: "Main Message 2", CreatedAt: time.Now(), ParentID: nil},
	}
	return mainMessages, nil
}
func (s *MockMessageStorage) GetMessageByUserID(UserID string) ([]*model.Message, error) {
	_, err := uuid.Parse(UserID)
	if err != nil {
		return nil, model.MessageBoardError{Code: http.StatusBadRequest, Message: "not a valid uuid"}
	}
	mainMessages := []*model.Message{
		{ID: 4, UserID: "user3", Content: "Main Message 1", CreatedAt: time.Now(), ParentID: nil},
		{ID: 5, UserID: "user4", Content: "Main Message 2", CreatedAt: time.Now(), ParentID: nil},
	}
	return mainMessages, nil
}

func (s *MockMessageStorage) GetMessageReplies(id int) ([]*model.Message, error) {
	return nil, nil
}
