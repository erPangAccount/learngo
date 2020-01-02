package model

import (
	"crawler/engine"
)

type SearchResult struct {
	Keyword  string
	PageInfo struct {
		Total      int64
		PrevPage   int
		TargetPage int
		NextPage   int
		PageSize   int
	}
	Items []engine.Item
}
