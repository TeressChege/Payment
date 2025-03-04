package tanda

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/TeressChege/Payment/helpers"
	"github.com/TeressChege/Payment/models"
)

type Tanda struct {
	Environment    Environment
	ClientId       string
	ClientSecret   string
	OrganizationId string
	ShortCode      string
	cacheToken     models.AccessTokenResponse
	Showlogs       bool
}

func NewTanda(clientId string, clientSecret string, organizationId string, shortCode string, environment Environment, showLogs bool) *Tanda {
	var accessToken = models.AccessTokenResponse{}
	return &Tanda{
		ClientId:       clientId,
		ClientSecret:   clientSecret,
		ShortCode:      shortCode,
		OrganizationId: organizationId,
		Environment:    environment,
		cacheToken:     accessToken,
		Showlogs:       showLogs,
	}
}

// setAccessToken retrieves and caches the access token.
func (t *Tanda) setAccessToken() (string, error) {
	if time.Until(t.cacheToken.ExpiresAt.UTC()).Seconds() > 0 {
		fmt.Println("Using cached token")
		return t.cacheToken.AccessToken, nil
	}
	url := "https://identity.tanda.africa/v1/oauth2/token"
	payload := []byte("grant_type=client_credentials&client_id=" + t.ClientId + "&client_secret=" + t.ClientSecret)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	res, err := helpers.NewReq(url, &payload, &headers, true)
	if err != nil {
		return "", err
	}

	if res.StatusCode() >= 200 && res.StatusCode() <= 300 {
		t.cacheToken, err = models.UnmarshalAccessTokenResponse(res.Body())
		if err != nil {
			return "", err
		}
		t.cacheToken.ExpiresAt = time.Now().Add(time.Duration(t.cacheToken.ExpiresIn) * time.Second)
	}
	return t.cacheToken.AccessToken, nil
}

// baseURL sets the appropriate environment URL.
func (t *Tanda) baseURL() string {
	if t.Environment == Production {
		return "https://api-v3.tanda.africa/io/"
	}
	return "https://api-v3-uat.tanda.africa/io"
}

func (t *Tanda) ReceivePayment(paymentRequest models.PaymentRequest) (*models.PaymentResponse, *models.APIErrorResponse) {
	token, err := t.setAccessToken()
	if err != nil {
		return nil, &models.APIErrorResponse{
			Description: err.Error(),
			Error:       err.Error(),
		}
	}
	headers := map[string]string{
		"authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}
	url := t.baseURL() + "/v3/organizations/" + t.OrganizationId + "/request"
	reqBody, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, &models.APIErrorResponse{
			Description: err.Error(),
			Error:       err.Error(),
		}
	}
	// alpha numeric only
	res, err := helpers.NewReq(url, &reqBody, &headers, t.Showlogs)
	if err != nil {
		return nil, &models.APIErrorResponse{
			Description: err.Error(),
			Error:       err.Error(),
		}
	}
	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		response, err := models.UnmarshalPaymentResponse(res.Body())
		if err != nil {
			return nil, &models.APIErrorResponse{
				Description: "unable to unmarshal payment response",
				Error:       err.Error(),
			}
		}

		return &response, nil
	} else {
		errResponse, err := models.UnmarshalAPIErrorResponse(res.Body())
		if err != nil {
			return nil, &models.APIErrorResponse{
				Description: "Unable to unmarshal api response",
				Error:       err.Error(),
			}
		}
		return nil, &errResponse
	}
}

