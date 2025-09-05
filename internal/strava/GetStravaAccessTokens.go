package strava

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func GetStravaAccessTokens(authcode string) (StravaTokenResponse, error) {
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
		fmt.Printf("Error creating request: %v\n", err)
		return StravaTokenResponse{}, err
	}
	// Setting the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	//Making the http request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return StravaTokenResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return StravaTokenResponse{}, err
	}

	//Parsing the response

	var tokenResponse StravaTokenResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return StravaTokenResponse{}, err
	}

	return tokenResponse, nil

}
