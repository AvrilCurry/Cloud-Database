package main

import "Cloud-database/service"
import "net/http"

func main() {
	server := service.NewServer()

	http.ListenAndServe(":8181", server)
}
