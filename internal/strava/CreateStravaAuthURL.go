package strava

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CreateStravaAuthURL() string {
	godotenv.Load()
	scopes := "read,read_all,activity:read,profile:read_all"
	redirectURL := os.Getenv("DOMAIN") + "/strava_token_exchange"
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	url := fmt.Sprintf(
		"http://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=%s",
		clientID, redirectURL, scopes)
	return url
}
