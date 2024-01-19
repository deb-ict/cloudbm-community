package core

type SortOrder byte

const (
	SortNone = iota
	SortAscending
	SortDescending
)

type Sort struct {
	Fields []SortField
}

type SortField struct {
	Name  string
	Order SortOrder
}

func (s SortOrder) String() string {
	switch s {
	case SortAscending:
		return "Ascending"
	case SortDescending:
		return "Descending"
	default:
		return "None"
	}
}
