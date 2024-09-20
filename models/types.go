package models

// Environment
// The server environment, either sandbox or production.
type Environment = string

const (
	EnvProduct Environment = "Production"
	EnvSandbox Environment = "Sandbox"
)

// Status
// The status of an auto-renewable subscription.
type Status = int32

const (
	// StatusActive
	// The auto-renewable subscription is active.
	StatusActive Status = iota + 1
	// StatusExpired
	// The auto-renewable subscription is expired.
	StatusExpired
	// StatusBillingRetry
	// The auto-renewable subscription is in a billing retry period.
	StatusBillingRetry
	// StatusBillingGracePeriod
	// The auto-renewable subscription is in a Billing Grace Period.
	StatusBillingGracePeriod
	// StatusRevoked
	// The auto-renewable subscription is revoked. The App Store refunded the transaction or revoked it from Family Sharing.
	StatusRevoked
)

// AutoRenewStatus
// The renewal status for an auto-renewable subscription.
type AutoRenewStatus = int32

const (
	// AutoRenewStatusOff
	// Automatic renewal is off. The customer has turned off automatic renewal for the subscription, and it won’t renew at the end of the current subscription period.
	AutoRenewStatusOff AutoRenewStatus = iota
	// AutoRenewStatusOn
	// Automatic renewal is on. The subscription renews at the end of the current subscription period.
	AutoRenewStatusOn
)

// ExpirationIntent
// The reason a subscription expired.
type ExpirationIntent = int32

const (
	// ExpirationIntentCustomerCancelled
	// The customer canceled their subscription.
	ExpirationIntentCustomerCancelled ExpirationIntent = iota + 1
	// ExpirationIntentBillingError
	// Billing error; for example, the customer’s payment information is no longer valid.
	ExpirationIntentBillingError
	// ExpirationIntentCustomerDidNotConsentToPriceIncrease
	// The customer didn’t consent to an auto-renewable subscription price increase that requires customer consent, allowing the subscription to expire.
	ExpirationIntentCustomerDidNotConsentToPriceIncrease
	// ExpirationIntentProductNotAvailable
	// The product wasn’t available for purchase at the time of renewal.
	ExpirationIntentProductNotAvailable
	// ExpirationIntentOther
	// The subscription expired for some other reason.
	ExpirationIntentOther
)

// OfferDiscountType
// The payment mode for a subscription offer for an auto-renewable subscription.
type OfferDiscountType = string

const (
	// OfferDiscountTypeFreeTrial
	// A payment mode of a product discount that indicates a free trial.
	OfferDiscountTypeFreeTrial OfferDiscountType = "FREE_TRIAL"
	// OfferDiscountTypePayAsYouGo
	// A payment mode of a product discount that customers pay over a single or multiple billing periods.
	OfferDiscountTypePayAsYouGo OfferDiscountType = "PAY_AS_YOU_GO"
	// OfferDiscountTypePayUpFront
	// A payment mode of a product discount that customers pay up front.
	OfferDiscountTypePayUpFront OfferDiscountType = "PAY_UP_FRONT"
)

// OfferType
// The type of subscription offer.
type OfferType = int32

const (
	// OfferTypeIntroductoryOffer
	// An introductory offer.
	OfferTypeIntroductoryOffer OfferType = iota + 1
	// OfferTypePromotionalOffer
	// A promotional offer.
	OfferTypePromotionalOffer
	// OfferTypeSubscriptionOfferCode
	// An offer with a subscription offer code.
	OfferTypeSubscriptionOfferCode
	// OfferTypeWinBackOffer
	// A win-back offer.
	OfferTypeWinBackOffer
)

type PriceIncreaseStatus = int32

const (
	// PriceIncreaseStatusCustomerHasNotResponded
	// The customer hasn’t yet responded to an auto-renewable subscription price increase that requires customer consent.
	PriceIncreaseStatusCustomerHasNotResponded PriceIncreaseStatus = iota
	// PriceIncreaseStatusCustomerConsentedOrWasNotifiedWithoutNeedingConsent
	// The customer consented to an auto-renewable subscription price increase that requires customer consent, or the App Store has notified the customer of an auto-renewable subscription price increase that doesn’t require consent.
	PriceIncreaseStatusCustomerConsentedOrWasNotifiedWithoutNeedingConsent
)

// PurchaseType
// The type of In-App Purchase products you can offer in your app.
type PurchaseType = string

const (
	// PurchaseTypeAutoRenewableSubscription
	// An auto-renewable subscription
	PurchaseTypeAutoRenewableSubscription PurchaseType = "Auto-Renewable Subscription"
	// PurchaseTypeNonConsumable
	// A non-consumable In-App Purchase
	PurchaseTypeNonConsumable PurchaseType = "Non-Consumable"
	// PurchaseTypeConsumable
	// A consumable In-App Purchase
	PurchaseTypeConsumable PurchaseType = "Consumable"
	// PurchaseTypeNonRenewingSubscription
	// A non-renewing subscription
	PurchaseTypeNonRenewingSubscription PurchaseType = "Non-Renewing Subscription"
)
