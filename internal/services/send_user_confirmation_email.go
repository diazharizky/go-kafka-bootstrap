package services

import "fmt"

type sendUserConfirmationEmailService struct{}

func NewSendUserConfirmationEmailService() sendUserConfirmationEmailService {
	return sendUserConfirmationEmailService{}
}

func (svc sendUserConfirmationEmailService) Consume(value string) {
	fmt.Println("Sending user confirmation email", value)
}
