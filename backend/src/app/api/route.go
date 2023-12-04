package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sedi/message-board/db"
)

type APIServer struct {
	listenAddr string
	userAPI    *UserAPI
	messageAPI *MessageAPI
}

func NewAPIServer(listenAddr string, storeUser db.UserStorage, storeMessage db.MessageStorage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		userAPI:    NewUserAPI(storeUser),
		messageAPI: NewMessageAPI(storeMessage),
	}
}

func (server *APIServer) SetupRouter() *mux.Router {
	router := mux.NewRouter()

	server.userAPI.RegisterRoutes(router.PathPrefix("/user").Subrouter())
	server.messageAPI.RegisterRoutes(router.PathPrefix("/message").Subrouter())
	return router
}

func (server *APIServer) Run() {
	envFile, _ := godotenv.Read(".env")
	
	router := server.SetupRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{envFile["CLIENT_ADDRESS"]})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	log.Println("Listening on", server.listenAddr)
	log.Fatal(http.ListenAndServe(server.listenAddr, corsRouter))
}
