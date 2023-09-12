package responses

type PaginateMeta struct {
	FirstPage int64		`json:"first_page"`
	LastPage int64		`json:"last_page"`
	Total int64			`json:"total"`
	From int64			`json:"from"`
	To int64			`json:"to"`
}

type PaginatedResults struct {
	Code int			`json:"code"`
	Data interface{}	`json:"data"`
	PaginateMeta		`json:"meta"`
}

