package deputy

type APIResponse interface{}

type Client interface {
	Me(deputyAPIResponse APIResponse) error
	GetResource(system string, id int, deputyAPIResponse APIResponse) error
	QueryResource(system string, options *QueryResourceOptions, deputyAPIResponse APIResponse) error
	ExecDeXML(scriptId string, params interface{}, deputyAPIResponse APIResponse) error
}
