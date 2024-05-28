package dto

// ListReq stores translation read request data
type ListReq struct {
	UserID        string `json:"user_id"`
	TranslationID string `json:"translation_id"`
}
