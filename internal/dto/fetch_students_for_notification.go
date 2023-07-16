package dto

type FetchStudentsForNotificationRequest struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}
