package payengine

type Timesheet struct {
	Employee int
	Bounds   TimeRange
	Breaks   []Break
}

type Break struct {
	Paid bool
	Bounds TimeRange
}

