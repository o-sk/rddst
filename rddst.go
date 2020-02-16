package rddst

import (
	"net/http"

	"golang.org/x/xerrors"
)

type Rddst interface {
	GetRedirectDestination(url string) (string, error)
}

type httpClient interface {
	Head(url string) (resp *http.Response, err error)
}

type rddst struct {
	client httpClient
}

func NewRddst(client httpClient) Rddst {
	return &rddst{
		client: client,
	}
}

func (r *rddst) GetRedirectDestination(url string) (string, error) {
	resp, err := r.client.Head(url)
	if err != nil {
		return "", err
	}
	dst := resp.Request.URL.String()
	if dst == url {
		return "", xerrors.New("The url is not redirect")
	}
	return dst, nil
}
