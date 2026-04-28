# go-apple-store-server [![GoDoc][1]][2] [![Go Report Card][3]][4] [![MIT licensed][5]][6]

---
[1]: https://godoc.org/github.com/meetleev/go-apple-store-server?status.svg
[2]: https://godoc.org/github.com/meetleev/go-apple-store-server
[3]: https://goreportcard.com/badge/github.com/meetleev/go-apple-store-server
[4]: https://goreportcard.com/report/github.com/meetleev/go-apple-store-server
[5]: https://img.shields.io/badge/license-Apache-blue.svg
[6]: LICENSE
[7]: https://github.com/meetleev/go-apple-store-server/actions/workflows/tests.yaml/badge.svg
[8]: https://github.com/meetleev/go-apple-store-server/actions/workflows/tests.yaml

The App Store Server Library

# Install

```bash
go get github.com/meetleev/go-apple-store-server
```

# Example

## 1. Create an API client

```go
package main

import (
	"fmt"

	apple_store_server "github.com/meetleev/go-apple-store-server"
	"github.com/meetleev/go-apple-store-server/types"
)

func main() {
	client, err := apple_store_server.NewAPIClientWithLocalPrivateKeyFilePath(
		"/path/to/AuthKey_ABC123XYZ.p8",
		"ABC123XYZ",
		"57246542-96fe-1a63-e053-0824d011072a",
		"com.example.app",
		types.EnvSandbox,
	)
	if err != nil {
		panic(err)
	}

	_ = client
	fmt.Println("client ready")
}
```

## 2. Fetch and decode transaction info

```go
package main

import (
	"fmt"

	apple_store_server "github.com/meetleev/go-apple-store-server"
	"github.com/meetleev/go-apple-store-server/models"
	"github.com/meetleev/go-apple-store-server/verifier"
	"github.com/meetleev/go-apple-store-server/types"
)

func main() {
	client, err := apple_store_server.NewAPIClientWithLocalPrivateKeyFilePath(
		"/path/to/AuthKey_ABC123XYZ.p8",
		"ABC123XYZ",
		"57246542-96fe-1a63-e053-0824d011072a",
		"com.example.app",
		types.EnvSandbox,
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.GetTransactionInfo("2000001161144043")
	if err != nil {
		panic(err)
	}

	payload := &models.JWSTransactionDecodedPayload{}
	sdv := verifier.NewParserWithDefault()
	if _, err := sdv.Parse(resp.SignedTransactionInfo, payload); err != nil {
		panic(err)
	}

	fmt.Printf("transactionId=%s bundleId=%s productId=%s\n", payload.TransactionId, payload.BundleId, payload.ProductId)
}
```

## 3. Verify client-uploaded `serverVerificationData`

```go
package main

import (
	"fmt"

	"github.com/meetleev/go-apple-store-server/models"
	"github.com/meetleev/go-apple-store-server/verifier"
)

func main() {
	serverVerificationData := "<client uploaded JWS>"

	payload := &models.JWSTransactionDecodedPayload{}
	sdv := verifier.NewParserWithDefault()
	if _, err := sdv.Parse(serverVerificationData, payload); err != nil {
		panic(err)
	}

	fmt.Printf("transactionId=%s bundleId=%s environment=%s\n", payload.TransactionId, payload.BundleId, payload.Environment)
}
```

## 4. Decode subscription status response

```go
package main

import (
	"fmt"

	"github.com/meetleev/go-apple-store-server/models"
	"github.com/meetleev/go-apple-store-server/verifier"
)

func main() {
	var resp models.StatusResponse
	_ = resp

	sdv := verifier.NewParserWithDefault()
	for _, group := range resp.Data {
		for _, item := range group.LastTransactions {
			tx := &models.JWSTransactionDecodedPayload{}
			if _, err := sdv.Parse(item.SignedTransactionInfo, tx); err != nil {
				panic(err)
			}
			fmt.Printf("transactionId=%s bundleId=%s\n", tx.TransactionId, tx.BundleId)
		}
	}
}
```
