package main

import (
	"github.com/victorradael/rest_api_go_gin/database"
	"github.com/victorradael/rest_api_go_gin/routes"
)

func main() {
	database.ConnectWithDatabase()
	routes.HandleRequests()
}
