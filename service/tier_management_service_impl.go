package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"github.com/indrawanagung/loyalty_management_api/repository"
	"github.com/indrawanagung/loyalty_management_api/util"
	"gorm.io/gorm"
)

type TierManagementServiceImpl struct {
	Database                 *gorm.DB
	Validate                 *validator.Validate
	TierManagementRepository repository.TierManagementRepositoryInterface
}

func NewTierManagementService(
	Database *gorm.DB,
	Validate *validator.Validate,
	TierManagementRepository repository.TierManagementRepositoryInterface,
) TierManagementServiceInterface {
	return &TierManagementServiceImpl{
		Database:                 Database,
		Validate:                 Validate,
		TierManagementRepository: TierManagementRepository,
	}
}

func (s TierManagementServiceImpl) FindAll() []web.TierManagementResponse {
	tiers := s.TierManagementRepository.FindAll(s.Database)
	return web.ToTierManagementResponses(tiers)
}

func (s TierManagementServiceImpl) FindByID(tierID int) web.TierManagementResponse {
	tier, err := s.TierManagementRepository.FindByID(s.Database, tierID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ToTierManagementResponse(tier)

}

func (s TierManagementServiceImpl) Save(request web.TierManagementRequest) int {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	id := s.TierManagementRepository.SaveOrUpdate(s.Database, domain.TierManagement{
		Name:     request.Name,
		MinPoint: request.MinPoint,
		MaxPoint: request.MaxPoint,
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	})

	return id
}

func (s TierManagementServiceImpl) Updated(request web.TierManagementRequest, tierID int) {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}
	tier, err := s.TierManagementRepository.FindByID(s.Database, tierID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.TierManagementRepository.SaveOrUpdate(s.Database, domain.TierManagement{
		ID:       tierID,
		Name:     request.Name,
		MinPoint: request.MinPoint,
		MaxPoint: request.MaxPoint,
		Timestamp: domain.Timestamp{
			CreatedAt: tier.CreatedAt,
			UpdatedAt: util.GetUnixTimestamp(),
		},
	})

}

func (s TierManagementServiceImpl) Delete(tierID int) {
	_, err := s.TierManagementRepository.FindByID(s.Database, tierID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	s.TierManagementRepository.Delete(s.Database, tierID)
}
