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
/*
 * Boot Script Server
 *
 * The boot script server will collect all information required to produce an
 * iPXE boot script for each node of a system.  This script will the be
 * generated on demand and delivered to the requesting node during an iPXE
 * boot.  The main items the script will deliver are the kernel image URL/path,
 * boot arguments, and the initrd URL/path.  Note that the kernel and initrd
 * images are specified with a URL or path.  A plain path will result in a tfpt
 * download from this server.  If a URL is provided, it can be from any
 * available service which iPXE supports, and any location that the iPXE client
 * has access to. It is not restricted to a particular Cray provided service.
 */

package main

import (
	"fmt"
	base "github.com/Cray-HPE/hms-base"
	"net/http"
)

const (
	baseEndpoint     = "/boot/v1"
	notifierEndpoint = baseEndpoint + "/scn"
	// We don't use the baseEndpoint here because cloud-init doesn't like them
	metaDataRoute  = "/meta-data"
	userDataRoute  = "/user-data"
	phoneHomeRoute = "/phone-home"
)

func initHandlers() {
	http.HandleFunc(baseEndpoint+"/", Index)
	// config
	http.HandleFunc(baseEndpoint+"/bootparameters", bootParameters)
	// boot
	http.HandleFunc(baseEndpoint+"/bootscript", bootScript)
	http.HandleFunc(baseEndpoint+"/hosts", hosts)
	http.HandleFunc(baseEndpoint+"/dumpstate", dumpstate)
	http.HandleFunc(baseEndpoint+"/service/", service)
	// cloud-init
	http.HandleFunc(metaDataRoute, metaDataGet)
	http.HandleFunc(userDataRoute, userDataGet)
	http.HandleFunc(phoneHomeRoute, phoneHomePost)
	// notifications
	http.HandleFunc(notifierEndpoint, scn)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func sendAllowable(w http.ResponseWriter, allowable string) {
	w.Header().Set("allow", allowable)
	base.SendProblemDetailsGeneric(w, http.StatusMethodNotAllowed, "allow "+allowable)
}

func bootParameters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		BootparametersGet(w, r)
	case http.MethodPut:
		BootparametersPut(w, r)
	case http.MethodPost:
		BootparametersPost(w, r)
	case http.MethodPatch:
		BootparametersPatch(w, r)
	case http.MethodDelete:
		BootparametersDelete(w, r)
	default:
		sendAllowable(w, "GET,PUT,POST,PATCH,DELETE")
	}
}

func bootScript(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		BootscriptGet(w, r)
	default:
		sendAllowable(w, "GET")
	}
}

func hosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		HostsGet(w, r)
	case http.MethodPost:
		HostsPost(w, r)
	default:
		sendAllowable(w, "GET,POST")
	}
}

func dumpstate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		DumpstateGet(w, r)
	default:
		sendAllowable(w, "GET")
	}
}

func service(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		serviceStatusAPI(w, r)
	default:
		sendAllowable(w, "GET")
	}
}

func scn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		stateChangeNotification(w, r)
	default:
		sendAllowable(w, "POST")
	}
}

func metaDataGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		metaDataGetAPI(w, r)
	default:
		sendAllowable(w, "GET")
	}
}

func userDataGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userDataGetAPI(w, r)
	default:
		sendAllowable(w, "GET")
	}
}

func phoneHomePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		phoneHomePostAPI(w, r)
	default:
		sendAllowable(w, "POST")
	}
}
