package rddst

import "net/http"

type Rddst interface {
	GetRedirectDestination(url string) (string, error)
}

type rddst struct {
	client *http.Client
}

func NewRddst(client *http.Client) Rddst {
	return &rddst{
		client: client,
	}
}

func (r *rddst) GetRedirectDestination(url string) (string, error) {
	return url, nil
}
