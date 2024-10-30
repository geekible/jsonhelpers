package models

type JsonResponse struct {
	IsSuccessful bool   `json:"is_successful"`
	Message      string `json:"message"`
	Data         any    `json:"data,omitempty"`
}
