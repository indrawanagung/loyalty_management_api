package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"github.com/indrawanagung/loyalty_management_api/service"
	"github.com/indrawanagung/loyalty_management_api/util"
)

type MembershipControllerImpl struct {
	MembershipService service.MembershipServiceInterface
}

func NewMembershipController(MembershipService service.MembershipServiceInterface) MembershipControllerInterface {
	return &MembershipControllerImpl{MembershipService: MembershipService}
}

func (c MembershipControllerImpl) FindByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c MembershipControllerImpl) FindByEmail(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c MembershipControllerImpl) Save(ctx *fiber.Ctx) error {
	request := new(web.MemberCreateOrUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	member := ctx.Locals("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberID := claims["id"].(float64)
	id := int(memberID)

	transactionID := c.MembershipService.Save(*request, id)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"transaction_id": transactionID},
	}
	return ctx.Status(201).JSON(webResponse)
}

func (c MembershipControllerImpl) SignIn(ctx *fiber.Ctx) error {
	request := new(web.SignInRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	id := c.MembershipService.SignIn(*request)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"access_token": id},
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c MembershipControllerImpl) AddMemberActivity(ctx *fiber.Ctx) error {
	request := new(web.AddMemberActivityRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	member := ctx.Locals("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberID := claims["id"].(float64)
	id := int(memberID)

	transactionID := c.MembershipService.AddMemberActivity(*request, id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"transaction_id": transactionID},
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c MembershipControllerImpl) RedeemedPoint(ctx *fiber.Ctx) error {
	request := new(web.RedeemedPointRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	member := ctx.Locals("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberID := claims["id"].(float64)
	id := int(memberID)

	c.MembershipService.AddRedeemedPoint(*request, id)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c MembershipControllerImpl) FindAllRedeemedPoint(ctx *fiber.Ctx) error {
	member := ctx.Locals("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberID := claims["id"].(float64)
	id := int(memberID)

	historyResponses := c.MembershipService.FindAllRedeemedPointHistory(id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   historyResponses,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c MembershipControllerImpl) FindAllEarnedPoint(ctx *fiber.Ctx) error {
	member := ctx.Locals("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberID := claims["id"].(float64)
	id := int(memberID)

	historyResponses := c.MembershipService.FindAllEarnedPointHistory(id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   historyResponses,
	}
	return ctx.Status(200).JSON(webResponse)
}
