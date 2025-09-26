package emailservice

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Kranold/hyrox/internal/logging"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendCoachingEmail(userName string, email string, advice string) error {
	godotenv.Load()
	logger := logging.CreateLogger()

	//Email content
	from := mail.NewEmail("Rolf Carlsen", "rolfkranold@gmail.com")
	subject := "Your Personalized Coaching Advice"
	to := mail.NewEmail(userName, email)
	htmlContent := "<strong>" + "Hello" + userName + "</strong>" + advice
	plainTextContent := advice

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	//Sending email
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)

	if err != nil {
		logger.Error(fmt.Sprintf("Error sending email http request with code %d", resp.StatusCode),
			slog.String("Error", err.Error()))
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		logger.Info("Coaching email sent successfully",
			slog.Int("StatusCode", resp.StatusCode),
			slog.String("ResponseBody", resp.Body))
	}

	if resp.StatusCode >= 400 {
		logger.Error("Error sending coaching email",
			slog.Int("StatusCode", resp.StatusCode),
			slog.String("ResponseBody", resp.Body))
		return fmt.Errorf("failed to send email, status code: %d", resp.StatusCode)
	}

	return nil
}
