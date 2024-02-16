package service

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
)

type MembershipServiceInterface interface {
	FindByID(ID int) domain.Membership
	FindByEmail(email string) domain.Membership
	Save(memberCreateRequest web.MemberCreateOrUpdateRequest) int
	SignIn(request web.SignInRequest) string
}
