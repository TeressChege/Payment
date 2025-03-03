package models

import "time"

import "encoding/json"

type WalletsResponse []WalletsResponseElement

func UnmarshalWalletsResponse(data []byte) (WalletsResponse, error) {
	var r WalletsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WalletsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WalletsResponseElement struct {
	ID              string      `json:"id"`
	Status          string      `json:"status"`
	Message         interface{} `json:"message"`
	Name            string      `json:"name"`
	Actual          float64     `json:"actual"`
	Available       float64     `json:"available"`
	DatetimeCreated time.Time   `json:"datetimeCreated"`
	LastModified    time.Time   `json:"lastModified"`
	Currency        Currency    `json:"currency"`
	Type            Type        `json:"type"`
}

type Currency struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type Type struct {
	Name string `json:"name"`
	K    string `json:"k"`
	ID   string `json:"id"`
}
