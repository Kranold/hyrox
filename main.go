package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	accessToken := "265004ae68ee4830e11a1c7e9ec01c40db7e3cdf"
	refreshToken := "8d21180b1aed2cb4fca1938743b18d888342be3a"
	GetStravaAuthURL()
	GetStravaAccessTokens("36e60ddbc996d314ad98862f429492ed5e891751")
	GetStravaAthlete(accessToken)
}
