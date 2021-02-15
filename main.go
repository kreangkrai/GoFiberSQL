package main

import (
	"GoFiberSQL/Controller"
	"GoFiberSQL/Router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := Controller.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	api := app.Group("/api") // /api

	device := api.Group("/device") // /api/v1

	device.Get("/get/:device", Router.Get)
	device.Get("/gets", Router.Gets)
	device.Post("/insert", Router.Insert)
	device.Put("/update", Router.Update)
	device.Delete("/delete/:device", Router.Delete)

	// var port = os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}

	app.Listen(":" + port)
}
