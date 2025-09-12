package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Kranold/hyrox/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type accessToken struct {
	TokenType   string
	AccessToken string
	ExpiresAt   time.Time
}

func (cfg *StravaService) GetNewStravaAccessToken(ctx context.Context, userID uuid.UUID) (accessToken, error) {
	godotenv.Load()
	var returnToken accessToken
	/*
		- Get token object
		- check if access token is expired
		- if expired, use refresh token to get new access token
		- update the access token in the database
		- update the refresh token in the database
		- return the new access token
	*/

	// Fetching tokens
	stravaAccessTokens, err := cfg.DB.GetStravaAccessTokensByApplicationUserID(ctx, userID)
	if err != nil {
		fmt.Printf("rolf sql error %s", err)
		return accessToken{}, err
	}
	// Checking if access token is still valid, and returning if it is
	if stravaAccessTokens.ExpiresAt.After(time.Now()) {
		returnToken = accessToken{
			TokenType:   stravaAccessTokens.TokenType,
			AccessToken: stravaAccessTokens.AccessToken,
			ExpiresAt:   stravaAccessTokens.ExpiresAt,
		}
		return returnToken, nil
	}

	// Access token is expired, using refresh token to get new access token

	//Creating the required request body parameters and creating the request
	authURL := "https://www.strava.com//api/v3/oauth/token"
	data := url.Values{}
	data.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", stravaAccessTokens.RefreshToken)
	fullURL := fmt.Sprintf("%s?%s", authURL, data.Encode())

	// making the http request
	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return accessToken{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return accessToken{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return accessToken{}, err
	}

	//Parsing the response
	var tokenResponse StravaTokenResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return accessToken{}, err
	}

	//Preparing the tokenResponse for return
	returnToken = accessToken{
		TokenType:   tokenResponse.TokenType,
		AccessToken: tokenResponse.AccessToken,
		ExpiresAt:   time.Unix(int64(tokenResponse.ExpiresAt), 0),
	}

	// Updating access and refreshtokens in the database
	updateParams := database.UpdateStravaAccessTokenParams{
		StravaUserID: stravaAccessTokens.StravaUserID,
		AccessToken:  tokenResponse.AccessToken,
		TokenType:    tokenResponse.TokenType,
		ExpiresAt:    time.Unix(int64(tokenResponse.ExpiresAt), 0),
		RefreshToken: tokenResponse.RefreshToken,
	}

	err = cfg.DB.UpdateStravaAccessToken(ctx, updateParams)
	if err != nil {
		fmt.Println(err)
		return returnToken, err
	}
	// Returning the new access token
	return returnToken, nil

}
