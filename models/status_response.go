package models

// StatusResponse
// A response that contains status information for all of a customer’s auto-renewable subscriptions in your app.
type StatusResponse struct {
	// An array of information for auto-renewable subscriptions, including App Store-signed transaction information and App Store-signed renewal information.
	Data []*SubscriptionGroupIdentifierItem `json:"data"`
	// The server environment, sandbox or production, in which the App Store generated the response.
	Environment Environment `json:"environment"`
	// The unique identifier of an app in the App Store.
	AppAppleId int64 `json:"appAppleId"`
	// The bundle identifier of an app.
	BundleId string `json:"bundleId"`
}
