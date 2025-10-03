package strava

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Kranold/hyrox/internal/logging"
	"github.com/joho/godotenv"
)

type StravaSubscription struct {
	ID            int    `json:"id"`
	ApplicationID int    `json:"application_id"`
	CallbackURL   string `json:"callback_url"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Domain        string
}

func GetSubscription() (StravaSubscription, error) {
	godotenv.Load()
	logger := logging.CreateLogger()

	requestDomain := "https://www.strava.com/api/v3/push_subscriptions"
	data := url.Values{}
	data.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))

	requestURL := fmt.Sprintf("%s?%s", requestDomain, data.Encode())

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		logger.Error("Error creating GetSubscription request",
			slog.String("Error", err.Error()))
		return StravaSubscription{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making GetSubscription request",
			slog.String("Error", err.Error()))
		return StravaSubscription{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Unexpected status code when getting subscription",
			slog.Int("StatusCode", resp.StatusCode))
		return StravaSubscription{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var subscription StravaSubscription

	decoder := json.NewDecoder(resp.Body)

	if !decoder.More() {
		logger.Info("No existing Strava subscription found")
		return StravaSubscription{}, nil
	}

	err = decoder.Decode(&subscription)
	if err != nil {
		logger.Error("Error decoding GetSubscription response",
			slog.String("Error", err.Error()))
		return StravaSubscription{}, err
	}

	// Extracting domain from CallbackURL
	endOfDomain := strings.Index(subscription.CallbackURL, "/validate_strava_subscription")
	domain := subscription.CallbackURL[:endOfDomain]
	subscription.Domain = domain

	return subscription, nil

}

func CreateSubscription() error {
	godotenv.Load()
	logger := logging.CreateLogger()

	requestDomain := "https://www.strava.com/api/v3/push_subscriptions"
	callBackURL := os.Getenv("DOMAIN") + "/validate_strava_subscription"

	data := url.Values{}
	data.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))
	data.Set("callback_url", callBackURL)
	data.Set("verify_token", "HYROX_APPLICATION_TOKEN")
	fullURL := fmt.Sprintf("%s?%s", requestDomain, data.Encode())

	request, err := http.NewRequest("POST", fullURL, nil)

	if err != nil {
		logger.Error("Error creating CreateSubscription request",
			slog.String("Error", err.Error()))
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
	}
	defer resp.Body.Close()

	return nil
}

func DeleteSubscription(subscriptionID int) error {
	godotenv.Load()
	logger := logging.CreateLogger()

	requestDomain := "https://www.strava.com/api/v3/push_subscriptions"

	data := url.Values{}
	data.Set("id", fmt.Sprintf("%d", subscriptionID))
	data.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))

	fullURL := fmt.Sprintf("%s?%s", requestDomain, data.Encode())

	request, err := http.NewRequest("DELETE", fullURL, nil)

	if err != nil {
		logger.Error("Error creating DeleteSubscription request",
			slog.String("Error", err.Error()))
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
	}
	defer resp.Body.Close()

	return nil
}
