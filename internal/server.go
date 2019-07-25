package internal

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nxshock/signaller/internal/models/capabilities"
)

var currServer *Server

type Server struct {
	httpServer *http.Server
	router     *mux.Router

	Address string

	Capabilities capabilities.Capabilities
	Backend      Backend
}

func New() *Server {
	router := mux.NewRouter()
	router.HandleFunc("/_matrix/client/versions", VersionHandler)
	router.HandleFunc("/_matrix/client/r0/login", LoginHandler)
	router.HandleFunc("/_matrix/client/r0/logout", LogoutHandler)
	router.HandleFunc("/_matrix/client/r0/logout/all", LogoutAllHandler)
	router.HandleFunc("/_matrix/client/r0/register", RegisterHandler)
	router.HandleFunc("/_matrix/client/r0/account/whoami", WhoAmIHandler)
	router.HandleFunc("/_matrix/client/r0/joined_rooms", JoinedRoomsHandler)
	router.HandleFunc("/_matrix/client/r0/sync", SyncHandler)
	router.HandleFunc("/_matrix/client/r0/capabilities", CapabilitiesHandler)
	router.HandleFunc("/", RootHandler)

	httpServer := new(http.Server)
	httpServer.Addr = ":8008"
	httpServer.Handler = router

	server := &Server{
		httpServer: httpServer,
		router:     router}

	currServer = server
	return server
}

func (server *Server) Run() error {
	return server.httpServer.ListenAndServe() // TODO: custom port
}
