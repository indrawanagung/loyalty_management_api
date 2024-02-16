package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"github.com/indrawanagung/loyalty_management_api/service"
	"github.com/indrawanagung/loyalty_management_api/util"
	"strconv"
)

type TierManagementControllerImpl struct {
	TierManagementService service.TierManagementServiceInterface
}

func NewTierManagementController(TierManagementService service.TierManagementServiceInterface) TierManagementControllerInterface {
	return &TierManagementControllerImpl{TierManagementService: TierManagementService}
}

func (c TierManagementControllerImpl) FindAll(ctx *fiber.Ctx) error {
	tierResponses := c.TierManagementService.FindAll()

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   tierResponses,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c TierManagementControllerImpl) FindByID(ctx *fiber.Ctx) error {
	tierID := ctx.Params("tierID")
	id, err := strconv.Atoi(tierID)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(err.Error()))
	}
	tierResponse := c.TierManagementService.FindByID(id)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   tierResponse,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c TierManagementControllerImpl) Save(ctx *fiber.Ctx) error {
	request := new(web.TierManagementRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	id := c.TierManagementService.Save(*request)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]int{"tier_id": id},
	}
	return ctx.Status(201).JSON(webResponse)
}

func (c TierManagementControllerImpl) Updated(ctx *fiber.Ctx) error {
	tierID := ctx.Params("tierID")
	id, err := strconv.Atoi(tierID)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(err.Error()))
	}

	request := new(web.TierManagementRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	c.TierManagementService.Updated(*request, id)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c TierManagementControllerImpl) Delete(ctx *fiber.Ctx) error {
	tierID := ctx.Params("tierID")
	id, err := strconv.Atoi(tierID)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(err.Error()))
	}
	c.TierManagementService.Delete(id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}
