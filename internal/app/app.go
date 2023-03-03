package app

import (
	"fmt"
	"github.com/ProSt1ll/makves-test/internal/store/storeCSV"
	"github.com/ProSt1ll/makves-test/internal/transport"
	"net/http"
)

func configureRoutes(serveMux *http.ServeMux, Server *transport.Server) {
	serveMux.HandleFunc("/get-items", Server.HandlerGetItems)
}

func Run() error {
	//создаем экземпляр Mux сервера
	serveMux := http.NewServeMux()

	//создаем экземпляр удовлетворяющей интерфейсу Store
	var store transport.Store
	store, err := storeCSV.New("internal/store/storeCSV/ueba.csv")
	if err != nil {
		return err
	}

	//создаем экзмепляр нашего сервера с обработчиками
	storeServer := transport.NewServer(store)

	//конфигурируем Mux сеовер
	configureRoutes(serveMux, storeServer)
	login := transport.Logging(serveMux)

	fmt.Println("Server started " + "0.0.0.0:8080")
	return http.ListenAndServe("0.0.0.0:8080", login)
}
