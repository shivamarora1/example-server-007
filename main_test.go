package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, r)

	res := w.Result()
	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)

	require.Nil(t, err, "expected nil error")
	require.JSONEq(t, string(resBytes), `{"response": "success"}`, "expected and actual response are not identical")
}

func TestHeadersHandler(t *testing.T) {

	headers := map[string]string{"Key_1": "value_1", "Key_2": "value_2"}

	r := httptest.NewRequest(http.MethodGet, "/headers", nil)
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	headersHandler(w, r)

	res := w.Result()
	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)

	require.Nil(t, err, "expected nil error")

	b, err := json.Marshal(headers)
	require.Nil(t, err, "expected nil error from Marshal")
	require.Contains(t, string(resBytes), string(b), "expected and actual response are not identical")
}
