package models

import "github.com/meetleev/go-apple-store-server/types"

// LastTransactionsItem
// The most recent App Store-signed transaction information and App Store-signed renewal information for an auto-renewable subscription.
type LastTransactionsItem struct {
	// The original transaction identifier of the auto-renewable subscription.
	OriginalTransactionId string `json:"originalTransactionId"`
	// The status of the auto-renewable subscription.
	Status types.Status `json:"status"`
	// The subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedRenewalInfo string `json:"signedRenewalInfo"`
	// The transaction information signed by the App Store, in JWS format.
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}
