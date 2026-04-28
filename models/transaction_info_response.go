package models

// TransactionInfoResponse
// A response that contains signed transaction information for a single transaction.
type TransactionInfoResponse struct {
	// A customer’s in-app purchase transaction, signed by Apple, in JSON Web Signature (JWS) format.
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}
