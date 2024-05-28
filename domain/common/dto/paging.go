package commondto

type Pagable struct {
	Total int64 `json:"total,omitempty" query:"total"`
	Page  int64 `json:"page,omitempty" query:"page"`
	Size  int64 `json:"size,omitempty" query:"size"`
	All   int64 `json:"all,omitempty" query:"all"`
}

type Page struct {
	Pagable
	Data []any `json:"data,omitempty" query:"data"`
}
