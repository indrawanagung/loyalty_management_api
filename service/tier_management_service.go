package service

import (
	"github.com/indrawanagung/loyalty_management_api/model/web"
)

type TierManagementServiceInterface interface {
	FindAll() []web.TierManagementResponse
	FindByID(tierID int) web.TierManagementResponse
	Save(request web.TierManagementRequest) int
	Updated(request web.TierManagementRequest, tierID int)
	Delete(tierID int)
}
