package main

import (
	"net/http"

	"github.com/IgorMijatovic/go-course/pkg/handlers"
)



const portNumber = ":8080";

func main() {
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/", handlers.Home)

	_ = http.ListenAndServe(portNumber, nil)
}
