// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type GptUser struct {
	ID        int64     `json:"id"`
	ChatID    string    `json:"chat_id"`
	GptToken  string    `json:"gpt_token"`
	CreatedAt time.Time `json:"created_at"`
}