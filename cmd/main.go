package main

import (
	"fmt"
	"github.com/crisguitar/paraules-noves/api"
	"net/http"
	"os"
)

func main() {
	r := api.NewRouter()
	port, found := os.LookupEnv("PORT")
	if !found || port == "" {
		port = "8000"
	}
	fmt.Printf("Starting app in port %s \n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		fmt.Println("Could not start app :(")
	}
}
