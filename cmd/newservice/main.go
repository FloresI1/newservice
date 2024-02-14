package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/FloresI1/newservice/internal/database/db"
)

/*
func EditNews(c *fiber.Ctx) error {

}

func ListNews(c *fiber.Ctx) error {

}

func AddNews(c *fiber.Ctx) error {

}

func DeleteNews(c *fiber.Ctx) error {

}
*/
func main() {
	app := fiber.New()
	// Инициализируем подключение к базе данных
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при инициализации базы данных: %v", err)
	}
	defer db.Close()

	//app.Post("/edit/:Id", EditNews)
	//app.Get("/list", ListNews)

	log.Fatal(app.Listen(":8080"))
}
