package Models

import "time"

type Session struct {
	ID         string    `json:"id"`
	Guesses    int       `json:"guesses"`
	IsComplete bool      `json:"iscomplete"`
	SecretWord string	 `json:"secretword"`
	Created    time.Time `json:"created"`
}
