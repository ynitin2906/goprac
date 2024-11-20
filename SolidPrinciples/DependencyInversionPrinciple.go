package solidprinciples

import "fmt"

type MessageSender interface {
	Send(to string, message string)
}

type EmailService struct{}
type SmsService struct{}

type NotificationService struct {
	messageSender MessageSender
}

func (email *EmailService) Send(to string, message string) {
	fmt.Printf("Sending Email To %s: %s\n", to, message)
}
func (sms *SmsService) Send(to string, message string) {
	fmt.Printf("Sending Sms To %s: %s\n", to, message)
}
func NewNotificationService(sender MessageSender) *NotificationService {
	return &NotificationService{messageSender: sender}
}
func (n *NotificationService) Notify(to string, message string) {
	n.messageSender.Send(to, message)
}

func RunDependencyInversionPrinciple() {
	email := EmailService{}
	notificationServiceEmail := NewNotificationService(&email)
	notificationServiceEmail.Notify("Nitin", "Hello Email")

	sms := SmsService{}
	notificationServiceSms := NewNotificationService(&sms)
	notificationServiceSms.Notify("Nitin", "Hello SMS")

}
