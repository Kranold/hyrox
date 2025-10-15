package emailservice

import (
	"github.com/Kranold/hyrox/internal/aiservice"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendCoachingEmail(userName string, email string, notes aiservice.CoachingNotes) error {

	//Email content
	from := mail.NewEmail("Rolf Carlsen", "rolfkranold@gmail.com")
	to := mail.NewEmail(userName, email)

	message := mail.NewV3Mail()

	message.SetFrom(from)
	message.SetTemplateID("d-56a8f5e04a2d491bbb4e8adc63b917ed")
	//Email content

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.SetDynamicTemplateData("user_first_name", userName)
	personalization.SetDynamicTemplateData("suggested_next_training", notes.SuggestedNextTraining)
	personalization.SetDynamicTemplateData("next_week_focus", notes.FocusForNextWeek)
	personalization.SetDynamicTemplateData("injury_prevention", notes.InjuryPrevention)

	message.AddPersonalizations(personalization)

	err := SendEmail(message)
	if err != nil {
		return err
	}

	return nil
}
