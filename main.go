package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/indrawanagung/loyalty_management_api/controller"
	"github.com/indrawanagung/loyalty_management_api/db"
	"github.com/indrawanagung/loyalty_management_api/repository"
	"github.com/indrawanagung/loyalty_management_api/route"
	"github.com/indrawanagung/loyalty_management_api/service"
	"github.com/indrawanagung/loyalty_management_api/util"
	"os"
)

func main() {
	validate := validator.New()

	config := util.LoadConfig(".")
	database := db.OpenConnection(config.DBSource)

	membershipRepository := repository.NewMembershipRepository()
	tierRepository := repository.NewTierManagementRepository()
	transactionRepository := repository.NewTransactionRepository()
	loyaltyRepository := repository.NewLoyaltyRepository()
	earnedRepository := repository.NewEarnedHistory()

	transactionService := service.NewTransactionService(database, validate, transactionRepository, tierRepository, membershipRepository, loyaltyRepository, earnedRepository)
	tierService := service.NewTierManagementService(database, validate, tierRepository)
	membershipService := service.NewMembershipService(database, membershipRepository, validate)

	transactionController := controller.NewTransactionController(transactionService)
	tierController := controller.NewTierManagementController(tierService)
	membershipController := controller.NewMembershipController(membershipService)

	app := route.New(membershipController, tierController, transactionController)
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	f, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(f)
	log.Info("server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
