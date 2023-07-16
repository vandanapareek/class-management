package dto

type RegisterStudentsRequest struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}
