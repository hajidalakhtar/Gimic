package main

import (
	"Gimic/routes"
	"Gimic/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	switch args[0] {

	case "migrate":
		database.Migrate()
		fmt.Println("Migrate : Success")

	case "test":
		fmt.Println("test")

	case "serve":
		m := routes.SetRoute()
		m.Run(":8080")

	case "microservice":
		service := web.NewService(
		web.Name("go.micro.api.user"),
		)
		service.Init()
		router := routes.SetRoute()
		router.HandleMethodNotAllowed = true
		router.NoMethod(func(c *gin.Context) {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
			return
		})
		router.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
			return
		})

		service.Handle("/", router)
		if err := service.Run(); err != nil {log.Fatal(err)}
	}


}


