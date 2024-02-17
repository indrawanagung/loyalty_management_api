package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/indrawanagung/loyalty_management_api/exception"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"github.com/indrawanagung/loyalty_management_api/repository"
	"github.com/indrawanagung/loyalty_management_api/util"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type MembershipServiceImpl struct {
	Database                  *gorm.DB
	Validate                  *validator.Validate
	MembershipRepository      repository.MembershipRepositoryInterface
	TierManagementRepository  repository.TierManagementRepositoryInterface
	LoyaltyRepository         repository.LoyaltyRepositoryInterface
	EarnedHistoryRepository   repository.EarnedHistoryInterface
	RedeemedHistoryRepository repository.RedeemedHistoryInterface
}

func NewMembershipService(
	Database *gorm.DB,
	Validate *validator.Validate,
	MembershipRepository repository.MembershipRepositoryInterface,
	TierManagementRepository repository.TierManagementRepositoryInterface,
	LoyaltyRepository repository.LoyaltyRepositoryInterface,
	EarnedHistoryRepository repository.EarnedHistoryInterface,
	RedeemedHistoryRepository repository.RedeemedHistoryInterface,
) MembershipServiceInterface {
	return &MembershipServiceImpl{
		Database:                  Database,
		Validate:                  Validate,
		MembershipRepository:      MembershipRepository,
		TierManagementRepository:  TierManagementRepository,
		LoyaltyRepository:         LoyaltyRepository,
		EarnedHistoryRepository:   EarnedHistoryRepository,
		RedeemedHistoryRepository: RedeemedHistoryRepository,
	}
}

