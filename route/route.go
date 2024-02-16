package route

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	//authentication := jwtware.New(jwtware.Config{
	//	SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("SECRET_KEY"))},
	//})
	//app.Use(recover.New())
	//app := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})
	app := fiber.New(fiber.Config{})
	api := app.Group("/api")

	api.Get("/users", func(ctx *fiber.Ctx) error {
		return ctx.SendString("test")
	})

	return app
}
