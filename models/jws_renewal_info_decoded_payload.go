package models

// JWSRenewalInfoDecodedPayload
// A decoded payload containing subscription renewal information for an auto-renewable subscription.
type JWSRenewalInfoDecodedPayload struct {
	// The product identifier of the product that renews at the next billing period.
	AutoRenewProductId string `json:"autoRenewProductId"`
	// The renewal status of the auto-renewable subscription.
	AutoRenewStatus AutoRenewStatus `json:"autoRenewStatus"`
	// The currency code for the renewalPrice of the subscription.
	Currency string `json:"currency"`
	// The list of win-back offer IDs that the customer is eligible for.
	EligibleWinBackOfferIds []string `json:"eligibleWinBackOfferIds"`
	// The server environment, either sandbox or production.
	Environment Environment `json:"environment"`
	// 	The reason the subscription expired.
	ExpirationIntent ExpirationIntent `json:"expirationIntent"`
	// 	The time when the Billing Grace Period for subscription renewals expires.
	GracePeriodExpiresDate int64 `json:"gracePeriodExpiresDate"`
	// A Boolean value that indicates whether the App Store is attempting to automatically renew a subscription that expired due to a billing issue.
	IsInBillingRetryPeriod bool `json:"isInBillingRetryPeriod"`
	// The payment mode of the discount offer.
	OfferDiscountType OfferDiscountType `json:"offerDiscountType"`
	// The offer code or the promotional offer identifier.
	OfferIdentifier string `json:"offerIdentifier"`
	// 	The type of subscription offer.
	OfferType OfferType `json:"offerType"`
	// The transaction identifier of the original purchase associated with this transaction.
	OriginalTransactionId string `json:"originalTransactionId"`
	// The status that indicates whether the auto-renewable subscription is subject to a price increase.
	PriceIncreaseStatus PriceIncreaseStatus `json:"priceIncreaseStatus"`
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
}
