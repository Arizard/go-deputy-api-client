package deputy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
