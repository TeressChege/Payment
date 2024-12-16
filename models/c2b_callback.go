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

type Result struct {
	Ref string `json:"ref"`
}
