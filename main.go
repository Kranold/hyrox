package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	GetStravaAuthURL()
	stravaCode := "XXX"
	GetStravaAccessTokens(stravaCode)
	GetStravaAthlete(accessToken)

	/* 
	FLOW 1: Initialize 
	Create auth URL
	Create user 
	Manually visit URL and get access code
	Paste code and get refresh token and strava user data
	Store refresh toke and strava data in db 

	FLOW 2: Regular use, implement as test afterwards
	Renew the refresh token from strava
	Get strava user data and store in db
	Update refresh token in db
	*/
}

func StravaTokenRefresh(refreshToken string) (string, error) {
	
}
