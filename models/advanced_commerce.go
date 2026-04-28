package models

import "github.com/meetleev/go-apple-store-server/types"

type advancedCommercePriceIncreaseInfo struct {
	DependentSKUs []string                                      `json:"dependentSKUs"`
	Price         int64                                         `json:"price"`
	Status        types.AdvancedCommercePriceIncreaseInfoStatus `json:"status"`
}

type advancedCommerceDescriptors struct {
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type advancedCommerceOffer struct {
	Period      string `json:"period"`
	PeriodCount int32  `json:"periodCount"`
	Price       int64  `json:"price"`
	// Possible Values: ACQUISITION WIN_BACK RETENTION
	Reason string `json:"reason"`
}
