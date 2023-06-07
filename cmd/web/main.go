package main

import (
	"fmt"
	"net/http"

	"github.com/moyvincent/GoWebApp/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	} else {
		fmt.Println("Server Running...")
	}
}
