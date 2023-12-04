package main

import (
	"log"
	db "server/database"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("init database fail: %s", err)
	}

	userReq := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userReq)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
