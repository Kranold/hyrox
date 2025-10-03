package strava

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubscription(t *testing.T) {

	//Mock response data

	noSubscriptionResponse := "[]"

	existingSubscriptionResponse := map[string]interface{}{
		"id":             123456,
		"application_id": 654321,
		"callback_url":   "https://test.com/validate_strava_subscription",
		"created_at":     "2023-10-01T12:00:00Z",
		"updated_at":     "2023-10-01T12:00:00Z",
	}

	//Creating a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("client_id") == "123456" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(existingSubscriptionResponse)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(noSubscriptionResponse)
		}

	}))
	defer mockServer.Close()

	client := &http.Client{}

	// Test: Existing subscription for the client ID
	os.Setenv("STRAVA_CLIENT_ID", "123456")
	StravaAPIDomain = mockServer.URL

	subscription, err := GetSubscription(client)
	assert.NoError(t, err)
	//	assert.Equal(t, existingSubscriptionResponse, subscription)
	assert.Equal(t, "https://test.com", subscription.Domain)

	// Test: No subscription for the client ID
	os.Setenv("STRAVA_CLIENT_ID", "12")
	subscription, err = GetSubscription(client)
	assert.NoError(t, err)
	//	assert.Equal(t, noSubscriptionResponse, subscription)

}
