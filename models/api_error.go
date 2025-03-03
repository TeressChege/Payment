package models

import "encoding/json"

type APIErrorResponse struct {
	Status      string `json:"status"`
	Category    string `json:"category"`
	Severity    string `json:"severity"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

// UnmarshalAPIErrorResponse unmarshals a JSON byte array into an APIErrorResponse struct.
func UnmarshalAPIErrorResponse(data []byte) (APIErrorResponse, error) {
	var errResp APIErrorResponse
	err := json.Unmarshal(data, &errResp)
	if err != nil {
		return APIErrorResponse{}, err
	}
	return errResp, nil
}
