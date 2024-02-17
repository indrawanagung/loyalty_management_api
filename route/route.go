package route

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/indrawanagung/loyalty_management_api/controller"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"github.com/spf13/viper"
)

func New(
	membershipController controller.MembershipControllerInterface,
	tierManagementController controller.TierManagementControllerInterface,
	transactionController controller.TransactionControllerInterface,
) *fiber.App {
	authentication := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("SECRET_KEY"))},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			webResponse := web.WebResponse{
				Header: web.Header{
					Message: "Invalid or Expired Token",
					Error:   true,
				},
				Data: nil,
			}
			return ctx.Status(401).JSON(webResponse)
		},
	})
	app := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})

	app.Use(recover.New())

	api := app.Group("/api")

	api.Post("/sign_in", membershipController.SignIn)

	api.Post("/memberships", authentication, membershipController.Save)
	api.Post("/memberships/activity", authentication, membershipController.AddMemberActivity)
	api.Post("/memberships/redeemed_point", authentication, membershipController.RedeemedPoint)
	api.Post("/memberships/redeemed_point/histories", authentication, membershipController.FindAllRedeemedPoint)
	api.Post("/memberships/earned_point/histories", authentication, membershipController.FindAllEarnedPoint)

	api.Get("/tiers", authentication, tierManagementController.FindAll)
	api.Get("/tiers/:tierID", authentication, tierManagementController.FindByID)
	api.Post("/tiers", authentication, tierManagementController.Save)
	api.Put("/tiers/:tierID", authentication, tierManagementController.Updated)
	api.Delete("/tiers/:tierID", authentication, tierManagementController.Delete)

	api.Post("/transaction", authentication, transactionController.AddTransaction)

	return app
}
