package email

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendCoachingEmail(userName string, email string, advice string) error {
	godotenv.Load()

	from := mail.NewEmail("Rolf Carlsen", "rolfkranold@gmail.com")
	subject := "Your Personalized Coaching Advice"
	to := mail.NewEmail(userName, email)
	htmlContent := "<strong>" + "Hello" + userName + "</strong>" + advice
	plainTextContent := advice
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		println(err.Error())
		return err
	} else {
		println(response.StatusCode)
		println(response.Body)
		println(response.Headers)
	}

	return nil
}
