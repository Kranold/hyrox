package strava

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func CreateSubscription() error {

	godotenv.Load()
	hookUrl := "https://www.strava.com/api/v3/push_subscriptions"

	data := url.Values{}
	data.Set("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))
	data.Set("callback_url", "https://hyrox-601006340303.europe-north2.run.app/validate_strava_subscription") //INSERT URL
	data.Set("verify_token", "HYROX_APPLICATION_TOKEN")
	fullURL := fmt.Sprintf("%s?%s", hookUrl, data.Encode())

	request, err := http.NewRequest("POST", fullURL, nil)

	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
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
