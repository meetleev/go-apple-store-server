package models

import (
	"errors"
	"github.com/meetleev/go-apple-store-server/verifier"
	logger "github.com/sirupsen/logrus"
)

// LastTransactionsItem
// The most recent App Store-signed transaction information and App Store-signed renewal information for an auto-renewable subscription.
type LastTransactionsItem struct {
	// The original transaction identifier of the auto-renewable subscription.
	OriginalTransactionId string `json:"originalTransactionId"`
	// The status of the auto-renewable subscription.
	Status Status `json:"status"`
	// The subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedRenewalInfo string `json:"signedRenewalInfo"`
	// The transaction information signed by the App Store, in JWS format.
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

func (t *LastTransactionsItem) DecodeTransactionInfo() (*JWSTransactionDecodedPayload, error) {
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

func (t *LastTransactionsItem) DecodeRenewalInfo() (*JWSRenewalInfoDecodedPayload, error) {
	singedVerifier := verifier.NewParserWithDefault()
	payload := &JWSRenewalInfoDecodedPayload{}
	v, err := singedVerifier.Parse(t.SignedTransactionInfo, payload)
	logger.Debugf("JWSRenewalInfoDecodedPayload %v", v)
	if nil != err {
		logger.Errorf("singed parse error:[%v]", err.Error())
		return nil, err
	}
	if !v.Valid {
		errMsg := "JWSRenewalInfoDecodedPayload singed verify failed"
		logger.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	return payload, nil
}
