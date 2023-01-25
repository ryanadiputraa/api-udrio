package pagination

type Page struct {
	CurrentPage  int `json:"current_page"`
	TotalPage    int `json:"total_page"`
	TotalData    int `json:"total_data"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
}

type Pagination struct {
	Size      int
	Page      int
	TotalPage int
	TotalData int
}

func NewPagination(size, page, totalData int) *Pagination {
	totalPage := 0
	if totalData >= 0 {
		totalPage = (totalData + size - 1) / size
	}
	return &Pagination{
		Size:      size,
		Page:      page,
		TotalPage: totalPage,
		TotalData: totalData,
	}
}

func Offset(size, page int) int {
	return (page - 1) * size
}

func (p *Pagination) NextPage() int {
	if p.TotalPage >= 0 && p.Page >= p.TotalPage {
		return p.TotalPage
	}
	return p.Page + 1
}

func (p *Pagination) PreviousPage() int {
	if p.Page > 1 {
		return p.Page - 1
	}
	return p.Page
}
