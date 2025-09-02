package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetStravaAthlete(accessToken string) (StravaAthlete, error) {
	athleteURL := "https://www.strava.com/api/v3/athlete"

	// Preparing the HTTP Request
	req, err := http.NewRequest("GET", athleteURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return StravaAthlete{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Making the HTTP Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return StravaAthlete{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return StravaAthlete{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parsing the response
	var athlete StravaAthlete
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&athlete)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return StravaAthlete{}, err
	}
	fmt.Printf("Retrieved athlete: %+v\n", athlete)

	return athlete, nil

}
