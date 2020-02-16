package rddst_test

import (
	"net/http"
	"testing"

	"github.com/o-sk/rddst"
	"github.com/stretchr/testify/assert"
)

func TestGetRedirectDestination(t *testing.T) {
	r := rddst.NewRddst(&http.Client{})
	u := "https://example.com"
	res, err := r.GetRedirectDestination(u)
	assert.Equal(t, u, res)
	assert.NoError(t, err)
}