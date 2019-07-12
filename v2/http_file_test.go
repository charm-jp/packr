package packr

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func fileBox() *Box {
	box := New("file box", "../fixtures")

	return box
}

func TestFileBox_GetHTTPIndex(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(fileBox()))

	req, err := http.NewRequest("GET", "/", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal(200, res.Code)

	r.Equal("<h1>Index!</h1>", strings.TrimSpace(res.Body.String()))
}

func TestFileBox_GetHTTPRoot(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(fileBox()))

	req, err := http.NewRequest("GET", "/root_test.html", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal(200, res.Code)

	r.Equal("<h1>Root Test!</h1>", strings.TrimSpace(res.Body.String()))
}

func TestFileBox_GetHTTPDeepIndex(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(fileBox()))

	req, err := http.NewRequest("GET", "/foo/bar/deep", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal(200, res.Code)

	r.Equal("<h1>Deep Index!</h1>", strings.TrimSpace(res.Body.String()))
}

func TestFileBox_GetHTTPDeepTest(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(fileBox()))

	req, err := http.NewRequest("GET", "/foo/bar/deep/deep_test.html", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal(200, res.Code)

	r.Equal("<h1>Deep test!</h1>", strings.TrimSpace(res.Body.String()))
}
