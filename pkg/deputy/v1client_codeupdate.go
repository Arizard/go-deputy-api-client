package deputy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/arizard/go-deputy-api-client/pkg/deputy/codeupdate"
)

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
