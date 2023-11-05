package main

import (
	"fmt"
	"github.com/adeemgoogle/go2/pkg/handlers/ProstoiGet"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func main() {
	fmt.Println()
	app := fiber.New()
	app.Get("/GetAuthors", ProstoiGet.GetAuthors)
	app.Get("/GetMembers", ProstoiGet.GetMembers)
	app.Get("/GetBooks", ProstoiGet.GetBooks)

	app.Listen(":8080")

	if err := initConfig(); err != nil {
		log.Fatal("error with init", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