func (s MembershipServiceImpl) FindByID(ID int) domain.Membership {
	member, err := s.MembershipRepository.FindByID(s.Database, ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return member
}

func (s MembershipServiceImpl) FindByEmail(email string) domain.Membership {
	member, err := s.MembershipRepository.FindByEmail(s.Database, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return member
}

func (s MembershipServiceImpl) Save(memberCreateRequest web.MemberCreateOrUpdateRequest, memberID int) string {
	tx := s.Database.Begin()
	defer tx.Rollback()

	err := s.Validate.Struct(memberCreateRequest)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	err = web.PasswordValidator(memberCreateRequest.Password)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(memberCreateRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Error()
		panic(err)
	}

	email, _ := s.MembershipRepository.FindByEmail(tx, memberCreateRequest.Email)
	if email.Email != "" {
		panic(exception.NewBadRequestError("email has been already exists"))
	}

	member := domain.Membership{
		Name:          memberCreateRequest.Name,
		Email:         memberCreateRequest.Email,
		Password:      string(passwordHash),
		Phone:         memberCreateRequest.Phone,
		BirthDate:     util.ConvertStringToTime(memberCreateRequest.BirthDate),
		Address:       memberCreateRequest.Address,
		JoinDate:      time.Now(),
		Referral:      uuid.NewString(),
		EarnedPoint:   0,
		RedeemedPoint: 0,
		RemainedPoint: 0,
		StatusID:      1,
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	}
	idCreateMember := s.MembershipRepository.Save(tx, member)

	addMemberHistories := domain.AddMemberHistory{
		MemberID:        memberID,
		PersonID:        idCreateMember,
		TransactionDate: time.Now(),
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	}

	idCreateHistory := s.MembershipRepository.AddMemberHistory(tx, addMemberHistories)

	memberCreator, err := s.MembershipRepository.FindByID(tx, memberID)
	if err != nil {
		panic(exception.NewNotFoundError("member is not found"))
	}

	loyalties, _ := s.TierManagementRepository.FindByPoint(tx, memberCreator.EarnedPoint)
	if len(loyalties) > 0 {
		var earnedPoint int64
		loyaltyPolicy, errLoyalty := s.LoyaltyRepository.FindByID(tx, loyalties[0].ID)
		if errLoyalty == nil {
			earnedPoint += loyaltyPolicy.CommunityPolicy.FixedPoint
		}
		s.EarnedHistoryRepository.Save(tx, domain.EarnedPointHistory{
			MemberID:               memberID,
			TransactionDate:        time.Now(),
			ReferenceTransactionID: idCreateHistory,
			LoyaltyProgramID:       loyalties[0].ID,
			ExistingPoint:          memberCreator.RemainedPoint,
			EarnedPoint:            earnedPoint,
			BalancePoint:           member.RemainedPoint + earnedPoint,
			Timestamp: domain.Timestamp{
				CreatedAt: util.GetUnixTimestamp(),
			},
		})
		memberCreator.EarnedPoint += earnedPoint
		memberCreator.RemainedPoint += earnedPoint
		s.MembershipRepository.Save(tx, memberCreator)
	}
	if err == nil {
		tx.Commit()
	}
	return idCreateHistory
}

func (s MembershipServiceImpl) SignIn(request web.SignInRequest) string {
	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	member, err := s.MembershipRepository.FindByEmail(s.Database, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError("member is not found"))
	}

	//compare password has
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    member.ID,
		"email": member.Email,
		"phone": member.Phone,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	return t

}

func (s MembershipServiceImpl) AddMemberActivity(request web.AddMemberActivityRequest, memberID int) string {
	tx := s.Database.Begin()
	defer tx.Rollback()

	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	transactionID := s.MembershipRepository.AddMemberActivityHistory(tx, domain.AddMemberActivityHistory{
		MemberID:        memberID,
		ActivityName:    request.ActivityName,
		TransactionDate: time.Now(),
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	})

	member, err := s.MembershipRepository.FindByID(tx, memberID)
	if err != nil {
		panic(exception.NewNotFoundError("member is not found"))
	}

	loyalties, _ := s.TierManagementRepository.FindByPoint(tx, member.EarnedPoint)
	if len(loyalties) > 0 {
		var earnedPoint int64
		loyaltyPolicy, errLoyalty := s.LoyaltyRepository.FindByID(tx, loyalties[0].ID)
		if errLoyalty == nil {
			earnedPoint += loyaltyPolicy.CommunityPolicy.FixedPoint
		}
		s.EarnedHistoryRepository.Save(tx, domain.EarnedPointHistory{
			MemberID:               memberID,
			TransactionDate:        time.Now(),
			ReferenceTransactionID: transactionID,
			LoyaltyProgramID:       loyalties[0].ID,
			ExistingPoint:          member.RemainedPoint,
			EarnedPoint:            earnedPoint,
			BalancePoint:           member.RemainedPoint + earnedPoint,
			Timestamp: domain.Timestamp{
				CreatedAt: util.GetUnixTimestamp(),
			},
		})
		member.EarnedPoint += earnedPoint
		member.RemainedPoint += earnedPoint
		s.MembershipRepository.Save(tx, member)
	}

	if err == nil {
		tx.Commit()
	}

	return transactionID
}

func (s MembershipServiceImpl) AddRedeemedPoint(request web.RedeemedPointRequest, memberID int) {
	tx := s.Database.Begin()
	defer tx.Rollback()

	err := s.Validate.Struct(request)
	errTrans := util.TranslateErroValidation(s.Validate, err)
	if err != nil {
		log.Error(err)
		panic(exception.NewBadRequestError(errTrans.Error()))
	}

	member, err := s.MembershipRepository.FindByID(tx, memberID)
	if err != nil {
		log.Error(err)
		panic(exception.NewNotFoundError("user is not found"))
	}

	if member.RemainedPoint < request.RedeemedPoint {
		log.Error()
		panic(exception.NewBadRequestError("tidak bisa diredeem kalau redeemed poin lebih besar dari earned poin"))
	}

	s.RedeemedHistoryRepository.AddRedeemed(tx, domain.RedeemedPointHistory{
		MemberID:        memberID,
		EarnedPoint:     member.RemainedPoint,
		RedeemedPoint:   request.RedeemedPoint,
		RemainingPoint:  member.RemainedPoint - request.RedeemedPoint,
		TransactionDate: time.Now(),
		Timestamp: domain.Timestamp{
			CreatedAt: util.GetUnixTimestamp(),
		},
	})

	member.RedeemedPoint += request.RedeemedPoint
	member.RemainedPoint -= request.RedeemedPoint
	s.MembershipRepository.Save(tx, member)

	if err == nil {
		tx.Commit()
	}
}

func (s MembershipServiceImpl) FindAllRedeemedPointHistory(memberID int, filter web.FilterRedeemedPoint) []web.RedeemedPointHistoryResponse {
	histories := s.RedeemedHistoryRepository.FindAllByMemberID(s.Database, memberID, filter)
	return web.ToRedeemdPointHistoryResponses(histories)
}

func (s MembershipServiceImpl) FindAllEarnedPointHistory(memberID int, filter web.FilterEarnedPoint) []web.EarnedPointHistoryResponse {
	histories := s.EarnedHistoryRepository.FindAllEarnedHistory(s.Database, memberID, filter)
	return web.ToEarnedPointHistoryResponses(histories)
}
