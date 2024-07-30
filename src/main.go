package main

import (
	"fmt"
	"log"
	"tractian-go/configs"
	"tractian-go/internal/controller"
	"tractian-go/internal/repository"
	"tractian-go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf(err.Error())
	}

	envs := configs.GetEnvs()
	db := configs.NewConnDB(*envs)

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar a conexão: ", err)
	}

	// dependency injection
	repo := repository.NewRepo(db)
	service := service.NewService(repo)
	controller := controller.NewController(*service)

	router := gin.New()

	routes := router.Group("/v1")
	routes.GET("/", controller.List)

	fmt.Println("to aqui")
}
