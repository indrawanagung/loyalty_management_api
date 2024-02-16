package web

import "github.com/indrawanagung/loyalty_management_api/model/domain"

type TierManagementResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	MinPoint int64  `json:"min_point"`
	MaxPoint int64  `json:"max_point"`
}

type TierManagementRequest struct {
	Name     string `validate:"required" json:"name"`
	MinPoint int64  `validate:"required,number" json:"min_point"`
	MaxPoint int64  `validate:"required,number" json:"max_point"`
}

func ToTierManagementResponse(tier domain.TierManagement) TierManagementResponse {
	return TierManagementResponse{
		ID:       tier.ID,
		Name:     tier.Name,
		MinPoint: tier.MinPoint,
		MaxPoint: tier.MaxPoint,
	}
}

func ToTierManagementResponses(tiers []domain.TierManagement) []TierManagementResponse {
	var tierManagementResponses []TierManagementResponse

	for _, tierManagement := range tiers {
		tierManagementResponses = append(tierManagementResponses, ToTierManagementResponse(tierManagement))
	}
	return tierManagementResponses
}
