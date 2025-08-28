package ports

type HttpConnOpts struct {
	BaseURL  string
	Username string
	Password string
}

type HttpConnector interface {
	DoGet(opts HttpConnOpts, path string) ([]byte, int, error)
	DoPost(opts HttpConnOpts, path string, bodyData any) ([]byte, int, error)
}
