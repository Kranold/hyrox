package strava

import (
	"fmt"
	"os"
)

func GetStravaAuthURL() string {
	scopes := "read,read_all,activity:read,profile:read_all"
	redirectURL := "http://localhost:8080/strava_token_exchange"
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	url := fmt.Sprintf(
		"http://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=%s",
		clientID, redirectURL, scopes)
	fmt.Println(url)
	return url
}

// http://www.strava.com/oauth/authorize?client_id=174704&response_type=code&redirect_uri=http://localhost:8080/strava_token_exchange&approval_prompt=force&scope=read,read_all,activity:read,profile:read_all
