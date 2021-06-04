package deputy

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (dc *V1Client) GetResource(system string, id int, deputyApiResponse APIResponse) error {
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

func (dc *V1Client) QueryResource(system string, options *QueryResourceOptions, deputyApiResponse APIResponse) error {
	url := fmt.Sprintf("%s/resource/%s/QUERY", dc.GetAPIUrl(), system)
	method := "POST"

	optionsPayload, err := json.Marshal(options)
	if err != nil {
		return err
	}

	err = dc.DoAuthorisedRequest(
		method,
		url,
		bytes.NewReader(optionsPayload),
		deputyApiResponse,
	)
	if err != nil {
		return err
	}

	return nil
}
