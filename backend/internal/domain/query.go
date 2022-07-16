package domain

type PaginationQuery struct {
	Skip  int64 `form:"skip"`
	Limit int64 `form:"limit"`
}

type SearchQuery struct {
	Search string `form:"search"`
}

type EmployeesFiltersQuery struct {
	RegisterDateFrom string `form:"registerDateFrom"`
	RegisterDateTo   string `form:"registerDateTo"`
}

type GetEmployeesQuery struct {
	PaginationQuery
	SearchQuery
	EmployeesFiltersQuery
}

type ShiftsFiltersQuery struct {
	StartAt  string `form:"startAt"`
	EndAt    string `form:"endAt"`
	DateFrom string `form:"dateFrom"`
	DateTo   string `form:"dateTo"`
	Status   string `form:"status"`
}

type GetOrdersQuery struct {
	PaginationQuery
	SearchQuery
	ShiftsFiltersQuery
}

type WorkstationsFiltersQuery struct {
	Status string `form:"status"`
}

type GetWorkstationsQuery struct {
	PaginationQuery
	SearchQuery
	ShiftsFiltersQuery
}

func (p PaginationQuery) GetSkip() *int64 {
	if p.Skip == 0 {
		return nil
	}

	return &p.Skip
}

func (p PaginationQuery) GetLimit() *int64 {
	if p.Limit == 0 {
		return nil
	}

	return &p.Limit
}
