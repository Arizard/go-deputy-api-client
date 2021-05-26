package main

import (
	"fmt"

	"github.com/arizard/go-deputy-api-client/pkg/deputy"
)

const (
	subdomain = "brycesflorist.au"
	bearer    = "5f9e9b70122f6e6fda282977ab6c9a1d"
)

func main() {
	// Create a new Deputy API Client
	dc := deputy.NewClient(
		subdomain,
		deputy.NewBearerTokenRequestAuthoriser(bearer),
	)

	// Get current user
	var me = deputy.MeResponse{}
	if err := dc.Me(&me); err != nil {
		fmt.Println(err)
		return
	}

	// Find timesheets for current user
	queryOptions := deputy.NewDeputyQueryResourceOptions()

	queryOptions.AddSearch("employeeIsMe", "Employee", "eq", me.EmployeeId, "")
	queryOptions.AddSearch("dateFrom", "Date", "ge", "2020-04-01T00:00:00+10:00", "")
	queryOptions.AddSearch("dateTo", "Date", "le", "2020-04-30T00:00:00+10:00", "")
	queryOptions.AddSort("Date", deputy.SortAscending)

	var timesheets []deputy.TimesheetResponse
	if err := dc.QueryResource("Timesheet", queryOptions, &timesheets); err != nil {
		fmt.Println(err)
		return
	}

	// Print the results
	fmt.Printf("Timesheets for employee %d (%s)\n", me.EmployeeId, me.Name)
	for _, ts := range timesheets {
		fmt.Printf("id: %d emp: %d date: %s\n", ts.Id, ts.Employee, ts.Date)
	}
}
