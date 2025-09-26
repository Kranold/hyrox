package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"time"

	"github.com/Kranold/hyrox/internal/database"
	"github.com/Kranold/hyrox/internal/logging"
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
	logger := logging.CreateLogger()

	var returnToken accessToken

	// Fetching tokens
	stravaAccessTokens, err := cfg.DB.GetStravaAccessTokensByApplicationUserID(ctx, userID)
	if err != nil {
		logger.Error("Error fetching strava access tokens from DB",
			slog.Any("userID", userID),
			slog.String("Error", err.Error()))
		return accessToken{}, err
	}
	// Checking if access token is still valid, and returning if it is
	if stravaAccessTokens.ExpiresAt.After(time.Now()) {
		returnToken = accessToken{
			TokenType:   stravaAccessTokens.TokenType,
			AccessToken: stravaAccessTokens.AccessToken,
			ExpiresAt:   stravaAccessTokens.ExpiresAt,
		}
		logger.Debug("Refreshed Strava access token",
			slog.String("token", returnToken.AccessToken))
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
		logger.Error("Error creating request with refresh token",
			slog.String("userID", userID.String()),
			slog.String("Error", err.Error()),
			slog.String("stacktrace", string(debug.Stack())))
		return accessToken{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error requesting new Strava access token",
			slog.String("userID", userID.String()),
			slog.String("Error", err.Error()),
			slog.String("stacktrace", string(debug.Stack())))
		return accessToken{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Error getting new strava access token",
			slog.Int("StatusCode", resp.StatusCode),
			slog.String("userID", userID.String()),
			slog.String("stacktrace", string(debug.Stack())))
		return accessToken{}, err
	}

	//Parsing the response
	var tokenResponse StravaTokenResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		logger.Error("Error decoding strava token response",
			slog.String("userID", userID.String()),
			slog.String("Error", err.Error()))
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
		logger.Error("Error updating strava access tokens in DB",
			slog.String("userID", userID.String()),
			slog.String("Error", err.Error()))
		return returnToken, err
	}
	// Returning the new access token
	logger.Debug("Refreshed Strava access token",
		slog.String("token", returnToken.AccessToken))
	return returnToken, nil

}
