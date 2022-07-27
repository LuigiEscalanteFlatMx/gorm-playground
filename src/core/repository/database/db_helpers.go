package database

func GetOffset(page int, totalRecords int) int {
	var offset int
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * totalRecords
	}
	return offset
}
