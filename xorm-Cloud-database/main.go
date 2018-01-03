package main

import (
	"net/http"
	"xorm-Cloud-database/service"
)

func main() {
	server := service.NewServer()
	http.ListenAndServe(":8282", server)
}
