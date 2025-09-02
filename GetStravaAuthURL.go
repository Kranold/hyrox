package main

import (
	"fmt"
	"os"
)

func GetStravaAuthURL() string {
	scopes := "read,read_all,activity:read,profile:read_all"
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	url := fmt.Sprintf("http://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost:8080/exchange_token&approval_prompt=force&scope=%s", clientID, scopes)
	fmt.Println(url)
	return url
}
