package index

import (
	"os"

	"github.com/blevesearch/bleve"
)

type service struct {
	i bleve.Index
	b *bleve.Batch
}

//NewService create service
func NewService() (Service, error) {
	var indexPath = os.Getenv("BLEVE_PATH")
	index, err := bleve.Open(indexPath)
	switch {
	case err == bleve.ErrorIndexPathDoesNotExist:
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
		if err != nil {
			return nil, err
		}
	case err != nil:
		return nil, err
	}
	batch := index.NewBatch()
	s := &service{
		i: index,
		b: batch,
	}
	return s, nil
}

func (s *service) Index(key string, d Data) error {
	s.b.Index(key, d)
	return s.i.Batch(s.b)
}

func (s *service) Search(q string) (*SearchResult, error) {
	query := bleve.NewMatchQuery(q)
	search := bleve.NewSearchRequest(query)
	searchResults, err := s.i.Search(search)
	if err != nil {
		return nil, err
	}
	r := &SearchResult{}
	var docs []*DocumentMatch
	st := &SearchStatus{
		Total:      searchResults.Status.Total,
		Failed:     searchResults.Status.Failed,
		Successful: searchResults.Status.Successful,
	}
	r.Status = st
	r.Total = searchResults.Total
	r.MaxScore = searchResults.MaxScore
	r.Took = searchResults.Took
	for _, i := range searchResults.Hits {
		d := &DocumentMatch{}
		d.Index = i.Index
		d.Key = i.ID
		d.ID = i.ID
		d.Score = i.Score
		d.Sort = i.Sort
		d.HitNumber = i.HitNumber
		docs = append(docs, d)
	}
	r.Hits = docs
	return r, nil
}
