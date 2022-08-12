package main

import (
	"example.com/m/v2/src/api"
	"example.com/m/v2/src/service"
	_ "github.com/lib/pq"
)

func main() {

	service.StartHttpService(8080, api.HttpService())
}
