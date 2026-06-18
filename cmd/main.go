package main

import (
	"go-conceptual/internal/config"
	"go-conceptual/internal/server"
)

func main() {
	env := config.LoadEnv()

	db := config.ConnectToDB(env)

	server.Start(env, db)
}
