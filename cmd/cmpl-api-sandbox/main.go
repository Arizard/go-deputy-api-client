package main

import (
	"fmt"

	"github.com/arizard/go-deputy-api-client/pkg/deputy"
)

const (
	subdomain = "brycesflorist.au"
	bearer    = "5f9e...9a1d"
)

func main() {
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
	// queryOptions.AddSearch("dateFrom", "Date", "ge", "2020-04-01T00:00:00+10:00", "")
	// queryOptions.AddSearch("dateTo", "Date", "le", "2020-04-30T00:00:00+10:00", "")
	queryOptions.AddSort("Date", deputy.SortAscending)

	var timesheets []deputy.Timesheet
	if err := dc.QueryResource("Timesheet", queryOptions, &timesheets); err != nil {
		fmt.Println(err)
		return
	}

	timesheets2, _ := deputy.GetAllTimesheetsForEmployee(dc, me.EmployeeId)

	// Print the results
	fmt.Printf("Timesheets for employee %d (%s)\n", me.EmployeeId, me.Name)
	for _, ts := range timesheets {
		fmt.Printf("id: %d emp: %d date: %s\n", ts.Id, ts.Employee, ts.Date)
	}

	fmt.Printf("length1: %d\n", len(timesheets))
	fmt.Printf("length2: %d\n", len(timesheets2))

	var scriptResponse map[string]interface{}

	err := dc.ExecDeXML("plygrnd_4mdHQcCJaXK8rAisNHMYs8", []byte{}, &scriptResponse)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", scriptResponse)
}
