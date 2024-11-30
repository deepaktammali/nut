package postgres

func ParseOrderBy(value string) string {
	if value == "desc" {
		return "DESC"
	} else {
		return "ASC"
	}
}
