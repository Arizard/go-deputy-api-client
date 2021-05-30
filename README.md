# Deputy API Client

## Description

A Deputy API client package written in Go. Use this to make API calls to Deputy accounts. This is very much unfinished!

## Installation

You can simply run the following in your shell:

```
go get github.com/Arizard/go-deputy-api-client
```

## Getting Started

### Make a Client

The `deputy` package exposes a `Client` interface with a set of public methods which abstract the Deputy API endpoints. In this case, `V1Client` is the struct that implements `Client`. 

**V1Client**

`V1Client` accesses `v1` endpoints of the Deputy API (`/api/v1/...`) using the Go `http` package. To create a new `V1Client`, you provide the subdomain of a Deputy account (e.g. `brycesflorist.au`) and a `RequestAuthoriser` function.

**RequestAuthoriser**

`RequestAuthoriser` is a kind of "request middleware". It is a function with one argument of type `*http.Request`, and no return value.

`V1Client` expects to be able to call a `RequestAuthoriser` before executing a `*http.Request` value so that the request is authorised.

`NewBearerTokenRequestAuthoriser(token string)` creates a new `RequestAuthoriser` which adds the `Authorization: Bearer ...` header onto a `*http.Request`, therefore making it authorised.

[Find out how to generate your bearer token (permanent token) by reading the Deputy API docs.](https://www.deputy.com/api-doc/API/Authentication)

**Initialise a Working Client**

```go
var dc deputy.Client = deputy.NewV1Client(
  "brycesflorist.au",
  deputy.NewBearerTokenRequestAuthoriser("abcd...1234"),
)
```

### Make an API Call

**It Is All About Me**

Test out a simple endpoint. `GET api/v1/me` using the Client.

Methods for `Client` expect different arguments, but will always expect an `APIResponse` as the last argument. `APIResponse` should be a pointer to a value which the Client will then attempt to unmarshal the response into. **The response value's type should match what you expect to be returned**.

```go
var me = deputy.MeResponse{}
if err := dc.Me(&me); err != nil {
  fmt.Println(err)
  return
}
```

`Me` will return an error if it had trouble unmarshaling the json response, or if the endpoint returned a non-2xx status code.

## Reference

### Me

**Get current user**

```go
var me = deputy.MeResponse{}
_ := dc.Me(&me)
```

### Resource

**Get a single record**

```go
var timesheet = deputy.Timesheet
var system = "Timesheet"
var id = 100
_ := dc.GetResource(system, id, &timesheet)
```

**Create QueryOptions**

To query resources, you need a `QueryOptions`. You can add search parameters with the `AddSearch` method. You can set the sorting with `AddSort`

`AddSearch` takes 5 arguments: 

1. Name or nickname (should be unique inside the `QueryOptions`)
2. Field Name (must be a field name on the resource being queried **or** the resource joined in the case of argument 5 being not blank.)
3. Comparison (can be `eq` `ne` `ns` `is` `gt` `lt` `ge` `le` `in` `like`)
4. A static value to compare against
5. An object to join onto the resource being queried.

`AddSort` takes 2 arguments:

1. Field Name (must be a field on the resource. It cannot be a field on a joined resource.)
2. The constant `SortAscending` or `SortDescending`

```go
var employeeId = 997

queryOptions := deputy.NewQueryResourceOptions()

queryOptions.AddSearch("employeeIsMe", "Employee", "eq", employeeId, "")
queryOptions.AddSearch("dateFrom", "Date", "ge", "2020-04-01T00:00:00+10:00", "")
queryOptions.AddSearch("dateTo", "Date", "le", "2020-04-30T00:00:00+10:00", "")
queryOptions.AddSort("Date", deputy.SortAscending)
```

**Query for many records**

```go
var timesheets []deputy.Timesheet // Expect a slice of timesheet records
_ := dc.QueryResource("Timesheet", queryOptions, &timesheets) // Query the Timesheet resource using queryOptions
```

### Codeupdate

**Update a Script**

To update a script, you need to first create a `ScriptOptions`.

In the example below, `scriptId` refers to the alphanumeric script ID in Deputy.

Script types and ORM triggers are in the `codeupdate` package as constants.

```go
decafBase64 := base64.StdEncoding.EncodeToString(decafScriptCode)
scriptId := "arie_test_codeupdate"

codeupdateOptions := codeupdate.ScriptOptions{
  DecafBase64: decafBase64,
  Label:       "Arie Test Codeupdate Go Client",
  ScriptType:  codeupdate.ApplicationLaunch,
  TriggerORM:  codeupdate.TriggerTimesheet,
}

var codeWasUpdated = false // Expected response from this is a boolean.
_ := dc.CodeUpdateScript(scriptId, codeupdateOptions, &codeWasUpdated)
```



### Execute a Script via API

You can also execute scripts via API as long as they have the _Executable via SOAP/Rest API_ script type.

```go
// We need some kind of serialisable struct to pass data into the request body.
options := map[string]string{
	"func": "helloWorld",
}

// We need a response object that matches our expected result.
var scriptResponse struct {
	Message   string `json:"message"`
	Timestamp uint64 `json:"timestamp"`
}

_ := dc.ExecDeXML(scriptId, options, &scriptResponse)
```

## Example Usage

```go
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
```
