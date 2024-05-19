package dto

import "obyoy-backend/model"

// ReadReq stores order read request data
type ReadReq struct {
	UserID             string
	ParallelsentenceID string
}

// ReadReq stores order read request data
type ReadResp struct {
	SourceSentence      string   `json:"source_sentence"`
	DestinationSentence string   `json:"destination_sentence"`
	TranslatorID        string   `json:"translator_id"`
	ReviewedLines       []string `json:"reviewed_lines"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(ps *model.Parallelsentence) {
	r.SourceSentence = ps.SourceSentence
	r.DestinationSentence = ps.DestinationSentence
	r.TranslatorID = ps.TranslatorID
	r.ReviewedLines = ps.ReviewedLines
}
