package models

// JWSTransactionDecodedPayload
// A decoded payload that contains transaction information.
type JWSTransactionDecodedPayload struct {
	// A UUID you create at the time of purchase that associates the transaction with a customer on your own service.
	// If your app doesn’t provide an appAccountToken, this string is empty.
	AppAccountToken string `json:"appAccountToken"`
	// The bundle identifier of the app.
	BundleId string `json:"bundleId"`
	// The three-letter ISO 4217 currency code associated with the price parameter. This value is present only if price is present.
	Currency string `json:"currency"`
	// The server environment, either sandbox or production.
	Environment Environment `json:"environment"`
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
	PurchaseType PurchaseType `json:"type"`
	// The unique identifier of subscription purchase events across devices, including subscription renewals.
	WebOrderLineItemId string `json:"webOrderLineItemId"`
}

func (J *JWSTransactionDecodedPayload) Validate() error {
	return nil
}
