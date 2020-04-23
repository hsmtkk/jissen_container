package smpwebsrv_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/jissen_container/pkg/smpwebsrv"
	"github.com/stretchr/testify/assert"
)

func TestHelloDocker(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(smpwebsrv.HelloDocker))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.Nil(t, err, "should be nil")
	got, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err, "should be nil")
	want := []byte("Hello Docker!")
	assert.Equal(t, want, got, "should be nil")
}
