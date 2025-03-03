package models

import (
	"encoding/json"
	"time"
)

func UnmarshalC2BCallbackResp(data []byte) (C2BCallbackResp, error) {
	var r C2BCallbackResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *C2BCallbackResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type C2BCallbackResp struct {
	TrackingID    string    `json:"trackingId"`
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Details       string    `json:"details"`
	Timestamp     time.Time `json:"timestamp"`
	Result        Result    `json:"result"`
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    b2BCallBack, err := UnmarshalB2BCallBack(bytes)
//    bytes, err = b2BCallBack.Marshal()

func UnmarshalB2BCallBack(data []byte) (B2BCallBack, error) {
	var r B2BCallBack
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *B2BCallBack) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type B2BCallBack struct {
	TrackingID    string    `json:"trackingId"`
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"` // S000000 for success
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
	Result        Result    `json:"result"`
}

type Result struct {
	Ref string `json:"ref"` // usually contains the transaction code
}
