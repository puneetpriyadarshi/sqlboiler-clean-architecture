package main

import (
	"fmt"
	httphandler "root/api/handlers"
	"root/api/usecase"
	"root/db"

	"github.com/gin-gonic/gin"
)

func main() {
	dbconn, err := db.Connect()
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}

	r := gin.New()

	student := usecase.NewStudentUsecase(dbconn)
	httphandler.NewStudentHandler(r, student)

	_ = r.Run(":5000")
}
