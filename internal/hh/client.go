package hh

type HHClient struct {
	APIKey string
}

func NewHHClient(key string) *HHClient {
	return &HHClient{
		APIKey: key,
	}
}
