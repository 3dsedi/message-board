package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sedi/message-board/db"
	"github.com/sedi/message-board/model"
)

type MessageAPI struct {
	store db.MessageStorage
}

func NewMessageAPI(store db.MessageStorage) *MessageAPI {
	return &MessageAPI{
		store: store,
	}
}

func (api *MessageAPI) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", makeHTTPHandleFunc(api.handleMessages)).Methods("GET", "POST", "PUT", "DELETE")
	router.HandleFunc("/{id}", makeHTTPHandleFunc(api.handleUserMessages)).Methods("GET", "DELETE")
	router.HandleFunc("/reply/{id}", makeHTTPHandleFunc(api.handleGetRepliesByID)).Methods("GET")
}

func (api *MessageAPI) handleMessages(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return api.handleMainMessages(w, r)
	case http.MethodPost:
		return api.handlePostMessage(w, r)
	case http.MethodPut:
		return api.handleEditMessage(w, r)
	case http.MethodDelete:
		return api.handleDeleteMessage(w, r)
	default:
		return fmt.Errorf("unsupported method")
	}
}

func (api *MessageAPI) handleUserMessages(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return api.handleGetMessageByUserID(w, r)
	case http.MethodDelete:
		return api.handleDeleteMessage(w, r)
	default:
		return fmt.Errorf("unsupported method")
	}
}

func (api *MessageAPI) handleMainMessages(resWriter http.ResponseWriter, req *http.Request) error {
	messages, err := api.store.GetMainMessages()
	if err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(resWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}

		log.Printf("Error fetching messages: %s", err.Error())
		return WriteJSON(resWriter, http.StatusInternalServerError, ApiError{Error: "Failed to fetch messages"})
	}
	return WriteJSON(resWriter, http.StatusOK, messages)
}

func (api *MessageAPI) handleGetMessageByUserID(respWriter http.ResponseWriter, reqs *http.Request) error {
	id := mux.Vars(reqs)["id"]
	log.Printf("received request to get message with user ID %s", id)
	_, err := uuid.Parse(id)
	if err != nil {
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "not a valid uuid"})
	}

	messages, err := api.store.GetMessageByUserID(id)
	if err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error fetching messages for user ID %s: %s", id, err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to fetch message"})
	}

	return WriteJSON(respWriter, http.StatusOK, messages)
}

func (api *MessageAPI) handleGetRepliesByID(respWriter http.ResponseWriter, reqs *http.Request) error {
	id, found := getIDFromRequest(reqs)

	if !found {
		log.Println("ID not found in request")
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "ID not found in request"})
	}

	replies, err := api.store.GetMessageReplies(id)
	if err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error fetching replies for ID %d: %s", id, err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to fetch message replies"})
	}
	return WriteJSON(respWriter, http.StatusOK, replies)
}

func (api *MessageAPI) handlePostMessage(respWriter http.ResponseWriter, reqs *http.Request) error {
	createMessageReq := new(model.CreateMessageRequest)
	isValid, writer := postValidation(reqs, createMessageReq, respWriter)
	if isValid {
		return writer
	}

	log.Printf("received request to create a new message: %+v", createMessageReq)

	var parentID *int
	if createMessageReq.ParentID != 0 {
		parentID = &createMessageReq.ParentID
	}
	message := model.NewMessage(createMessageReq.Content, createMessageReq.UserID, parentID)

	if err := api.store.CreateMessage(message); err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error creating message: %s", err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to create message"})
	}
	return WriteJSON(respWriter, http.StatusOK, map[string]string{"message": "Message added successfully"})
}

func postValidation(reqs *http.Request, createMessageReq *model.CreateMessageRequest, respWriter http.ResponseWriter) (bool, error) {
	if err := json.NewDecoder(reqs.Body).Decode(createMessageReq); err != nil {
		return true, WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Invalid request body"})
	}

	if createMessageReq.Content == "" || createMessageReq.UserID == "" {
		return true, WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Missing required fields"})
	}

	_, err := uuid.Parse(createMessageReq.UserID)
	if err != nil {
		return true, WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "not a valid uuid"})
	}
	return false, nil
}

func (api *MessageAPI) handleEditMessage(respWriter http.ResponseWriter, reqs *http.Request) error {

	editReq := new(model.EditRequest)
	err := json.NewDecoder(reqs.Body).Decode(editReq)
	if err != nil {
		log.Println("invalid request")
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Invalid request body"})
	}
	message := &model.Message{
		ID:      editReq.ID,
		Content: editReq.Content,
	}
	err = api.store.UpdateMessage(message)
	if err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error updating message: %s", err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to update message"})
	}
	err = WriteJSON(respWriter, http.StatusOK, map[string]string{"message": "Message updated successfully"})
	if err != nil {
		log.Printf("Error writing JSON response: %s", err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to update message"})
	}
	return nil
}

func (api *MessageAPI) handleDeleteMessage(respWriter http.ResponseWriter, reqs *http.Request) error {
	id, found := getIDFromRequest(reqs)
	log.Println("found", found)

	if !found {
		log.Println("ID not found or invalid in request")
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Invalid ID format"})
	}

	log.Printf("Attempting to delete message with ID: %d\n", id)

	if err := api.store.DeleteMessage(id); err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error deleting message with ID %d: %s\n", id, err.Error())
		return fmt.Errorf("error deleting message: %w", err)
	}

	log.Printf("Message with ID %d deleted successfully\n", id)

	response := map[string]string{"message": "Message deleted successfully"}
	if err := WriteJSON(respWriter, http.StatusOK, response); err != nil {
		return fmt.Errorf("error writing JSON response: %w", err)
	}

	return nil
}
