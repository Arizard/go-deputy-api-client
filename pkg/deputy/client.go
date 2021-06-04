package deputy

import "github.com/arizard/go-deputy-api-client/pkg/deputy/codeupdate"

type APIResponse interface{}

type Client interface {
	Me(deputyAPIResponse APIResponse) error
	ExecDeXML(scriptId string, params interface{}, deputyAPIResponse APIResponse) error
	// Resource API
	GetResource(system string, id int, deputyAPIResponse APIResponse) error
	QueryResource(system string, options *QueryResourceOptions, deputyAPIResponse APIResponse) error
	// Codeupdate
	CodeUpdateScript(scriptId string, options codeupdate.ScriptOptions, deputyAPIResponse APIResponse) error
	CodeUpdateReport(reportId string, options codeupdate.ReportOptions, deputyAPIResponse APIResponse) error
	CodeUpdateCustomApp(appId string, options codeupdate.CustomAppOptions, deputyAPIResponse APIResponse) error
	// Supervise
}
