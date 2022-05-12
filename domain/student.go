package dto

type Student struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateStudentRequest struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
type CreateStudentResponse struct {
	ID int64 `json:"createdId"`
}

type UpdateStudentRequest struct {
	Firstname string `json:"firstname" `
	Lastname  string `json:"lastname"`
}

type UpdateBodyRequest struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
