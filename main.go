package main

import (
	"fmt"
	DELETE2 "github.com/adeemgoogle/go2/internal/handlers/DELETE"
	GET2 "github.com/adeemgoogle/go2/internal/handlers/GET"
	"github.com/adeemgoogle/go2/internal/handlers/PATCH"
	POST2 "github.com/adeemgoogle/go2/internal/handlers/POST"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func main() {
	fmt.Println()
	app := fiber.New()
	app.Get("/GetAuthors", GET2.GetAuthors)
	app.Get("/GetMembers", GET2.GetMembers)
	app.Get("/GetBooks", GET2.GetBooks)
	app.Get("/borrowedBook", GET2.BorrowedBooks)
	app.Get("/authors/:id/books", GET2.AuthorsBook)

	app.Post("/createAuthor", POST2.CreateAuthor)
	app.Post("/createMembers", POST2.CreateMember)
	app.Post("/createBook", POST2.CreateBook)
	app.Post("/createBooksb", POST2.CreateBRbook)

	app.Delete("/author/:id", DELETE2.AuthorDelete)
	app.Delete("/Books/:id", DELETE2.BookDelete)
	app.Patch("/author/:id", PATCH.UpdAuthors)

	err := app.Listen(":8080")
	if err != nil {
		return
	}

	if err := initConfig(); err != nil {
		log.Fatal("error with init", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
