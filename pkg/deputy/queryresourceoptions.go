package deputy

type DeputyQueryResourceSortDirection string

const (
	SortAscending  DeputyQueryResourceSortDirection = "asc"
	SortDescending                                  = "desc"
)

type DeputyQueryResourceSearchOptions map[DeputyQueryResourceSearchElementName]DeputyQueryResourceSearchElement
type DeputyQueryResourceSortOptions map[DeputyQueryResourceFieldName]DeputyQueryResourceSortDirection
type DeputyQueryResourceSearchElementName string
type DeputyQueryResourceFieldName string
type DeputyQueryResourceSearchElement struct {
	Field DeputyQueryResourceFieldName `json:"field,omitempty"`
	Type  string                       `json:"type,omitempty"`
	Data  interface{}                  `json:"data,omitempty"`
	Join  string                       `json:"join,omitempty"`
}
type DeputyQueryResourceOptions struct {
	Search DeputyQueryResourceSearchOptions `json:"search,omitempty"`
	Start  int                              `json:"start,omitempty"`
	Max    int                              `json:"max,omitempty"`
	Sort   DeputyQueryResourceSortOptions   `json:"sort,omitempty"`
}

func NewDeputyQueryResourceOptions() *DeputyQueryResourceOptions {
	return &DeputyQueryResourceOptions{
		Search: DeputyQueryResourceSearchOptions{},
		Sort:   DeputyQueryResourceSortOptions{},
		Start:  0,
		Max:    500,
	}
}

func (q *DeputyQueryResourceOptions) AddSearch(label DeputyQueryResourceSearchElementName, field DeputyQueryResourceFieldName, t string, data interface{}, join string) {
	q.Search[label] = DeputyQueryResourceSearchElement{
		Field: field,
		Type:  t,
		Data:  data,
		Join:  join,
	}
}

func (q *DeputyQueryResourceOptions) AddSort(field DeputyQueryResourceFieldName, direction DeputyQueryResourceSortDirection) {
	q.Sort[field] = direction
}
