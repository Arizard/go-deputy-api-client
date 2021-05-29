package main

import (
	"fmt"
	"os"

	"github.com/arizard/go-deputy-api-client/pkg/deputy"
)

var subdomain string
var bearer string

func main() {
	subdomain = os.Getenv("EXAMPLE_SUBDOMAIN")
	bearer = os.Getenv("EXAMPLE_BEARER")

	// Create a new Deputy API V1Client
	var dc deputy.Client = deputy.NewV1Client(
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
	queryOptions := deputy.NewQueryResourceOptions()

	queryOptions.AddSearch("employeeIsMe", "Employee", "eq", me.EmployeeId, "")
	queryOptions.AddSearch("dateFrom", "Date", "ge", "2020-04-01T00:00:00+10:00", "")
	queryOptions.AddSearch("dateTo", "Date", "le", "2020-04-30T00:00:00+10:00", "")
	queryOptions.AddSort("Date", deputy.SortAscending)

	var timesheets []deputy.Timesheet
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
