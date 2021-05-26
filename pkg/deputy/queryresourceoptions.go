package deputy

type QueryResourceSortDirection string

const (
	SortAscending  QueryResourceSortDirection = "asc"
	SortDescending                            = "desc"
)

type QueryResourceSearchOptions map[QueryResourceSearchElementName]QueryResourceSearchElement
type QueryResourceSortOptions map[QueryResourceFieldName]QueryResourceSortDirection
type QueryResourceSearchElementName string
type QueryResourceFieldName string
type QueryResourceSearchElement struct {
	Field QueryResourceFieldName `json:"field,omitempty"`
	Type  string                 `json:"type,omitempty"`
	Data  interface{}            `json:"data,omitempty"`
	Join  string                 `json:"join,omitempty"`
}
type QueryResourceOptions struct {
	Search QueryResourceSearchOptions `json:"search,omitempty"`
	Start  int                        `json:"start,omitempty"`
	Max    int                        `json:"max,omitempty"`
	Sort   QueryResourceSortOptions   `json:"sort,omitempty"`
}

func NewQueryResourceOptions() *QueryResourceOptions {
	return &QueryResourceOptions{
		Search: QueryResourceSearchOptions{},
		Sort:   QueryResourceSortOptions{},
		Start:  0,
		Max:    500,
	}
}

func (q *QueryResourceOptions) AddSearch(label QueryResourceSearchElementName, field QueryResourceFieldName, t string, data interface{}, join string) {
	q.Search[label] = QueryResourceSearchElement{
		Field: field,
		Type:  t,
		Data:  data,
		Join:  join,
	}
}

func (q *QueryResourceOptions) AddSort(field QueryResourceFieldName, direction QueryResourceSortDirection) {
	q.Sort[field] = direction
}
