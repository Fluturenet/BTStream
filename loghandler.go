package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.Host, r.URL)
		handler.ServeHTTP(w, r)
	})

}
