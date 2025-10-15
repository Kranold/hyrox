package strava

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/Kranold/hyrox/internal/logging"
)

func LinkStravaAccount(client *http.Client, authcode string) (StravaTokenResponse, error) {
	logger := logging.CreateLogger()

	authURL := "https://www.strava.com/oauth/token"

	//Creating the required request body parameters and creating the request
	data := url.Values{}
	data.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))
	data.Set("code", authcode)
	data.Set("grant_type", "authorization_code")

	//Constructing the full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", authURL, data.Encode())
	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		logger.Error("Error creating request to link Strava account",
			slog.String("Error", err.Error()))
		return StravaTokenResponse{}, err
	}
	// Setting the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	//Making the http request
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making request to link Strava account",
			slog.String("Error", err.Error()))
		return StravaTokenResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Unexpected status code when linking Strava account",
			slog.Int("StatusCode", resp.StatusCode))
		return StravaTokenResponse{}, err
	}

	//Parsing the response

	var tokenResponse StravaTokenResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		logger.Error("Error decoding token response when linking Strava account",
			slog.String("Error", err.Error()))
		return StravaTokenResponse{}, err
	}

	return tokenResponse, nil

}
