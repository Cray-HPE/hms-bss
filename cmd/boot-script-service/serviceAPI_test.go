package main

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func CallServiceStatusAPI(path string, message string, passCode int) bool {
	req := httptest.NewRequest("GET", path, strings.NewReader(message))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	serviceStatusAPI(recorder, req)
	fmt.Println("Response: ", recorder.Body)
	fmt.Println("Code: ", recorder.Code)
	if recorder.Code == passCode {
		fmt.Println("PASS")
		return (true)
	}
	fmt.Println("FAIL")
	return (false)
}

const URL = "http://localhost:27778"

const SSPATH = "/boot/v1/service/status"
const SVPATH = "/boot/v1/service/version"
const SS2PATH = "/boot/v1/serviceStatus"
const SV2PATH = "/boot/v1/serviceVersion"

func TestServiceStatusAPI(t *testing.T) {
	fmt.Println("RUNNING SERVICE STATUS TESTING")
	fmt.Println("Service Status Test " + SSPATH)
	if !CallServiceStatusAPI(URL+SSPATH, "", 200) {
		t.Fail()
	}
	fmt.Println("Service Status Test " + SS2PATH)
	if !CallServiceStatusAPI(URL+SS2PATH, "", 200) {
		t.Fail()
	}
	fmt.Println("Service Status Test " + SVPATH)
	if !CallServiceStatusAPI(URL+SVPATH, "", 200) {
		t.Fail()
	}
	fmt.Println("Service Status Test " + SV2PATH)
	if !CallServiceStatusAPI(URL+SV2PATH, "", 200) {
		t.Fail()
	}
	fmt.Println("Service Status Test " + SSPATH + "/version")
	if !CallServiceStatusAPI(URL+SSPATH+"/version", "", 200) {
		t.Fail()
	}
	/* Need to start up a HSM emulator for these tests to work
	fmt.Println("Service Status Test " + SSPATH + "/hsm")
	if !CallServiceStatusAPI(URL+SSPATH+"/hsm", "", 200) {
		t.Fail()
	}
	fmt.Println("Service Status Test " + SSPATH + "/all")
	if !CallServiceStatusAPI(URL+SSPATH+"/all", "", 200) {
		t.Fail()
	}
	*/
	fmt.Println("Service Status Test " + SSPATH + "/none")
	if !CallServiceStatusAPI(URL+SSPATH+"/none", "", 200) {
		t.Fail()
	}
}
