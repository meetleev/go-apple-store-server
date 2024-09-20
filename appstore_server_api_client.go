package apple_store_server

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/meetleev/go-apple-store-server/internal"
	"github.com/meetleev/go-apple-store-server/models"
	logger "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

const (
	ProductionUrl = "https://api.storekit.itunes.apple.com"
	SandboxUrl    = "https://api.storekit-sandbox.itunes.apple.com"
)

type ErrorPayload struct {
	ErrorCode    int64  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type APIError struct {
	HttpStatusCode int    `json:"statusCode"`
	ErrorCode      int64  `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
}

func (a APIError) Error() string {
	return fmt.Sprintf("{statusCode:%d, errorCode:%d, errorMessage:%s}", a.HttpStatusCode, a.ErrorCode, a.ErrorMessage)
}

type AppStoreServerAPIClient struct {
	// Your private key ID from App Store Connect
	keyId string
	// Your private key downloaded from App Store Connect
	privateKey *ecdsa.PrivateKey
	// Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	issuer string
	// Your app’s bundle ID (Ex: “com.example.testbundleid”)
	BundleId string
	urlBase  string
}

func NewAPIClientWithLocalPrivateKeyFilePath(privateKeyFilePath, keyId, issuer, bundleId string, environment models.Environment) (*AppStoreServerAPIClient, error) {
	privateKey, err := PrivateKeyFromFile(privateKeyFilePath)
	if nil != err {
		return nil, err
	}
	p := &AppStoreServerAPIClient{privateKey: privateKey, keyId: keyId, BundleId: bundleId, issuer: issuer}
	p.SetEnv(environment)
	return p, nil
}

func NewAPIClient(privateKey *ecdsa.PrivateKey, keyId, issuer, bundleId string) *AppStoreServerAPIClient {
	p := &AppStoreServerAPIClient{privateKey: privateKey, keyId: keyId, BundleId: bundleId, issuer: issuer}
	p.SetEnv(models.EnvProduct)
	return p
}

func (c *AppStoreServerAPIClient) SetEnv(environment models.Environment) {
	if models.EnvProduct == environment {
		c.urlBase = ProductionUrl
	} else {
		c.urlBase = SandboxUrl
	}
}

func (c *AppStoreServerAPIClient) makeRequest(path, method string, options ...internal.RequestDataOption) ([]byte, error) {
	reqData := &internal.RequestData{}
	for _, option := range options {
		option(reqData)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var bodyReader io.Reader = nil
	if nil != reqData.Body {
		body, err := json.Marshal(reqData.Body)
		if err != nil {
			logger.Error("json marshal error: %v", err)
			return nil, err
		}
		bodyReader = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.urlBase, path), bodyReader)
	if err != nil {
		logger.Error("Error creating request: %v", err)
		return nil, err
	}

	token, err := c.generateBearerToken()
	if err != nil {
		logger.Error("Error jwt token: %v", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/json")
	if nil != reqData.QueryParameters {
		for k, v := range reqData.QueryParameters {
			req.URL.Query()[k] = v
		}
	}
	if nil != reqData.Body {
		req.Header.Add("Content-Type", "application/json")
	}
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalf("Error reading response body: %v", err)
		return nil, err
	}
	if http.StatusOK != resp.StatusCode {
		errPayload := &ErrorPayload{}
		if err = json.Unmarshal(body, errPayload); err != nil {
			logger.Errorf("parse err response body failed [%v]", err.Error())
			return nil, APIError{HttpStatusCode: resp.StatusCode}
		}
		return nil, APIError{HttpStatusCode: resp.StatusCode}
	}
	return body, nil
}
func (c *AppStoreServerAPIClient) generateBearerToken() (string, error) {
	bta := &internal.BearerTokenAuthenticator{BundleId: c.BundleId, PrivateKey: c.privateKey, Issuer: c.issuer, KeyId: c.keyId}
	return bta.Generate()
}

// GetTransactionInfo
// Get information about a single transaction for your app.
// @param transactionId The identifier of a transaction that belongs to the customer, and which may be an original transaction identifier.
// @return A response that contains signed transaction information for a single transaction.
// @throws APIException If a response was returned indicating the request could not be processed
// @see <a href="https://developer.apple.com/documentation/appstoreserverapi/get_transaction_info">Get Transaction Info</a>
func (c *AppStoreServerAPIClient) GetTransactionInfo(transactionId string) (*models.TransactionInfoResponse, error) {
	body, err := c.makeRequest(fmt.Sprintf("/inApps/v1/transactions/%s", transactionId), "GET")
	if nil != err {
		return nil, err
	}
	response := &models.TransactionInfoResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		logger.Errorf("parse TransactionInfo response body failed [%v]", err.Error())
		return nil, err
	}

	return response, nil
}

// GetAllSubscriptionStatuses
// Get the statuses for all of a customer’s auto-renewable subscriptions in your app.
// @param transactionId The identifier of a transaction that belongs to the customer, and which may be an original transaction identifier.
// @param status An optional filter that indicates the status of subscriptions to include in the response. Your query may specify more than one status query parameter.
// @return A response that contains status information for all of a customer’s auto-renewable subscriptions in your app.
// @throws APIException If a response was returned indicating the request could not be processed
// {@link https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses Get All Subscription Statuses}
func (c *AppStoreServerAPIClient) GetAllSubscriptionStatuses(transactionId string, status []models.Status) (*models.StatusResponse, error) {
	query := make(map[string][]string)
	var statusList []string
	for _, v := range status {
		statusList = append(statusList, string(v))
	}
	if 0 < len(statusList) {
		query["status"] = statusList
	}
	body, err := c.makeRequest(fmt.Sprintf("/inApps/v1/subscriptions/%s", transactionId), "GET", internal.WithQuery(query))
	if nil != err {
		return nil, err
	}
	response := &models.StatusResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		logger.Errorf("parse SubscriptionStatuses response body failed [%v]", err.Error())
		return nil, err
	}

	return response, nil
}
