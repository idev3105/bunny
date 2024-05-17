package commondto

type Paging struct {
	Data  []any  `json:"data,omitempty"`
	Total uint64 `json:"total,omitempty"`
	Page  uint64 `json:"page,omitempty"`
}
