package util

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"time"
)

func ConvertStringToTime(timeString string) time.Time {
	date, err := time.Parse("2006-01-02", timeString)

	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(err.Error()))
	}
	return date
}
