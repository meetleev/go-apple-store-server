package verifier

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

const rootCaBase64Encoded = "MIICQzCCAcmgAwIBAgIILcX8iNLFS5UwCgYIKoZIzj0EAwMwZzEbMBkGA1UEAwwSQXBwbGUgUm9vdCBDQSAtIEczMSYwJAYDVQQLDB1BcHBsZSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTETMBEGA1UECgwKQXBwbGUgSW5jLjELMAkGA1UEBhMCVVMwHhcNMTQwNDMwMTgxOTA2WhcNMzkwNDMwMTgxOTA2WjBnMRswGQYDVQQDDBJBcHBsZSBSb290IENBIC0gRzMxJjAkBgNVBAsMHUFwcGxlIENlcnRpZmljYXRpb24gQXV0aG9yaXR5MRMwEQYDVQQKDApBcHBsZSBJbmMuMQswCQYDVQQGEwJVUzB2MBAGByqGSM49AgEGBSuBBAAiA2IABJjpLz1AcqTtkyJygRMc3RCV8cWjTnHcFBbZDuWmBSp3ZHtfTjjTuxxEtX/1H7YyYl3J6YRbTzBPEVoA/VhYDKX1DyxNB0cTddqXl5dvMVztK517IDvYuVTZXpmkOlEKMaNCMEAwHQYDVR0OBBYEFLuw3qFYM4iapIqZ3r6966/ayySrMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMAoGCCqGSM49BAMDA2gAMGUCMQCD6cHEFl4aXTQY2e3v9GwOAEZLuN+yRhHFD/3meoyhpmvOwgPUnPWTxnS4at+qIxUCMG1mihDK1A3UT82NQz60imOlM27jbdoXt2QfyFMm+YhidDkLF1vLUagM6BgD56KyKA=="

// CertFromBase64 loads a .X509 certificate from base64Cert and
// returns an *x509.Certificate
func CertFromBase64(base64Cert string) ([]*x509.Certificate, error) {
	certBytes, err := base64.StdEncoding.DecodeString(base64Cert)
	if err != nil {
		return nil, err
	}
	return x509.ParseCertificates(certBytes)
}

type JWTSignData struct {
	Raw       string                 // Raw contains the raw token.  Populated when you [Parse] a token
	Method    jwt.SigningMethod      // Method is the signing method used or to be used
	Header    map[string]interface{} // Header is the first segment of the token in decoded form
	Payload   interface{}            // Payload is the second segment of the Payload in decoded form
	Signature []byte                 // Signature is the third segment of the token in decoded form.  Populated when you Parse a token
	Valid     bool
}

type SignedDataVerifier struct {
	rootCertificates []*x509.Certificate

	parse *jwt.Parser
}

func NewParser(rootCertificates []*x509.Certificate) *SignedDataVerifier {
	o := &SignedDataVerifier{rootCertificates: rootCertificates}
	o.parse = jwt.NewParser()
	return o
}

func NewParserWithDefault() *SignedDataVerifier {
	rootCertificates, _ := CertFromBase64(rootCaBase64Encoded)
	o := &SignedDataVerifier{rootCertificates: rootCertificates}
	o.parse = jwt.NewParser()
	return o
}

func (p *SignedDataVerifier) Parse(tokenString string, payload interface{}) (*JWTSignData, error) {
	token, parts, err := p.parseUnverified(tokenString, payload)
	if err != nil {
		return token, err
	}

	// Decode signature
	token.Signature, err = p.parse.DecodeSegment(parts[2])
	if err != nil {
		return token, newError("could not base64 decode signature", jwt.ErrTokenMalformed, err)
	}
	text := strings.Join(parts[0:2], ".")

	sChain, ok := token.Header["x5c"]
	if !ok {
		return token, errors.New("invalid cert")
	}
	chain, ok := sChain.([]interface{})
	if !ok || 3 != len(chain) {
		return token, errors.New("invalid chain length")
	}

	var certificateChain []*x509.Certificate
	for _, v := range chain {
		s, ok := v.(string)
		if !ok {
			return token, errors.New("cert is not string")
		}
		cert, err := CertFromBase64(s)
		if nil != err {
			return token, err
		}
		certificateChain = append(certificateChain, cert...)
	}
	key, err := p.verifyCertificateChain(certificateChain[0], certificateChain[1])
	if err != nil {
		return token, err
	}

	err = token.Method.Verify(text, token.Signature, key)

	if err != nil {
		return token, newError("", jwt.ErrTokenSignatureInvalid, err)
	}

	// No errors so far, token is valid.
	token.Valid = true

	return token, nil
}

func (p *SignedDataVerifier) parseUnverified(data string, payload interface{}) (token *JWTSignData, parts []string, err error) {
	parts = strings.Split(data, ".")
	if len(parts) != 3 {
		return nil, parts, newError("data contains an invalid number of segments", jwt.ErrTokenMalformed)
	}

	token = &JWTSignData{Raw: data}

	// parse Header
	var headerBytes []byte
	if headerBytes, err = p.parse.DecodeSegment(parts[0]); err != nil {
		return token, parts, newError("could not base64 decode header", jwt.ErrTokenMalformed, err)
	}
	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
		return token, parts, newError("could not JSON decode header", jwt.ErrTokenMalformed, err)
	}

	// parse Payload
	token.Payload = payload

	claimBytes, err := p.parse.DecodeSegment(parts[1])
	if err != nil {
		return token, parts, newError("could not base64 decode claim", jwt.ErrTokenMalformed, err)
	}

	err = json.Unmarshal(claimBytes, &payload)

	if err != nil {
		return token, parts, newError("could not JSON decode payload", jwt.ErrTokenMalformed, err)
	}

	// Lookup signature method
	if method, ok := token.Header["alg"].(string); ok {
		if token.Method = jwt.GetSigningMethod(method); token.Method == nil {
			return token, parts, newError("signing method (alg) is unavailable", jwt.ErrTokenUnverifiable)
		}
	} else {
		return token, parts, newError("signing method (alg) is unspecified", jwt.ErrTokenUnverifiable)
	}

	return token, parts, nil
}

func (p *SignedDataVerifier) verifyCertificateChain(leaf *x509.Certificate, intermediate *x509.Certificate) (*ecdsa.PublicKey, error) {
	rootCAs := x509.NewCertPool()
	for _, root := range p.rootCertificates {
		rootCAs.AddCert(root)
	}
	intermediateCAs := x509.NewCertPool()
	intermediateCAs.AddCert(intermediate)

	_, err := leaf.Verify(x509.VerifyOptions{Roots: rootCAs, Intermediates: intermediateCAs})
	pubKey := (leaf.PublicKey).(*ecdsa.PublicKey)
	return pubKey, err
}

func newError(message string, err error, more ...error) error {
	var format string
	var args []any
	if message != "" {
		format = "%w: %s"
		args = []any{err, message}
	} else {
		format = "%w"
		args = []any{err}
	}

	for _, e := range more {
		format += ": %w"
		args = append(args, e)
	}

	err = fmt.Errorf(format, args...)
	return err
}
