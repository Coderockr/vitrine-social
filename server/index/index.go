package index

import (
	"time"
)

//Service to store application metrics
type Service interface {
	Index(key string, data Data) error
	Search(q string) (*SearchResult, error)
}

//Data to be indexed
type Data struct {
	Key  string
	ID   int64
	Data interface{}
}

//SearchResult results
type SearchResult struct {
	Status   *SearchStatus    `json:"status"`
	Hits     []*DocumentMatch `json:"hits"`
	Total    uint64           `json:"total_hits"`
	MaxScore float64          `json:"max_score"`
	Took     time.Duration    `json:"took"`
}

//SearchStatus status
type SearchStatus struct {
	Total      int `json:"total"`
	Failed     int `json:"failed"`
	Successful int `json:"successful"`
}

//DocumentMatch documents
type DocumentMatch struct {
	Index     string   `json:"index,omitempty"`
	Key       string   `json:"key"`
	ID        int64    `json:"id"`
	Score     float64  `json:"score"`
	Sort      []string `json:"sort,omitempty"`
	HitNumber uint64   `json:"-"`
}
