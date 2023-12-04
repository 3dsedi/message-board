package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sedi/message-board/db"
	"github.com/sedi/vmessage-board/model"
	"github.com/stretchr/testify/assert"
)

func setup() (*mux.Router, *MockMessageStorage) {
	router := mux.NewRouter()
	mockStorage := &MockMessageStorage{}

	return router, mockStorage
}

func TestGetMainMessages(t *testing.T) {
	router, mockStorage := setup()
	router.HandleFunc("/message", makeHandleGetMessages(mockStorage)).Methods("GET")
	req, err := http.NewRequest("GET", "/message", nil)

	respRecord := httptest.NewRecorder()
	router.ServeHTTP(respRecord, req)

	assert.Equal(t, http.StatusOK, respRecord.Code)

	var response []model.Message
	err = json.NewDecoder(respRecord.Body).Decode(&response)
	assert.NoError(t, err)

	assert.NotNil(t, response)
}

func TestAddNewMessages(t *testing.T) {
	message := model.Message{
		ID:      1,
		Content: "Test message",
		UserID:  "c81a2c15-0ded-4b57-99b9-618185063e6b",
	}
	router, mockStorage := setup()
	router.HandleFunc("/message", makeHandleAddNewMessage(mockStorage, message)).Methods("POST")

	messageJSON, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}
	req, err := http.NewRequest("POST", "/message", bytes.NewBuffer(messageJSON))
	req.Header.Set("Content-Type", "application/json")

	respRecord := httptest.NewRecorder()
	router.ServeHTTP(respRecord, req)

	assert.Equal(t, http.StatusOK, respRecord.Code)
}
func TestUserMEssagesByNotCorrectUserId(t *testing.T) {
	router, mockStorage := setup()
	wrongUserId := "123-456"
	url := fmt.Sprintf("/message/%s", wrongUserId)

	router.HandleFunc("/message/{id}", makeHandleGetMessagesByUserId(mockStorage, wrongUserId)).Methods("GET")
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		t.Fatal(err)
	}
	respRecord := httptest.NewRecorder()
	router.ServeHTTP(respRecord, req)

	assert.Equal(t, http.StatusBadRequest, respRecord.Code)

}

func makeHandleGetMessagesByUserId(storage db.MessageStorage, userID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messages, err := storage.GetMessageByUserID(userID)
		if err != nil {
			if boardErr, ok := err.(model.MessageBoardError); ok {
				http.Error(w, boardErr.Error(), boardErr.Code)
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	}
}
func makeHandleGetMessages(storage db.MessageStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messages, err := storage.GetMainMessages()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	}
}

func makeHandleAddNewMessage(storage db.MessageStorage, msg model.Message) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := storage.CreateMessage(&msg)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
