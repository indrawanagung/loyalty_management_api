package controller

import (
	"github.com/gofiber/fiber/v2"
)

type MembershipControllerInterface interface {
	FindByID(ctx *fiber.Ctx) error
	FindByEmail(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
}
