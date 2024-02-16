package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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
	fmt.Println(123)
	id := c.MembershipService.Save(*request)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]int{"member_id": id},
	}
	return ctx.Status(201).JSON(webResponse)
}

func (c MembershipControllerImpl) SignIn(ctx *fiber.Ctx) error {
	request := new(web.SignInRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	fmt.Println(123)
	id := c.MembershipService.SignIn(*request)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"access_token": id},
	}
	return ctx.Status(200).JSON(webResponse)
}
