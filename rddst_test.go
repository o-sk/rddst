package rddst_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/o-sk/rddst"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

type mockClient struct {
	resp *http.Response
	err  error
}

func NewMockClient(resp *http.Response, err error) rddst.HttpClient {
	return &mockClient{
		resp: resp,
		err:  err,
	}
}

func (m *mockClient) Head(url string) (*http.Response, error) {
	return m.resp, m.err
}

func TestGetRedirectDestination(t *testing.T) {
	type mockParam struct {
		resp *http.Response
		err  error
	}
	type expect struct {
		resp string
		err  error
	}
	for _, tt := range []struct {
		description string
		mockParam   *mockParam
		expect      *expect
	}{
		{
			description: "The url is redirect url",
			mockParam: &mockParam{
				resp: &http.Response{
					Request: &http.Request{
						URL: func() *url.URL {
							u, _ := url.Parse("https://redirect.example.com")
							return u
						}(),
					},
				},
				err: nil,
			},
			expect: &expect{
				resp: "https://redirect.example.com",
				err:  nil,
			},
		},
		{
			description: "The url is not redirect url",
			mockParam: &mockParam{
				resp: &http.Response{
					Request: &http.Request{
						URL: func() *url.URL {
							u, _ := url.Parse("https://example.com")
							return u
						}(),
					},
				},
				err: nil,
			},
			expect: &expect{
				resp: "",
				err:  xerrors.New("The url is not redirect"),
			},
		},
		{
			description: "The http client return error",
			mockParam: &mockParam{
				resp: &http.Response{},
				err:  xerrors.New("Something has occured"),
			},
			expect: &expect{
				resp: "",
				err:  xerrors.New("Something has occured"),
			},
		},
	} {
		t.Run(tt.description, func(t *testing.T) {
			r := rddst.NewRddst(NewMockClient(tt.mockParam.resp, tt.mockParam.err))
			resp, err := r.GetRedirectDestination("https://example.com")
			assert.Equal(t, tt.expect.resp, resp)
			if tt.expect.err != nil {
				assert.Equal(t, tt.expect.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
