package main

import (
	"fmt"

	deputy "github.com/arizard/go-deputy-api-client/pkg/client"
)

const (
	subdomain = "brycesflorist.au"
	bearer = "5f9e9b70122f6e6fda282977ab6c9a1d"
	timesheetId = 37968
)

func main() {
	dc := deputy.NewDeputyClient(
		subdomain,
		deputy.NewBearerTokenRequestAuthoriser(bearer),
	)

	ts := deputy.DeputyTimesheet{}

	if err := dc.GetResource("Timesheet", timesheetId, &ts); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", ts)
	}
}
