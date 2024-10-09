package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"playground/routes"
)

func main() {
	router := setupRoutes()

	fmt.Println("Running on 3333")
	err := http.ListenAndServe(":3333", router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func setupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.GetRoot)
	mux.HandleFunc("/hello", routes.GetHello)
	return mux
}
