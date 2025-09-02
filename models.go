package main

type StravaAthlete struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Premium  bool   `json:"premium"`
}

type StravaTokenResponse struct {
	TokenType    string        `json:"token_type"`
	ExpiresAt    int           `json:"expires_at"`
	ExpiresIn    int           `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
	AccessToken  string        `json:"access_token"`
	Athlete      StravaAthlete `json:"athlete"`
}
