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

func SendEmail(mail *mail.SGMailV3) error {
	logger := logging.CreateLogger()
	godotenv.Load()

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(mail)

	if err != nil {
		logger.Error(fmt.Sprintf("Error sending email http request with code %d", resp.StatusCode),
			slog.String("Error", err.Error()),
			slog.Any("mail", mail))
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		logger.Info("Email sent successfully",
			slog.Int("StatusCode", resp.StatusCode),
			slog.Any("mail", mail),
			slog.String("ResponseBody", resp.Body))
	}

	if resp.StatusCode >= 400 {
		logger.Error("Error sending email",
			slog.Int("StatusCode", resp.StatusCode),
			slog.Any("mail", mail),
			slog.String("ResponseBody", resp.Body))
		return fmt.Errorf("failed to send email, status code: %d", resp.StatusCode)
	}

	return nil
}
