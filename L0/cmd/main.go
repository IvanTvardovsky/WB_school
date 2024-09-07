package main

import (
	"wb_l0/internal/repo"
	"wb_l0/internal/server"
)

func main() {
	repo.InitDB()
	repo.RestoreCacheFromDB()

	go repo.SubscribeToNATS()

	server.StartHTTPServer()
}
