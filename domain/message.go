package domain

import "time"

type Message struct {
	MessageId string `json:"id"`
	SenderSourceSystem string `json:"sender_source"`
	SenderSourceSystemId string `json:"sender_source_system_id"`
	MessageBody string `json:"message_body"`
	ReceivedAt *time.Time `json:"received_at"`
	DeliveredAt *time.Time `json:"delivered_at"`
}