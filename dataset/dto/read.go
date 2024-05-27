package dto

import "obyoy-backend/model"

// ReadReq stores order read request data
type ReadReq struct {
	UserID    string
	DatasetID string
}

// ReadReq stores order read request data
type ReadResp struct {
	DatasetID       string `json:"dataset_id"`
	Name            string `json:"name"`
	TotalLines      int    `json:"total_lines"`
	SourceLanguage  string `json:"source_language"`
	Description     string `json:"description"`
	Remarks         string `json:"remarks"`
	UploaderID      string `json:"uploader_id"`
	TranslatedLines int    `json:"translated_lines"`
	ReviewedLines   int    `json:"reviewed_lines"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(dataset *model.Dataset) {
	r.DatasetID = dataset.ID
	r.Name = dataset.Name
	r.TotalLines = int(dataset.TotalLines)
	r.ReviewedLines = int(dataset.ReviewedLines)
	r.SourceLanguage = dataset.SourceLanguage
	r.UploaderID = dataset.UploaderID
	r.TranslatedLines = int(dataset.TranslatedLines)
	r.Description = dataset.Description
	r.Remarks = dataset.Remarks
}
