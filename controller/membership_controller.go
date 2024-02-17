package controller

import (
	"github.com/gofiber/fiber/v2"
)

type MembershipControllerInterface interface {
	FindByID(ctx *fiber.Ctx) error
	FindByEmail(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	AddMemberActivity(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	RedeemedPoint(ctx *fiber.Ctx) error
	FindAllRedeemedPoint(ctx *fiber.Ctx) error
	FindAllEarnedPoint(ctx *fiber.Ctx) error
}
