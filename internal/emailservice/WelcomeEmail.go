package emailservice

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendWelcomeEmail(userName string, email string) error {
	//Email content
	from := mail.NewEmail("Rolf Carlsen", "rolfkranold@gmail.com")
	to := mail.NewEmail(userName, email)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID("d-364443d4214049daa68bc099acbef549")

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.SetDynamicTemplateData("userName", userName)
	message.AddPersonalizations(personalization)

	err := SendEmail(message)
	if err != nil {
		return err
	}

	return nil

}
