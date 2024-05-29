package dto

import "obyoy-backend/model"

// ReadReq stores parallelsentence read request data
type ReadReq struct {
	UserID             string
	ParallelsentenceID string
}

// ReadReq stores parallelsentence read request data
type ReadResp struct {
	SourceSentence      string   `json:"source_sentence"`
	DestinationSentence string   `json:"destination_sentence"`
	TranslatorID        string   `json:"translator_id"`
	ReviewedLines       []string `json:"reviewed_lines"`
	Name                string   `json:"name"`
	DatasetID           string   `json:"dataset_id"`
	LineNumber          int      `json:"line_number"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(ps *model.Parallelsentence) {
	r.SourceSentence = ps.SourceSentence
	r.DestinationSentence = ps.DestinationSentence
	r.TranslatorID = ps.TranslatorID
	r.ReviewedLines = ps.ReviewedLines
	r.Name = ps.DatasetName
	r.DatasetID = ps.DatasetID
	r.LineNumber = ps.LineNumber
}
