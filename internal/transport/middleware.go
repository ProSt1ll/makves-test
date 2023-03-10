package transport

import (
	"log"
	"net/http"
)

//Logging - каждый запрос проходит через данный обработчик и выводится его метод, URI и IP
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, req.RemoteAddr)
	})
}
