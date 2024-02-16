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
	Database             *gorm.DB
	MembershipRepository repository.MembershipRepositoryInterface
	Validate             *validator.Validate
}

func NewMembershipService(Database *gorm.DB,
	MembershipRepository repository.MembershipRepositoryInterface,
	Validate *validator.Validate) MembershipServiceInterface {
	return &MembershipServiceImpl{
		Database:             Database,
		MembershipRepository: MembershipRepository,
		Validate:             Validate,
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

func (s MembershipServiceImpl) Save(memberCreateRequest web.MemberCreateOrUpdateRequest) int {
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

	email, _ := s.MembershipRepository.FindByEmail(s.Database, memberCreateRequest.Email)
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

	return s.MembershipRepository.Save(s.Database, member)
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
