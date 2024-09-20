package apple_store_server

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// Possible errors when parsing a .p8 file.
var (
	errAuthKeyNotPem   = errors.New("token: AuthKey must be a valid .p8 PEM file")
	errAuthKeyNotECDSA = errors.New("token: AuthKey must be of type ecdsa.PrivateKey")
	//errAuthKeyNil      = errors.New("token: AuthKey was nil")
)

// PrivateKeyFromFile loads a .p8 certificate from a local file and returns a
// *ecdsa.PrivateKey.
func PrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(bytes)
}

// PrivateKeyFromBytes loads a .p8 certificate from an in memory byte array and
// returns an *ecdsa.PrivateKey.
func PrivateKeyFromBytes(bytes []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errAuthKeyNotPem
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if pk, ok := key.(*ecdsa.PrivateKey); ok {
		return pk, nil
	}
	return nil, errAuthKeyNotECDSA
}
