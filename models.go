package gorm

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Size  int   `json:"size" form:"size"`
	Total int64 `json:"total" form:"-"`
}

type Sorting struct {
	By  string `form:"by" json:"by"`
	Asc bool   `json:"asc" form:"asc"`
}

type Clause struct {
	Field     string    `json:"field"`
	Operator  string    `json:"operator"`
	Value     any       `json:"value"`
	Encrypted bool      `json:"encrypted"`
	Children  []*Clause `json:"children"`
}
