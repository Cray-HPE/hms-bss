// MIT License
//
// (C) Copyright [2021,2025] Hewlett Packard Enterprise Development LP
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	base "github.com/Cray-HPE/hms-base/v2"
)

type serviceStatus struct {
	Version    string `json:"bss-version,omitempty"`
	Status     string `json:"bss-status,omitempty"`
	HSMStatus  string `json:"bss-status-hsm,omitempty"`
	EctdStatus string `json:"bss-status-etcd,omitempty"`
}

func serviceStatusAPI(w http.ResponseWriter, req *http.Request) {
	var bssStatus serviceStatus
	var httpStatus = http.StatusOK

	defer base.DrainAndCloseRequestBody(req)

	if strings.Contains(strings.ToUpper(req.URL.Path), "STATUS") ||
		strings.Contains(strings.ToUpper(req.URL.Path), "ALL") {
		bssStatus.Status = "running"
	}
	if strings.Contains(strings.ToUpper(req.URL.Path), "VERSION") ||
		strings.Contains(strings.ToUpper(req.URL.Path), "ALL") {
		dat, err := ioutil.ReadFile(".version")
		if err != nil {
			dat, err = ioutil.ReadFile("../../.version")
			if err != nil {
				httpStatus = http.StatusInternalServerError
				dat = []byte("error")
				log.Printf("Cannot read version file: %s", err)
			}
		}
		bssStatus.Version = strings.TrimSpace(string(dat))
	}
	if strings.Contains(strings.ToUpper(req.URL.Path), "HSM") ||
		strings.Contains(strings.ToUpper(req.URL.Path), "ALL") {
		bssStatus.HSMStatus = "connected"
		url := smBaseURL + "/service/values/class"
		rsp, err := smClient.Get(url)
		defer base.DrainAndCloseResponseBody(rsp)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			bssStatus.HSMStatus = "error"
			log.Printf("Cannot connect to HSM: %s", err)
		} else {
			_, err = ioutil.ReadAll(rsp.Body)
			if err != nil {
				httpStatus = http.StatusInternalServerError
				bssStatus.HSMStatus = "error"
				log.Printf("Cannot read /service/values/class response from HSM: %s", err)
			}
		}
	}
	if strings.Contains(strings.ToUpper(req.URL.Path), "ETCD") ||
		strings.Contains(strings.ToUpper(req.URL.Path), "ALL") {
		bssStatus.EctdStatus = "connected"
		randnum := rand.Intn(255)
		err := etcdTestStore(randnum)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			bssStatus.EctdStatus = "error"
			log.Printf("Test store to etcd failed: %s", err)
		} else {
			ret, err := etcdTestGet()
			if err != nil || ret != randnum {
				httpStatus = http.StatusInternalServerError
				bssStatus.EctdStatus = "error"
				if err != nil {
					log.Printf("Test read from etcd failed: %s", err)
				} else {
					log.Printf("Test read from etcd miscompare: Expected %d, Actual %d", randnum, ret)
				}
			}
		}
	}
	w.WriteHeader(httpStatus)
	out, _ := json.Marshal(bssStatus)
	fmt.Fprintln(w, string(out))
}

func etcdTestStore(testId int) error {
	data, err := json.Marshal(testId)
	err = kvstore.Store("/bss/etcdTest", string(data))
	return err
}

func etcdTestGet() (testId int, err error) {
	data, exists, err := kvstore.Get("/bss/etcdTest")
	if exists {
		err = json.Unmarshal([]byte(data), &testId)
	} else if err == nil {
		err = fmt.Errorf("Key /bss/etcdTest does not exist")
	}
	return testId, err
}
