package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(w http.ResponseWriter, r *http.Request) error {
	return testingUserError("user error")
}

//
func errNotFound(w http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(w http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}

func errUnknown(w http.ResponseWriter, r *http.Request) error {
	return errors.New("unknown error")
}

func noError(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "no error")
	return nil
}

//
func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil)
		f(resp, req)
		verifyResponse(tt.code, tt.message, resp.Result(), t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(tt.code, tt.message, resp, t)
	}
}

func verifyResponse(expectCode int, expectMessage string, resp *http.Response, t *testing.T) {
	bytes, _ := ioutil.ReadAll(resp.Body)
	msg := strings.Trim(string(bytes), "\n")
	if resp.StatusCode != expectCode || msg != expectMessage {
		t.Errorf("expect (%d, %s), got (%d, %s)", expectCode, expectMessage, resp.StatusCode, msg)
	}
}
