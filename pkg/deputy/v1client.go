package deputy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/arizard/go-deputy-api-client/pkg/deputy/codeupdate"
)

type V1Client struct {
	Subdomain string
	Authorise RequestAuthoriser
	Client    *http.Client
}

func NewV1Client(subdomain string, authorise RequestAuthoriser) *V1Client {
	if subdomain == "" {
		panic("subdomain cannot be empty")
	}
	return &V1Client{Subdomain: subdomain, Authorise: authorise, Client: &http.Client{}}
}

func (dc *V1Client) GetAPIUrl() string {
	return fmt.Sprintf("https://%s.deputy.com/api/v1", dc.Subdomain)
}

func (dc *V1Client) DoAuthorisedRequest(method string, url string, body io.Reader, deputyApiResponse APIResponse) (err error) {
	client := dc.Client

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	dc.Authorise(req)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return errors.New(fmt.Sprint("request returned non 2xx status code: ", res.Status))
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, deputyApiResponse); err != nil {
		return err
	}

	return nil
}

func (dc *V1Client) Me(deputyApiResponse APIResponse) error {
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

func (dc *V1Client) ExecDeXML(scriptId string, params interface{}, deputyAPIResponse APIResponse) error {
	var body io.Reader

	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	body = bytes.NewReader(payload)

	if err := dc.DoAuthorisedRequest("POST", fmt.Sprintf("%s/execdexml/%s", dc.GetAPIUrl(), scriptId), body, deputyAPIResponse); err != nil {
		return err
	}

	return nil
}

func codeUpdate(dc *V1Client, mode string, id string, options interface{}, deputyAPIResponse APIResponse) error {
	var body io.Reader
	url := fmt.Sprintf("%s/codeupdate/%s/%s", dc.GetAPIUrl(), mode, id)

	optionsPayload, err := json.Marshal(options)
	if err != nil {
		return err
	}
	body = bytes.NewReader(optionsPayload)

	if err := dc.DoAuthorisedRequest("POST", url, body, deputyAPIResponse); err != nil && err.Error() != "unexpected end of JSON input" {
		return err
	}
	return nil
}

func (dc *V1Client) CodeUpdateScript(id string, options codeupdate.ScriptOptions, deputyAPIResponse APIResponse) error {
	return codeUpdate(dc, "dexml", id, options, deputyAPIResponse)
}

func (dc *V1Client) CodeUpdateReport(id string, options codeupdate.ReportOptions, deputyAPIResponse APIResponse) error {
	return codeUpdate(dc, "report", id, options, deputyAPIResponse)
}

func (dc *V1Client) CodeUpdateCustomApp(id string, options codeupdate.CustomAppOptions, deputyAPIResponse APIResponse) error {
	return codeUpdate(dc, "application", id, options, deputyAPIResponse)
}
