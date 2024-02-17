package web

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"strings"
	"time"
)

type SignInRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}

type RedeemedPointRequest struct {
	RedeemedPoint int64 `validate:"required,gte=1000" json:"redeemed_point"`
}
type MemberCreateOrUpdateRequest struct {
	Name      string `validate:"required" json:"name"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
	Phone     string `validate:"required,e164" json:"phone"`
	BirthDate string `validate:"required,datetime=2006-01-02" json:"birth_date"`
	Address   string `json:"address"`
}

type AddMemberActivityRequest struct {
	ActivityName string `validate:"required,min=5" json:"activity_name"`
}

type RedeemedPointHistoryResponse struct {
	ID              int    `json:"id"`
	MemberID        int    `json:"member_id"`
	EarnedPoint     int64  `json:"earned_point"`
	RedeemedPoint   int64  `json:"redeemed_point"`
	RemainingPoint  int64  `json:"remaining_point"`
	TransactionDate string `json:"transaction_date"`
}

func ToRedeemedPointHistoryResponse(redeemed domain.RedeemedPointHistory) RedeemedPointHistoryResponse {
	return RedeemedPointHistoryResponse{
		ID:              redeemed.ID,
		MemberID:        redeemed.MemberID,
		EarnedPoint:     redeemed.EarnedPoint,
		RedeemedPoint:   redeemed.RedeemedPoint,
		RemainingPoint:  redeemed.RemainingPoint,
		TransactionDate: redeemed.TransactionDate.String(),
	}
}

func ToRedeemdPointHistoryResponses(redeemedPoints []domain.RedeemedPointHistory) []RedeemedPointHistoryResponse {
	var redeemedPointResponses []RedeemedPointHistoryResponse

	for _, redeemed := range redeemedPoints {
		redeemedPointResponse := ToRedeemedPointHistoryResponse(redeemed)
		redeemedPointResponses = append(redeemedPointResponses, redeemedPointResponse)
	}
	return redeemedPointResponses
}

type EarnedPointHistoryResponse struct {
	ID                     int       `json:"id"`
	MemberID               int       `json:"member_id"`
	TransactionDate        time.Time `json:"transaction_date"`
	ReferenceTransactionID string    `json:"reference_transaction_id"`
	LoyaltyProgramID       int       `json:"loyalty_program_id"`
	ExistingPoint          int64     `json:"existing_point"`
	EarnedPoint            int64     `json:"earned_point"`
	BalancePoint           int64     `json:"balance_point"`
}

func ToEarnedPointHistoryResponse(earned domain.EarnedPointHistory) EarnedPointHistoryResponse {
	return EarnedPointHistoryResponse{
		ID:                     earned.ID,
		MemberID:               earned.MemberID,
		TransactionDate:        earned.TransactionDate,
		ReferenceTransactionID: earned.ReferenceTransactionID,
		LoyaltyProgramID:       earned.LoyaltyProgramID,
		ExistingPoint:          earned.ExistingPoint,
		EarnedPoint:            earned.EarnedPoint,
		BalancePoint:           earned.BalancePoint,
	}
}

func ToEarnedPointHistoryResponses(earneds []domain.EarnedPointHistory) []EarnedPointHistoryResponse {
	var earnedPointHistories []EarnedPointHistoryResponse

	for _, earned := range earneds {
		earnedPointHistories = append(earnedPointHistories, ToEarnedPointHistoryResponse(earned))
	}

	return earnedPointHistories
}

func containSpecialCharacter(str string, specialCharacter []string) bool {
	for _, sc := range specialCharacter {
		if strings.Contains(str, sc) {
			return true
		}
	}

	return false
}

func PasswordValidator(password string) error {
	validate := validator.New()

	err := validate.Var(password, "containsany=abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		return errors.New("password must contain lowercase")
	}

	err = validate.Var(password, "containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		return errors.New("password must contain uppercase")
	}

	listSpecialCharacter := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "'", "\"", "/", "%", "?", "\n", "<", ">"}
	symbol := containSpecialCharacter(password, listSpecialCharacter)
	if !symbol {
		return errors.New("password must contain special character")
	}

	return nil
}
