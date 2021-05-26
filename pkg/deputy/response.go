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

type Timesheet struct {
	Id                  int
	Employee            int
	EmployeeHistory     int
	EmployeeAgreement   int
	Date                string
	StartTime           int
	EndTime             int
	Mealbreak           string
	MealbreakSlots      string
	TotalTime           float64
	TotalTimeInv        float64
	Cost                float64
	Roster              int
	EmployeeComment     string
	SupervisorComment   string
	Supervisor          int
	Disputed            bool
	TimeApproved        bool
	TimeApprover        int
	Discarded           bool
	ValidationFlag      int
	OperationalUnit     int
	IsInProgress        bool
	IsLeave             bool
	LeaveId             int
	LeaveRule           int
	Invoiced            bool
	InvoiceComment      string
	PayRuleApproved     bool
	Exported            bool
	StagingId           int
	PayStaged           bool
	PaycycleId          int
	File                int
	CustomFieldData     int
	RealTime            bool
	AutoProcessed       bool
	AutoRounded         bool
	AutoTimeApproved    bool
	AutoPayRuleApproved bool
	Creator             int
	Created             string
	Modified            string
}
