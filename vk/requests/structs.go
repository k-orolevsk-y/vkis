package requests

type Requester struct {
	Token string
	Debug *bool

	apiVersion  string
	apiEndpoint string
}
