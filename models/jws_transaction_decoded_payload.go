package models

import "github.com/meetleev/go-apple-store-server/types"

type RevocationType = string

const (
	// The transaction has a full refund.
	RevocationTypeRefundFull RevocationType = "REFUND_FULL"
	// The transaction has a prorated refund.
	RevocationTypeRefundProrated RevocationType = "REFUND_PRORATED"
	// The transaction is revoked from Family Sharing.
	RevocationTypeFamilyRevoked RevocationType = "FAMILY_REVOKED"
)

type transactionCommitmentInfo struct {
	// Minimum: 1 Maximum: 12
	BillingPeriodNumber   int32 `json:"billingPeriodNumber"`
	CommitmentExpiresDate int64 `json:"commitmentExpiresDate"`
	CommitmentPrice       int64 `json:"commitmentPrice"`
	TotalBillingPeriods   int32 `json:"totalBillingPeriods"`
}

type advancedCommerceRefund struct {
	RefundAmount int64 `json:"refundAmount"`
	RefundDate   int64 `json:"refundDate"`
	/*
		Possible Values
		UNINTENDED_PURCHASE
		FULFILLMENT_ISSUE
		UNSATISFIED_WITH_PURCHASE
		LEGAL
		OTHER
		MODIFY_ITEMS_REFUND
		SIMULATE_REFUND_DECLINE
	*/
	RefundReason string `json:"refundReason"`
	/*
		Possible Values:
		FULL
		PRORATED
		CUSTOM
	*/
	RefundType string `json:"refundType"`
}

type advancedCommerceTransactionItem struct {
	SKU            string                    `json:"SKU"`
	Description    string                    `json:"description"`
	DisplayName    string                    `json:"displayName"`
	Offer          *advancedCommerceOffer    `json:"offer"`
	Price          int64                     `json:"price"`
	Refunds        []*advancedCommerceRefund `json:"refunds"`
	RevocationDate int64                     `json:"revocationDate"`
}
type advancedCommerceTransactionInfo struct {
	Descriptors        *advancedCommerceDescriptors       `json:"descriptors"`
	EstimatedTax       int64                              `json:"estimatedTax"`
	Items              []*advancedCommerceTransactionItem `json:"items"`
	Period             string                             `json:"period"`
	RequestReferenceId string                             `json:"requestReferenceId"`
	TaxCode            string                             `json:"taxCode"`
	TaxExclusivePrice  int64                              `json:"taxExclusivePrice"`
	TaxRate            string                             `json:"taxRate"`
}

// JWSTransactionDecodedPayload
// A decoded payload that contains transaction information.
type JWSTransactionDecodedPayload struct {
	// A UUID you create at the time of purchase that associates the transaction with a customer on your own service.
	// If your app doesn’t provide an appAccountToken, this string is empty.
	AppAccountToken string `json:"appAccountToken"`
	// The unique identifier of the app download transaction.
	AppTransactionId string `json:"appTransactionId"`
	// The bundle identifier of the app.
	BundleId string `json:"bundleId"`
	// The three-letter ISO 4217 currency code associated with the price parameter. This value is present only if price is present.
	Currency string `json:"currency"`
	// The server environment, either sandbox or production.
	Environment types.Environment `json:"environment"`
	// The UNIX time, in milliseconds, that the subscription expires or renews.
	ExpiresDate int64 `json:"expiresDate"`
	// A string that describes whether the transaction was purchased by the customer, or is available to them through Family Sharing.
	InAppOwnershipType string `json:"inAppOwnershipType"`
	// A Boolean value that indicates whether the customer upgraded to another subscription.
	IsUpgraded bool `json:"isUpgraded"`
	// The payment mode you configure for the subscription offer, such as Free Trial, Pay As You Go, or Pay Up Front.
	OfferDiscountType string `json:"offerDiscountType"`
	// The identifier that contains the offer code or the promotional offer identifier.
	OfferIdentifier string `json:"offerIdentifier"`
	// The duration of the offer applied to the transaction. This field is in ISO 8601 duration format.
	OfferPeriod string `json:"offerPeriod"`
	// A value that represents the promotional offer type.
	OfferType int32 `json:"offerType"`
	// The UNIX time, in milliseconds, that represents the purchase date of the original transaction identifier.
	OriginalPurchaseDate int64 `json:"originalPurchaseDate"`
	// The transaction identifier of the original purchase.
	OriginalTransactionId string `json:"originalTransactionId"`
	// An integer value that represents the price multiplied by 1000 of the in-app purchase or subscription offer you configured in App Store Connect and that the system records at the time of the purchase.
	Price int64 `json:"price"`
	// The unique identifier of the product.
	ProductId string `json:"productId"`
	// The UNIX time, in milliseconds, that the App Store charged the customer’s account for a purchase, restored product, subscription, or subscription renewal after a lapse.
	PurchaseDate int64 `json:"purchaseDate"`
	// The number of consumable products the customer purchased.
	Quantity int32 `json:"quantity"`
	// The UNIX time, in milliseconds, that the App Store refunded the transaction or revoked it from Family Sharing.
	RevocationDate int64 `json:"revocationDate"`
	// The reason that the App Store refunded the transaction or revoked it from Family Sharing.
	RevocationReason int32 `json:"revocationReason"`
	// The type of the refund or revocation that applies to the transaction.
	RevocationType RevocationType `json:"revocationType"`
	// The UNIX time, in milliseconds, that the App Store signed the JSON Web Signature (JWS) data.
	SignedDate int64 `json:"signedDate"`
	// The three-letter code that represents the country or region associated with the App Store storefront for the purchase.
	Storefront string `json:"storefront"`
	// An Apple-defined value that uniquely identifies the App Store storefront associated with the purchase.
	StorefrontId string `json:"storefrontId"`
	// The identifier of the subscription group to which the subscription belongs.
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	// The unique identifier of the transaction.
	TransactionId string `json:"transactionId"`
	// The reason for the purchase transaction, which indicates whether it’s a customer’s purchase or a renewal for an auto-renewable subscription that the system initates.
	TransactionReason string `json:"transactionReason"`
	// The type of the in-app purchase.
	PurchaseType types.PurchaseType `json:"type"`
	// The unique identifier of subscription purchase events across devices, including subscription renewals.
	WebOrderLineItemId string `json:"webOrderLineItemId"`
	// Transaction information that is present only for Advanced Commerce SKUs.
	AdvancedCommerceInfo *advancedCommerceTransactionInfo `json:"advancedCommerceInfo,omitempty"`
	// Possible Values: BILLED_UPFRONT, MONTHLY
	BillingPlanType string `json:"billingPlanType,omitempty"`

	CommitmentInfo *transactionCommitmentInfo `json:"commitmentInfo,omitempty"`
}

func (J *JWSTransactionDecodedPayload) Validate() error {
	return nil
}

func (J *JWSTransactionDecodedPayload) BundleID() string {
	return J.BundleId
}

func (J *JWSTransactionDecodedPayload) EnvironmentValue() string {
	return string(J.Environment)
}
