package controller

import (
	"github.com/gofiber/fiber/v2"
)

type TierManagementControllerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	Updated(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
