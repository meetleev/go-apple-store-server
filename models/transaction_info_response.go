package models

import (
	"errors"
	"github.com/meetleev/go-apple-store-server/verifier"
	logger "github.com/sirupsen/logrus"
)

// TransactionInfoResponse
// A response that contains signed transaction information for a single transaction.
type TransactionInfoResponse struct {
	// A customer’s in-app purchase transaction, signed by Apple, in JSON Web Signature (JWS) format.
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

func (t *TransactionInfoResponse) DecodeTransactionInfo() (*JWSTransactionDecodedPayload, error) {
	singedVerifier := verifier.NewParserWithDefault()
	payload := &JWSTransactionDecodedPayload{}
	v, err := singedVerifier.Parse(t.SignedTransactionInfo, payload)
	logger.Debugf("JWSTransactionDecodedPayload %v", v)
	if nil != err {
		logger.Errorf("singed parse error:[%v]", err.Error())
		return nil, err
	}
	if !v.Valid {
		errMsg := "JWSTransactionDecodedPayload singed verify failed"
		logger.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	return payload, nil
}
