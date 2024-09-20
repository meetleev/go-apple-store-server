package models

// SubscriptionGroupIdentifierItem
// Information for auto-renewable subscriptions, including signed transaction information and signed renewal information, for one subscription group.
type SubscriptionGroupIdentifierItem struct {
	// The subscription group identifier of the auto-renewable subscriptions in the lastTransactions array.
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	// An array of the most recent App Store-signed transaction information and App Store-signed renewal information for all auto-renewable subscriptions in the subscription group.
	LastTransactions []*LastTransactionsItem `json:"lastTransactions"`
}
