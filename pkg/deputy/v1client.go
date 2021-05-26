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

// RequestAuthoriser takes a *http.Request and makes it authorised.
// For example, this could add the Authorization header.
type RequestAuthoriser func(r *http.Request)

func NewBearerTokenRequestAuthoriser(token string) RequestAuthoriser {
	return func(req *http.Request) {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

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
		return errors.New(fmt.Sprint("request returned non 2xx status code:", res.Status))
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

func (dc *V1Client) QueryResource(system string, options *DeputyQueryResourceOptions, deputyApiResponse APIResponse) error {
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