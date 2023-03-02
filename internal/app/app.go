package app

import (
	"github.com/ProSt1ll/makves-test/internal/store/storeCSV"
	"github.com/ProSt1ll/makves-test/internal/transport"
	"net/http"
)

func configureRoutes(serveMux *http.ServeMux, Server *transport.Server) {
	serveMux.HandleFunc("/get-items", Server.HandlerGetItems)
}

func Run() error {

	serveMux := http.NewServeMux()
	var store transport.Store
	store = storeCSV.New()
	storeServer := transport.NewServer(store)
	configureRoutes(serveMux, storeServer)

	login := transport.Logging(serveMux)

	return http.ListenAndServe("localhost:8080", login)
}
