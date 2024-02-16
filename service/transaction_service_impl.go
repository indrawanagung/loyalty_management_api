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
	"time"
)

type TransactionServiceImpl struct {
	Database                 *gorm.DB
	Validate                 *validator.Validate
	TransactionRepository    repository.TransactionRepositoryInterface
	TierManagementRepository repository.TierManagementRepositoryInterface
	MemberRepository         repository.MembershipRepositoryInterface
	LoyaltyRepository        repository.LoyaltyRepositoryInterface
	EarnedHistoryRepository  repository.EarnedHistoryInterface
}

func NewTransactionService(
	Database *gorm.DB,
	Validate *validator.Validate,
	TransactionRepository repository.TransactionRepositoryInterface,
	TierManagementRepository repository.TierManagementRepositoryInterface,
	MemberRepository repository.MembershipRepositoryInterface,
	LoyaltyRepository repository.LoyaltyRepositoryInterface,
	EarnedHistoryRepository repository.EarnedHistoryInterface,
) TransactionServiceInterface {
	return &TransactionServiceImpl{
		Database:                 Database,
		Validate:                 Validate,
		TransactionRepository:    TransactionRepository,
		TierManagementRepository: TierManagementRepository,
		MemberRepository:         MemberRepository,
		LoyaltyRepository:        LoyaltyRepository,
		EarnedHistoryRepository:  EarnedHistoryRepository,
	}
}

func (s TransactionServiceImpl) AddTransaction(request web.TransactionCreateRequest, memberID int) int {
	var items []domain.ItemTransactionDetails
	var totalAmount int64
	for _, itemRequest := range request.Transaction {
		item := domain.ItemTransactionDetails{
			ItemName:     itemRequest.ItemName,
			Price:        itemRequest.Price,
			Qty:          itemRequest.Qty,
			ItemSubtotal: itemRequest.Price * int64(itemRequest.Qty),
			CreatedAt:    util.GetUnixTimestamp(),
		}
		totalAmount += item.ItemSubtotal
		items = append(items, item)
	}

	member, err := s.MemberRepository.FindByID(s.Database, memberID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	var earnedPoint int64

	loyalties, err := s.TierManagementRepository.FindByPoint(s.Database, member.EarnedPoint)
	if err == nil {

		loyaltyTransactionPolicy, err := s.LoyaltyRepository.FindByID(s.Database, loyalties[0].ID)
		if err == nil {
			if totalAmount >= loyaltyTransactionPolicy.TransactionPolicy.TransactionAmount {
				earnedPoint += totalAmount * loyaltyTransactionPolicy.TransactionPolicy.Percentage / 100
			}
		}
	}

	itemTransaction := domain.ItemTransaction{
		MemberID:        memberID,
		TotalAmount:     totalAmount,
		TransactionDate: time.Now(),
		TransactionID:   util.GetUnixTimestamp(),
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
		ItemDetail: items,
	}

	itemTransactionId := s.TransactionRepository.AddTransaction(s.Database, itemTransaction)

	earnedPointHistory := domain.EarnedPointHistory{
		MemberID:               memberID,
		TransactionDate:        time.Now(),
		ReferenceTransactionID: itemTransactionId,
		LoyaltyProgramID:       loyalties[0].ID,
		ExistingPoint:          member.RemainedPoint,
		EarnedPoint:            earnedPoint,
		BalancePoint:           earnedPoint + member.RemainedPoint,
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	}
	member.EarnedPoint += earnedPoint
	member.RemainedPoint += earnedPoint
	member.UpdatedAt = util.GetUnixTimestamp()
	s.MemberRepository.Save(s.Database, member)
	s.EarnedHistoryRepository.Save(s.Database, earnedPointHistory)
	return itemTransactionId

}
