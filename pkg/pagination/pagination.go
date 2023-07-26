package pagination

var (
	defaultPageSize  = 20
	defaultPageIndex = 1
	defaultOrder     = "id ASC"
)

func GetDefaultPagination(pageIndex, pageSize int, order string) (int, int, string) {
	newPageSize := pageSize
	newPageIndex := pageIndex
	newOrder := order
	if newPageIndex == 0 {
		newPageIndex = defaultPageIndex
	}
	if newPageSize == 0 {
		newPageSize = defaultPageSize
	}
	if newOrder == "" {
		newOrder = defaultOrder
	}
	return newPageIndex, newPageSize, newOrder
}
