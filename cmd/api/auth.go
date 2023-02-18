package main

import "time"

type Auth struct {
	Issuer string
	Audience string
	Secret string
	TokenExpiry time.Duration
	RefreshExpiry time.Duration
	CookieDomain string
	CookiePath string
	CookieName string
}


type jwtUser struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}