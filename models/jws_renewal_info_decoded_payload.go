package models

import "github.com/meetleev/go-apple-store-server/types"

type renewalCommitmentInfo struct {
	AutoRenewProductId     string                `json:"commitmentAutoRenewProductId"`
	AutoRenewStatus        types.AutoRenewStatus `json:"commitmentAutoRenewStatus"`
	RenewalBillingPlanType string                `json:"commitmentRenewalBillingPlanType,omitempty"`
	RenewalDate            int64                 `json:"commitmentRenewalDate,omitempty"`
	RenewalPrice           int64                 `json:"commitmentRenewalPrice,omitempty"`
}

type advancedCommerceRenewalItem struct {
	SKU               string                             `json:"SKU"`
	Description       string                             `json:"description"`
	DisplayName       string                             `json:"displayName"`
	Offer             *advancedCommerceOffer             `json:"offer"`
	Price             int64                              `json:"price"`
	PriceIncreaseInfo *advancedCommercePriceIncreaseInfo `json:"priceIncreaseInfo,omitempty"`
}

type advancedCommerceRenewalInfo struct {
	ConsistencyToken   string                         `json:"consistencyToken"`
	Descriptors        *advancedCommerceDescriptors   `json:"descriptors,omitempty"`
	Items              []*advancedCommerceRenewalItem `json:"items,omitempty"`
	Period             string                         `json:"period,omitempty"`
	RequestReferenceId string                         `json:"requestReferenceId,omitempty"`
	TaxCode            string                         `json:"taxCode,omitempty"`
}

// JWSRenewalInfoDecodedPayload
// A decoded payload containing subscription renewal information for an auto-renewable subscription.
type JWSRenewalInfoDecodedPayload struct {
	// A UUID you create at the time of purchase that associates the transaction with a customer on your own service.
	// If your app doesn’t provide an appAccountToken, this string is empty.
	AppAccountToken string `json:"appAccountToken"`
	// The unique identifier of the app download transaction.
	AppTransactionId string `json:"appTransactionId"`
	// The product identifier of the product that renews at the next billing period.
	AutoRenewProductId string `json:"autoRenewProductId"`
	// The renewal status of the auto-renewable subscription.
	AutoRenewStatus types.AutoRenewStatus `json:"autoRenewStatus"`
	// The currency code for the renewalPrice of the subscription.
	Currency string `json:"currency"`
	// The list of win-back offer IDs that the customer is eligible for.
	EligibleWinBackOfferIds []string `json:"eligibleWinBackOfferIds"`
	// The server environment, either sandbox or production.
	Environment types.Environment `json:"environment"`
	// 	The reason the subscription expired.
	ExpirationIntent types.ExpirationIntent `json:"expirationIntent"`
	// 	The time when the Billing Grace Period for subscription renewals expires.
	GracePeriodExpiresDate int64 `json:"gracePeriodExpiresDate"`
	// A Boolean value that indicates whether the App Store is attempting to automatically renew a subscription that expired due to a billing issue.
	IsInBillingRetryPeriod bool `json:"isInBillingRetryPeriod"`
	// The payment mode of the discount offer.
	OfferDiscountType types.OfferDiscountType `json:"offerDiscountType"`
	// The offer code or the promotional offer identifier.
	OfferIdentifier string `json:"offerIdentifier"`
	// The duration of the offer. This field is in ISO 8601 duration format.
	OfferPeriod string `json:"offerPeriod"`
	// 	The type of subscription offer.
	OfferType types.OfferType `json:"offerType"`
	// The transaction identifier of the original purchase associated with this transaction.
	OriginalTransactionId string `json:"originalTransactionId"`
	// The status that indicates whether the auto-renewable subscription is subject to a price increase.
	PriceIncreaseStatus types.PriceIncreaseStatus `json:"priceIncreaseStatus"`
	//	The product identifier of the In-App Purchase.
	ProductId string `json:"productId"`
	// The earliest start date of the auto-renewable subscription in a series of subscription purchases that ignores all lapses of paid service that are 60 days or fewer.
	RecentSubscriptionStartDate int64 `json:"recentSubscriptionStartDate"`
	// The UNIX time, in milliseconds, when the most recent auto-renewable subscription purchase expires.
	RenewalDate int64 `json:"renewalDate"`
	// The renewal price, in milli units, of the auto-renewable subscription that renews at the next billing period.
	RenewalPrice int64 `json:"renewalPrice"`
	// The UNIX time, in milliseconds, that the App Store signed the JSON Web Signature (JWS) data.
	SignedDate int64 `json:"signedDate"`

	AdvancedCommerceInfo *advancedCommerceRenewalInfo `json:"advancedCommerceInfo,omitempty"`
	CommitmentInfo       *renewalCommitmentInfo       `json:"commitmentInfo,omitempty"`
	/*
		Possible Values:
		BILLED_UPFRONT
		MONTHLY
	*/
	RenewalBillingPlanType string `json:"renewalBillingPlanType,omitempty"`
}

func (J *JWSRenewalInfoDecodedPayload) EnvironmentValue() string {
	return string(J.Environment)
}
