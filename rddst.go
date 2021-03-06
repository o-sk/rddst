package rddst

import (
	"net/http"

	"golang.org/x/xerrors"
)

type Rddst interface {
	GetRedirectDestination(url string, strictly bool) (string, error)
}

type HttpClient interface {
	Head(url string) (resp *http.Response, err error)
}

type rddst struct {
	client HttpClient
}

func NewRddst(client HttpClient) Rddst {
	return &rddst{
		client: client,
	}
}

func (r *rddst) GetRedirectDestination(url string, strictly bool) (string, error) {
	resp, err := r.client.Head(url)
	if err != nil {
		return "", err
	}
	dst := resp.Request.URL.String()
	if strictly && dst == url {
		return "", xerrors.New("The url is not redirect")
	}
	return dst, nil
}
