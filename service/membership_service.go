package service

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
)

type MembershipServiceInterface interface {
	FindByID(ID int) domain.Membership
	FindByEmail(email string) domain.Membership
	Save(memberCreateRequest web.MemberCreateOrUpdateRequest, memberID int) string
	AddMemberActivity(request web.AddMemberActivityRequest, memberID int) string
	SignIn(request web.SignInRequest) string
	AddRedeemedPoint(request web.RedeemedPointRequest, memberID int)
	FindAllRedeemedPointHistory(memberID int) []web.RedeemedPointHistoryResponse
	FindAllEarnedPointHistory(memberID int) []web.EarnedPointHistoryResponse
}
