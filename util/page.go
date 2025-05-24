package util

const (
	MaxPageSize     = 100
	DefaultPageSize = 10
)

func OffsetLimit(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	} else if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	offset := (page - 1) * pageSize
	return offset, pageSize
}

func PageCount(total, pageSize int) int {
	if total <= 0 {
		return 0
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	} else if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return (total + pageSize - 1) / pageSize
}
