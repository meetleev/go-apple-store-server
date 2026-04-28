package models

import "github.com/meetleev/go-apple-store-server/types"

// AppTransactionDecodedPayload
// A decoded payload that contains app transaction information.
type AppTransactionDecodedPayload struct {
	// The unique identifier the App Store uses to identify the app.
	AppAppleId uint64 `json:"appAppleId"`
	// The unique identifier of the app download transaction.
	AppTransactionID string `json:"appTransactionID"`
	// The bundle identifier of the app.
	BundleId string `json:"bundleID"`

	// The app version that the customer originally purchased from the App Store.
	OriginalAppVersion string `json:"originalApplicationVersion"`
	// The platform on which the customer originally purchased the app.
	OriginalPlatform string `json:"originalPlatform"`
	// The UNIX time, in milliseconds, that the customer originally purchased the app from the App Store.
	OriginalPurchaseDate int64 `json:"originalPurchaseDate"`

	// The date the customer placed an order for the app before it’s available in the App Store.
	PreorderDate int64 `json:"preorderDate"`
	// The date that the App Store signed the JWS app transaction.
	ReceiptCreationDate int64 `json:"receiptCreationDate"`
	// The server environment, either sandbox or production.
	Environment types.Environment `json:"environment"`

	// The app version.
	AppVersion string `json:"appVersion"`
	// The number that the App Store uses to uniquely identify the version of the app.
	AppVersionID uint64 `json:"appVersionID"`

	// The device verification value used to verify whether the app transaction belongs to the device.
	DeviceVerification []byte `json:"deviceVerification"`
	// The UUID used to compute the device verification value.
	DeviceVerificationNonce string `json:"deviceVerificationNonce"`
	// The UNIX time, in milliseconds, that the App Store signed the JSON Web Signature (JWS) data.
	SignedDate int64 `json:"signedDate"`
}

func (a *AppTransactionDecodedPayload) Validate() error {
	return nil
}

func (a *AppTransactionDecodedPayload) BundleID() string {
	if a == nil {
		return ""
	}
	return a.BundleId
}

func (a *AppTransactionDecodedPayload) EnvironmentValue() string {
	if a == nil {
		return ""
	}
	return string(a.Environment)
}
