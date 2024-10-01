package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	// Adiciona o handler para /account
	mux.Handle("/account", reverseProxy("http://localhost:8081/account"))

	// Adiciona o CORS handler
	http.ListenAndServe(":8888", corsHandler(mux))
}

func reverseProxy(target string) http.Handler {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Método HTTP: %s, Path: %s", r.Method, r.URL.Path)

		if strings.HasPrefix(r.URL.Path, "/account") {
			println("Passou pelo reverse proxy")
			proxy.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATH, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
