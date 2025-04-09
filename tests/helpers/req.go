package helpers

import (
	"encoding/json"
	"go-task-app/tests/globals"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func CreateReaderFromStruct(arg any) io.Reader {
	jsonData, _ := json.Marshal(arg)
	return strings.NewReader(string(jsonData))
}

func Request(req *http.Request, accessToken *string) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()

	if accessToken != nil {
		SetAuthHeader(req, *accessToken)
	}

	globals.Router.ServeHTTP(httpRecorder, req)

	return httpRecorder
}
