// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    statusQueryResponse, err := UnmarshalStatusQueryResponse(bytes)
//    bytes, err = statusQueryResponse.Marshal()

package models

import (
	"encoding/json"
	"time"
)

func UnmarshalStatusQueryResponse(data []byte) (StatusQueryResponse, error) {
	var r StatusQueryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StatusQueryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StatusQueryResponse struct {
	TrackingID        string            `json:"trackingId"`
	TransactionID     string            `json:"transactionId"`
	Reference         string            `json:"reference"`
	Status            string            `json:"status"`
	Message           string            `json:"message"`
	DatetimeCreated   time.Time         `json:"datetimeCreated"`
	LastUpdated       time.Time         `json:"lastUpdated"`
	DatetimeCompleted time.Time         `json:"datetimeCompleted"`
	Result            ResultStatusQuery `json:"result"`
	Request           Request           `json:"request"`
	Ipn               Ipn               `json:"ipn"`
	Results           interface{}       `json:"results"`
}

type Ipn struct {
	URL string `json:"url"`
}

type Request struct {
	AccountNumber string `json:"accountNumber"`
	Amount        string `json:"amount"`
	CommandID     string `json:"commandId"`
	MobileNumber  string `json:"mobileNumber"`
	Narration     string `json:"narration"`
}

type ResultStatusQuery struct {
	AccountNumber string `json:"accountNumber"`
	Ref           string `json:"ref"`
}
