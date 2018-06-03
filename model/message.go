package model

type Message struct {
	DeliveryID string `json:"id"`
	Content    string `json:"content"`
}
