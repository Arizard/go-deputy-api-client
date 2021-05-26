package deputy

type Client interface {
	Me(deputyApiResponse APIResponse) error
	GetResource(system string, id int, deputyApiResponse APIResponse) error
	QueryResource(system string, options *DeputyQueryResourceOptions, deputyApiResponse APIResponse) error
}
