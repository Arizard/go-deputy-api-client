package client

import (
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

type DeputyClient struct {
	Subdomain string
	Authorise RequestAuthoriser
	Client    *http.Client
}

func NewDeputyClient(subdomain string, authorise RequestAuthoriser) *DeputyClient {
	if subdomain == "" {
		panic("subdomain cannot be empty")
	}
	return &DeputyClient{Subdomain: subdomain, Authorise: authorise, Client: &http.Client{}}
}

func (dc *DeputyClient) GetAPIUrl() string {
	return fmt.Sprintf("https://%s.deputy.com/api/v1", dc.Subdomain)
}

func (dc *DeputyClient) DoAuthorisedRequest(method string, url string, body io.Reader, deputyApiResponse DeputyAPIResponse) (err error) {
	client := dc.Client

	req, err := http.NewRequest(method, url, nil)
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
