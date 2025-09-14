package aiservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func (cfg *AIService) AICoaching(ctx context.Context, userid uuid.UUID) (string, error) {
	godotenv.Load()

	// Fetch user data and activities from Strava

	userActivities, err := cfg.DB.GetStravaActivitiesByUserID(ctx, userid)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	userData, _ := cfg.DB.GetUserByID(ctx, userid)

	contextData := map[string]interface{}{
		"userId":     userid.String(),
		"userdata":   userData,
		"activities": userActivities,
	}
	contextJSON, _ := json.MarshalIndent(contextData, "", "  ")

	responseJSON := `{
        "Suggested_next_training": "",
        "Focus_for_next_week": "",
        "Injury_preventaion": ""
    }`

	testJSON, _ := json.Marshal(responseJSON)
	// create various prompts

	contextPrompt := "Here is some context about the athelete as well as their recent workouts from stava. The data was fetched from the strava API and follows that schema. \n"
	contextPrompt += string(contextJSON)

	systemPrompt := "You are a professional fitness coach. Provide personalized advice and training suggestions based on the user's data and activities."
	systemPrompt += "Respond only in JSON with the follow structure, retunr nothing else than the json \n" + string(testJSON)

	userPrompt := "Analyze the  user data and recent activities to provide tailored fitness advice and suggest the next training session. Also suggest the focus areas for next weeks training, and give very specific advice regarding injury risk. Consider if the user could be risking injuries from current training volume and intensity and if so give specific advice how to reduce injury risk. Make your self concise and only give specific and concrete advice. "

	totalPromt := contextPrompt + "\n" + systemPrompt + "\n" + userPrompt

	// Initialize Gemini client

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(totalPromt),
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	res := result.Text()[7 : len(result.Text())-3]
	return res, nil
}
