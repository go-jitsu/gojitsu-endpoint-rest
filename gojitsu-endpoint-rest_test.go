package main

import (
	"testing"
	"net/http"
	"strings"
	"io/ioutil"
)

func TestDataIngest(t *testing.T) {
	go main()

	req, err := http.NewRequest("POST", "http://localhost:8080/data/foo", strings.NewReader(`{foo: "bar"}`))
	if(err != nil) {
		t.FailNow()
	}
	req.Header.Set("KEY", "key")
	resp, err := http.DefaultClient.Do(req)
	if(err != nil) {
		t.Errorf("Error executing request: " + err.Error())
		t.FailNow()
	}
	if(resp.StatusCode != 200) {
		t.Errorf("Invalid return code: %d", resp.StatusCode)
		t.FailNow()
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		t.FailNow()
	}
	if(len(respBody) != 0) {
		t.FailNow()
	}
}