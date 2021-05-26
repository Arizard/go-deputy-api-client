package deputy

type APIResponse interface{}

type MeResponse struct {
	Name         string
	FirstName    string
	LastName     string
	Company      int
	Login        string
	EmployeeId   int
	UserId       int
	InProgressTS *int
	PrimaryPhone string
	PrimaryEmail string
}
