package dto

import (
	"strconv"
)

// CountResp holds the response data for counting comments
type CountResp struct {
	DatasetCount string `json:"translation_count"`
}

// FromModel converts the model data to response data
func (r *CountResp) FromModel(datasetcnt int64) {
	r.DatasetCount = strconv.FormatInt(datasetcnt, 10)
}
