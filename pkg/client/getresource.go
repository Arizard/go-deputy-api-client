package client

import "fmt"

func (dc *DeputyClient) GetResource(system string, id int, deputyApiResponse DeputyAPIResponse) error {
	url := fmt.Sprintf("%s/resource/%s/%d", dc.GetAPIUrl(), system, id)
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
