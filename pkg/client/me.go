package client

import "fmt"

func (dc *DeputyClient) Me(deputyApiResponse DeputyAPIResponse) error {
	url := fmt.Sprintf("%s/me", dc.GetAPIUrl())
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
