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

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type userErr string

func (userErr userErr) Error() string {
	return userErr.Message()
}

func (userErr userErr) Message() string {
	return string(userErr)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return userErr("user error")
}

func errNotFounf(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown error")
}

func noErr(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       errHandlerType
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFounf, 404, "Not Found"},
	{errPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noErr, 200, "no error"},
}

func TestErrHandler(t *testing.T) {
	for _, tt := range tests {
		f := errHandler(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://test.com", nil)
		f(response, request)
		validate(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrHandlerInServer(t *testing.T) {
	for _, tt := range tests {
		f := errHandler(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		validate(response, tt.code, tt.message, t)
	}
}

func validate(response *http.Response, exportCode int, exportBody string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != exportCode || body != exportBody {
		t.Errorf("code=%d, message=%s;but result is code=%d, message=%s", exportCode, exportBody, response.StatusCode, body)
	}
}
