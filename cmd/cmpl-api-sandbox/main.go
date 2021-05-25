package main

import (
	"fmt"

	deputy "github.com/arizard/go-deputy-api-client/pkg/client"
)

const (
	subdomain = "brycesflorist.au"
	bearer    = "5f9e9b70122f6e6fda282977ab6c9a1d"
)

func main() {
	dc := deputy.NewDeputyClient(
		subdomain,
		deputy.NewBearerTokenRequestAuthoriser(bearer),
	)

	// Get current user
	me := deputy.DeputyMeResponse{}
	if err := dc.Me(&me); err != nil {
		fmt.Println(err)
	}

	// Find timesheets for current user
	queryOptions := deputy.NewDeputyQueryResourceOptions()
	queryOptions.AddSearch("employee", "Employee", "eq", me.EmployeeId, "")
	queryOptions.AddSearch("dateFrom", "Date", "ge", "2020-04-01T00:00:00+10:00", "")
	queryOptions.AddSearch("dateTo", "Date", "le", "2020-04-30T00:00:00+10:00", "")

	var timesheets []deputy.DeputyTimesheet

	if err := dc.QueryResource("Timesheet", queryOptions, &timesheets); err != nil {
		fmt.Println(err)
	} else {
		for _, ts := range timesheets {
			fmt.Printf("id:%d emp:%d date:%s\n", ts.Id, ts.Employee, ts.Date)
		}
	}
}
