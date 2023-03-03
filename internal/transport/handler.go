package transport

import (
	"net/http"
	"strings"
)

func NewServer(store Store) *Server {
	return &Server{store: store}
}

//HandlerGetItems метод для получения Item'ов
func (ss *Server) HandlerGetItems(w http.ResponseWriter, r *http.Request) {
	//получаем ключи
	str := handlerDataGet(r, "id")
	//достаем их из базы
	res, err := ss.store.Get("id", str)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")

	//проверяем наличие данных, если нет - не найдено
	if res == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	//пишем данные
	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

//handlerDataGet функция парсинга ключей
func handlerDataGet(r *http.Request, key string) []string {
	var ids []string
	idString := r.FormValue(key)
	if idString != "" {
		idStrings := strings.Split(idString, ",")
		for _, v := range idStrings {
			ids = append(ids, v)
		}
	}
	return ids
}
