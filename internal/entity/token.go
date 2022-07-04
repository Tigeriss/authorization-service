package entity

import "time"

type Token struct {
	Token   string    `json:"token"`
	UserID  int64     `json:"user_id"`
	Scope   string    `json:"scope"`
	Created time.Time `json:"created"`
	Expired time.Time `json:"expired"`
}
