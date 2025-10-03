package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Kranold/hyrox/internal/logging"
	"github.com/Kranold/hyrox/internal/strava"
	"github.com/joho/godotenv"
)

/* This function checks if the Strava webhooks subscription either is not created
or the callback URL is different from the one in the current environment.
If either is the case, it creates a new subscription with the correct callback URL.
If any error occurs, it will exit without deleting or creating a subscription and
log the error. This then requires manual investigation to fix. This is intended behavior
to not disrupt any existing subscriptions.
It is intended to be called manually via an API endpoint or during the deployment pipeline to production
*/

func (cfg *APIConfig) CreateStravaWebhookSubscription(w http.ResponseWriter, r *http.Request) {
	logger := logging.CreateLogger()
	godotenv.Load()

	client := &http.Client{}
	// Check existing subscription
	subscription, err := strava.GetSubscription(client)
	// if not subscription create one and return
	if err != nil {
		logger.Error("Error getting strava webhook subscription",
			slog.String("Error", err.Error()))
		http.Error(w, "Error getting strava webhook subscription", http.StatusInternalServerError)
		return
	}
	if subscription.ID == 0 {
		err = strava.CreateSubscription(client)
		if err != nil {
			logger.Error("Error creating strava webhook subscription",
				slog.String("Error", err.Error()))
			http.Error(w, "Error creating strava webhook subscription", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Created new strava webhook subscription"))

		logger.Info("Created new strava webhook subscription",
			slog.Any("Domain", os.Getenv("DOMAIN")))
		return
	}
	// if there is a subscription compare it to current domain
	if subscription.Domain == os.Getenv("DOMAIN") {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Domain already has a subscription"))
		return
	}

	// if they are different delete and create a new one
	err = strava.DeleteSubscription(client, subscription.ID)
	if err != nil {
		logger.Error("Error delete strava webhook subscription",
			slog.String("Error", err.Error()))
		http.Error(w, "Error delete strava webhook subscription", http.StatusInternalServerError)
		return
	}
	err = strava.CreateSubscription(client)
	if err != nil {
		logger.Error("Error creating strava webhook subscription",
			slog.String("Error", err.Error()))
		http.Error(w, "Error creating strava webhook subscription", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated domain for strava webhook subscription"))

	logger.Info("Created new strava webhook subscription",
		slog.String("Domain", os.Getenv("DOMAIN")),
		slog.Int("SubscriptionID", subscription.ID))

}
