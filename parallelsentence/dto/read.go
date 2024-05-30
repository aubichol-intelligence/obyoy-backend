package dto

import "obyoy-backend/model"

// ReadReq stores parallelsentence read request data
type ReadReq struct {
	UserID             string
	ParallelsentenceID string
}

// ReadReq stores parallelsentence read request data
type ReadResp struct {
	ID                  string   `json:"id"`
	State               string   `json:"state"`
	SourceSentence      string   `json:"source_sentence"`
	SourceLanguage      string   `json:"source_language"`
	DestinationSentence string   `json:"destination_sentence"`
	DestinationLanguage string   `json:"destination_language"`
	TranslatorID        string   `json:"translator_id"`
	ReviewedLines       []string `json:"reviewed_lines"`
	Name                string   `json:"name"`
	DatasetID           string   `json:"dataset_id"`
	DatastreamID        string   `json:"datastream_id"`
	LineNumber          int      `json:"line_number"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(ps *model.Parallelsentence) {
	r.SourceSentence = ps.SourceSentence
	r.SourceLanguage = ps.SourceLanguage
	r.DestinationSentence = ps.DestinationSentence
	r.DestinationLanguage = ps.DestinationLanguage
	r.TranslatorID = ps.TranslatorID
	r.ReviewedLines = ps.ReviewedLines
	r.Name = ps.DatasetName
	r.State = ps.State
	r.ID = ps.ID
	r.DatasetID = ps.DatasetID
	r.DatastreamID = ps.DatastreamID
	r.TranslatorID = ps.TranslatorID
	r.LineNumber = ps.LineNumber
}
