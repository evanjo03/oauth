package main

import (
	"fmt"
	"net/http"

	"github.com/evanjo03/oauth/internal/routes"
)

const (
	Port = 8080
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.FileHandler)
	mux.HandleFunc("/oauth", routes.OauthHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", Port), mux)
}
