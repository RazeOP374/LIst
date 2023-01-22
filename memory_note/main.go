package main

import (
	"GOproject/GIT/memory_note/config"
	"GOproject/GIT/memory_note/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}

//"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcl9OYW1lIjoid2FuZ2NhbyIsImV4cCI6MTY3NDI3NjgxOSwiaXNzIjoibGlzdCJ9.8eloAqcQVyqdDJt-AqPl_RN289vFn_5YCnDT-z_qWkE"
