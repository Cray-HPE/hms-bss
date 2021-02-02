// MIT License
//
// (C) Copyright [2021] Hewlett Packard Enterprise Development LP
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

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
