package controller

import "github.com/gofiber/fiber/v2"

type TransactionControllerInterface interface {
	AddTransaction(ctx *fiber.Ctx) error
}
