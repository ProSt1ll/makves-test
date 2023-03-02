package transport

import (
	"net/http"
	"strings"
)

func NewServer(store Store) *Server {
	return &Server{store: store}
}

func (ss *Server) HandlerGetItems(w http.ResponseWriter, r *http.Request) {
	str := handlerDataGet(r, "id")
	res, err := ss.store.Get("id", str)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

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
