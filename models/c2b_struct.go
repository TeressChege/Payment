package models

import "encoding/json"

type PaymentRequest struct {
	CommandId         CommandId      `json:"commandId"`
	ServiceProviderId string         `json:"serviceProviderId"`
	Reference         string         `json:"reference"`
	AccountNumber     string         `json:"accountNumber"`
	AccountNumberRef  string         `json:"accountNumberRef"`
	AccountName       string         `json:"accountName"`
	ShortCode         string         `json:"shortCode"`
	BankCode          string         `json:"bankCode"`
	CallBackUrl       string         `json:"callBackUrl"`
	Amount            string         `json:"amount"`
	Narration         string         `json:"narration"`
	Request           []PaymentField `json:"request"`
}

type PaymentField struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Label string `json:"label"`
}

type PaymentResponse struct {
	TrackingId string `json:"trackingId"`
	Reference  string `json:"reference"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// UnmarshalPaymentRequest unmarshals a JSON byte array into a PaymentRequest struct.
func UnmarshalPaymentRequest(data []byte) (PaymentRequest, error) {
	var req PaymentRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return PaymentRequest{}, err
	}
	return req, nil
}

// UnmarshalPaymentResponse unmarshals a JSON byte array into a PaymentResponse struct.
func UnmarshalPaymentResponse(data []byte) (PaymentResponse, error) {
	var resp PaymentResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return PaymentResponse{}, err
	}
	return resp, nil
}

// Marshal converts a PaymentRequest struct into JSON bytes.
func (p *PaymentRequest) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// Marshal converts a PaymentResponse struct into JSON bytes.
func (p *PaymentResponse) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
