package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"github.com/indrawanagung/loyalty_management_api/service"
	"github.com/indrawanagung/loyalty_management_api/util"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionServiceInterface
}

func NewTransactionController(TransactionService service.TransactionServiceInterface) TransactionControllerInterface {
	return &TransactionControllerImpl{TransactionService: TransactionService}
}

func (c TransactionControllerImpl) AddTransaction(ctx *fiber.Ctx) error {
	member := ctx.Locals("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberID := claims["id"].(float64)
	id := int(memberID)

	request := new(web.TransactionCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	idResponse := c.TransactionService.AddTransaction(*request, id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]int{"item_transaction_id": idResponse},
	}
	return ctx.Status(201).JSON(webResponse)
}
