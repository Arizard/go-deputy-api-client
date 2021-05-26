package deputy

func GetAllTimesheetsForEmployee(dc Client, employeeId int) (timesheets []Timesheet, err error) {
	queryOptions := NewDeputyQueryResourceOptions()
	queryOptions.AddSearch("employeeIsId", "Employee", "eq", employeeId, "")
	queryOptions.Max = 5

	for {
		var timesheetBatch []Timesheet
		if err := dc.QueryResource("Timesheet", queryOptions, &timesheetBatch); err != nil {
			return timesheets, err
		} else {
			if len(timesheetBatch) == 0 {
				break
			}
			for _, ts := range timesheetBatch {
				timesheets = append(timesheets, ts)
			}
		}
		queryOptions.Start = queryOptions.Start + queryOptions.Max
	}

	return timesheets, nil

}
