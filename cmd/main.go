package main

import (
	"go-chat/db"
	"go-chat/internal/user"
	"go-chat/internal/ws"
	"go-chat/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalln("Could not create database connection")
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)
	userPresentation := user.NewPresentation(*userHandler)

	hub := ws.NewHub()
	go hub.Run()
	wsHandler := ws.NewHandler(hub)
	wsPresentation := ws.NewPresentation(*wsHandler)


	router.InitRouter(userPresentation, wsPresentation)
	router.Start("0.0.0.0:8080")
}