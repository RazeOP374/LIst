package main

import (
	"GOproject/GIT/memory_note/config"
	"GOproject/GIT/memory_note/routes"
)

func main() {
	// http://localhost:3000/swagger/index.html
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}