// MerchantToCustomer handles mobile money payments to customers
func (t *Tanda) MerchantToCustomer(body models.PaymentRequest) (*models.PaymentResponse, error) {
	token, err := t.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}
	url := t.baseURL() + "/v3/organizations/" + t.OrganizationId + "/request"

	if body.CommandId == models.CustomerToMerchantMobileMoneyPayment {
		req := []models.PaymentField{
			{ID: "amount", Value: body.Amount, Label: "Amount"},
			{ID: "narration", Value: body.Narration, Label: "Narration"},
			{ID: "ipnUrl", Value: body.CallBackUrl, Label: "Notification URL"},
			{ID: "shortCode", Value: body.ShortCode, Label: "Short code"},
			{ID: "accountNumber", Value: body.AccountNumber, Label: "Phone number"},
		}
		body.Request = req
	}
	if body.CommandId == models.MerchantTo3rdPartyMerchantPayment {
		req := []models.PaymentField{
			{ID: "amount", Value: body.Amount, Label: "Amount"},
			{ID: "narration", Value: body.Narration, Label: "Narration"},
			{ID: "ipnUrl", Value: body.CallBackUrl, Label: "Notification URL"},
			{ID: "partyA", Value: body.ShortCode, Label: "Short code"},
			{ID: "partyB", Value: body.AccountNumber, Label: "Till"},
		}
		body.Request = req
	}
	if body.CommandId == models.MerchantTo3rdPartyBusinessPayment {
		req := []models.PaymentField{
			{ID: "amount", Value: body.Amount, Label: "Amount"},
			{ID: "narration", Value: body.Narration, Label: "Narration"},
			{ID: "ipnUrl", Value: body.CallBackUrl, Label: "Notification URL"},
			{ID: "shortCode", Value: body.ShortCode, Label: "Short code"},
			{ID: "businessNumber", Value: body.AccountNumber, Label: "Business number"},
			{ID: "accountReference", Value: body.AccountNumberRef, Label: "Account Reference"},
		}
		body.Request = req
	}
	if body.CommandId == models.MerchantToCustomerBankPayment {
		req := []models.PaymentField{
			{ID: "amount", Value: body.Amount, Label: "Amount"},
			{ID: "narration", Value: body.Narration, Label: "Narration"},
			{ID: "ipnUrl", Value: body.CallBackUrl, Label: "Notification URL"},
			{ID: "shortCode", Value: body.ShortCode, Label: "Short code"},
			{ID: "accountNumber", Value: body.AccountNumber, Label: "Account Number"},
			{ID: "bankCode", Value: body.BankCode, Label: "Bank code"},
			{ID: "accountName", Value: body.AccountNumber, Label: "Account Name"},
		}
		body.Request = req
	}

	if body.CommandId == models.MerchantToCustomerMobileMoneyPayment {
		req := []models.PaymentField{
			{ID: "amount", Value: body.Amount, Label: "Amount"},
			{ID: "narration", Value: body.Narration, Label: "Narration"},
			{ID: "ipnUrl", Value: body.CallBackUrl, Label: "Notification URL"},
			{ID: "shortCode", Value: body.ShortCode, Label: "Short code"},
			{ID: "accountNumber", Value: body.AccountNumber, Label: "Account Number"},
		}
		body.Request = req
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	res, err := helpers.NewReq(url, &reqBody, &headers, t.Showlogs)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errResponse, err := models.UnmarshalAPIErrorResponse(res.Body())
		if err != nil {
			return nil, errors.New(string(res.Body()))
		}
		return nil, errors.New(errResponse.Description)
	}

	response, err := models.UnmarshalPaymentResponse(res.Body())
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Struct to hold the status query request and response
type StatusQueryRequest struct {
	TrackingID string `json:"trackingId"`
	ShortCode  string `json:"shortCode"`
}

// StatusQuery checks the status of a transaction using the tracking ID and shortcode
func (t *Tanda) StatusQuery(trackingID string, shortCode string) (*models.StatusQueryResponse, *models.APIErrorResponse) {
	// Step 1: Retrieve the access token using the existing token function.
	token, err := t.setAccessToken()
	if err != nil {
		return nil, &models.APIErrorResponse{
			Description: err.Error(),
		}
	}
	// Step 2: Construct the URL using the tracking ID and shortcode provided
	url := fmt.Sprintf("%s/v3/organizations/%s/request/%s?shortCode=%s", t.baseURL(), t.OrganizationId, trackingID, shortCode)
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	res, err := helpers.NewReq(url, nil, &headers, t.Showlogs)
	if err != nil {
		return nil, &models.APIErrorResponse{
			Description: err.Error(),
		}
	}
	// Step 3: Check the status code
	if res.StatusCode() != 200 {
		errResponse, err := models.UnmarshalAPIErrorResponse(res.Body())
		if err != nil {
			return nil, &models.APIErrorResponse{
				Description: "Unable to unmarshal api response",
				Error:       err.Error(),
			}
		}
		return nil, &errResponse
	}
	// Step 4: Unmarshal the response
	var statusResp models.StatusQueryResponse
	err = json.Unmarshal(res.Body(), &statusResp)
	if err != nil {
		return nil, &models.APIErrorResponse{
			Description: "Unable to unmarshal response",
			Error:       err.Error(),
		}
	}

	return &statusResp, nil
}

func (t *Tanda) GetWalletBalances(storeUUid string) (string, *models.APIErrorResponse) {
	token, err := t.setAccessToken()
	if err != nil {
		return "", &models.APIErrorResponse{
			Description: err.Error(),
			Error:       err.Error(),
		}
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}
	url := fmt.Sprintf("https://api-v3.tanda.africa/cps/v1/head-offices/%v/stores/%v/wallets", t.OrganizationId, storeUUid)
	res, err := helpers.NewReq(url, nil, &headers, t.Showlogs)
	if err != nil {
		return "", &models.APIErrorResponse{
			Description: err.Error(),
			Error:       err.Error(),
		}
	}

	if res.StatusCode() != 200 {
		errResponse, err := models.UnmarshalAPIErrorResponse(res.Body())
		if err != nil {
			return "", &models.APIErrorResponse{
				Description: "Unable to unmarshal api response",
				Error:       err.Error(),
			}
		}
		return "", &errResponse
	}
	var walletsResponse models.WalletsResponse
	err = json.Unmarshal(res.Body(), &walletsResponse)
	if err != nil {
		return "", &models.APIErrorResponse{
			Description: "Unable to unmarshal response",
			Error:       err.Error(),
		}
	}

	return string(res.Body()), nil
}
