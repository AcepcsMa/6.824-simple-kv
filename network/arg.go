package network

type Status string

const (
	SUCCESS = "success"
	FAIL = "fail"
)

type GetArg struct {
	Key string
}

type PutArg struct {
	Key string
	Value string
}

type GetResponse struct {
	Value string
	StatusCode Status
}

type PutResponse struct {
	StatusCode Status
}