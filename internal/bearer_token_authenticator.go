package internal

import (
	"crypto/ecdsa"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const appStoreConnectAudience = "appstoreconnect-v1"

type authenticatorClaims struct {
	BundleId string `json:"bid"`
	Audience string `json:"aud"`
	jwt.RegisteredClaims
}

type BearerTokenAuthenticator struct {
	KeyId      string
	PrivateKey *ecdsa.PrivateKey
	// Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	Issuer string
	// Your app’s bundle ID (Ex: “com.example.testbundleid”)
	BundleId string
}

func (c *BearerTokenAuthenticator) Generate() (string, error) {
	if c.PrivateKey == nil {
		return "", errors.New("PrivateKey not given")
	}
	issuedAt := time.Now()
	expirationTime := issuedAt.Add(time.Hour)
	claims := &authenticatorClaims{
		BundleId: c.BundleId,
		Audience: appStoreConnectAudience,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: c.Issuer,
			//Audience:  jwt.ClaimStrings{appStoreConnectAudience},
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": jwt.SigningMethodES256.Alg(),
			"kid": c.KeyId,
		},
		Claims: claims,
		Method: jwt.SigningMethodES256,
	}
	sToken, err := token.SignedString(c.PrivateKey)
	if nil != err {
		return "", err
	}
	return sToken, err
}
