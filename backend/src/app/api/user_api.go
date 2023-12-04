package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sedi/messageBoard/db"
	"github.com/sedi/messageBoard/model"
)

type UserAPI struct {
	store db.UserStorage
}

func NewUserAPI(store db.UserStorage) *UserAPI {
	return &UserAPI{
		store: store,
	}
}

func (api *UserAPI) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", makeHTTPHandleFunc(api.handleUsers)).Methods("GET", "POST")
	router.HandleFunc("/login", makeHTTPHandleFunc(api.handleLoginUser)).Methods("POST")
}

func (api *UserAPI) handleUsers(w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case http.MethodGet:
		return api.handleGetUsers(w, req)
	case http.MethodPost:
		return api.handleCreateUser(w, req)
	default:
		return fmt.Errorf("unsupported method")
	}
}

func (api *UserAPI) handleGetUsers(respWriter http.ResponseWriter, reqs *http.Request) error {
	users, err := api.store.GetUsers()
	if err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}

		log.Printf("Error fetching users: %s", err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to fetch users"})
	}
	return WriteJSON(respWriter, http.StatusOK, users)
}

func (api *UserAPI) handleCreateUser(respWriter http.ResponseWriter, reqs *http.Request) error {
	createUser := new(model.CreateUserRequest)
	if err := json.NewDecoder(reqs.Body).Decode(createUser); err != nil {
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Invalid request body"})
	}
	log.Printf("Create user request: %+v", createUser)
	if createUser.UserName == "" || createUser.Email == "" || createUser.Password == "" {
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Missing required fields"})
	}
	user := model.NewUser(createUser.UserName, createUser.Email, createUser.Password)
	if err := api.store.CreateUser(user); err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error creating user: %s", err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to create user"})
	}
	return WriteJSON(respWriter, http.StatusOK, map[string]string{"message": "User added successfully"})
}

func (api *UserAPI) handleLoginUser(respWriter http.ResponseWriter, reqs *http.Request) error {
	getUser := new(model.GetUserRequest)
	if err := json.NewDecoder(reqs.Body).Decode(getUser); err != nil {
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Invalid request body"})
	}
	if getUser.Email == "" || getUser.Password == "" {
		return WriteJSON(respWriter, http.StatusBadRequest, ApiError{Error: "Missing required fields"})
	}

	user, err := api.store.GetUserByEmailAndPassword(getUser.Email, getUser.Password)
	if err != nil {
		if boardErr, ok := err.(model.MessageBoardError); ok {
			return WriteJSON(respWriter, boardErr.Code, ApiError{Error: boardErr.Error()})
		}
		log.Printf("Error fetching user: %s", err.Error())
		return WriteJSON(respWriter, http.StatusInternalServerError, ApiError{Error: "Failed to fetch user"})
	}
	return WriteJSON(respWriter, http.StatusOK, user)
}
