package deputy

import "fmt"

func (dc *V1Client) SuperviseGetEmployee(id int, deputyApiResponse APIResponse) error {
	url := fmt.Sprintf("%s/supervise/employee/%d", dc.GetAPIUrl(), id)
	method := "GET"

	err := dc.DoAuthorisedRequest(
		method,
		url,
		nil,
		deputyApiResponse,
	)
	if err != nil {
		return err
	}

	return nil
}

func (dc *V1Client) SuperviseAssociateEmployeeToCompany(employeeId int, companyId int, deputyApiResponse APIResponse) error {
	url := fmt.Sprintf("%s/supervise/employee/%d/assoc/%d", dc.GetAPIUrl(), employeeId, companyId)
	method := "GET"

	err := dc.DoAuthorisedRequest(
		method,
		url,
		nil,
		deputyApiResponse,
	)
	if err != nil {
		return err
	}

	return nil
}

func (dc *V1Client) SuperviseUnassociateEmployeeFromCompany(employeeId int, companyId int, deputyApiResponse APIResponse) error {
	url := fmt.Sprintf("%s/supervise/employee/%d/unassoc/%d", dc.GetAPIUrl(), employeeId, companyId)
	method := "GET"

	err := dc.DoAuthorisedRequest(
		method,
		url,
		nil,
		deputyApiResponse,
	)
	if err != nil {
		return err
	}

	return nil
}
