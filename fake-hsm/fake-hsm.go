// MIT License
// 
// (C) Copyright [2020-2021] Hewlett Packard Enterprise Development LP
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
    "net/http"
    "log"
    "os"
	"strings"
    "encoding/json"
    "io/ioutil"
)

type ScnSubscribe struct {
    Subscriber string       `json:"Subscriber"`               //[service@]xname (nodes) or 'hmnfd'
    Components []string     `json:"Components,omitempty"`     //SCN components (usually nodes)
    Url string              `json:"Url"`                      //URL to send SCNs to
    States []string         `json:"States,omitempty"`         //Subscribe to these HW SCNs
    Enabled *bool           `json:"Enabled,omitempty"`        //true==all enable/disable SCNs
    SoftwareStatus []string `json:"SoftwareStatus,omitempty"` //Subscribe to these SW SCNs
    //Flag bool               `json:"Flags,omitempty"`        //Subscribe to flag changes
    Roles []string          `json:"Roles,omitempty"`          //Subscribe to role changes
}

type hsmComponent struct {
	ID string `json:"ID"`
	Type string `json:"Type"`
	State string `json:"State"`
	Flag string `json:"Flag"`
}

type hsmComponentList struct {
	Components []hsmComponent `json:"Components"`
}

type grpMembers struct {
	IDS []string `jtag:"ids"`
}

type hsmGroup struct {
	Label          string     `jtag:"label"`
	Description    string     `jtag:"description"`
	Tags           []string   `jtag:"tags"`
	ExclusiveGroup string     `jtag:"exclusiveGroup"`
	Members        grpMembers `jtag:"members"`
}

type hsmGroupList []hsmGroup


type httpStuff struct {
    stuff string
}


var Groups hsmGroupList
var Components hsmComponentList

func (p *httpStuff) hsmComponents(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")
	log.Printf("Sender: '%s'",agent)

	if (r.Method == "GET") {
		ba,baerr := json.Marshal(&Components)
		if (baerr != nil) {
			log.Printf("ERROR: problem marshalling component list: '%v'\n",baerr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(ba)
	} else if (r.Method == "POST") {
		//var comps hsmComponentList
		body,err := ioutil.ReadAll(r.Body)
		if (err != nil) {
			log.Printf("ERROR: problem reading request body: '%v'\n",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body,&Components)
		if (err != nil) {
			log.Printf("ERROR: problem unmarshalling request body: '%v'\n",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//copy(Components.Components,comps.Components)
		w.WriteHeader(http.StatusOK)
	} else {
        log.Printf("ERROR: request is not a GET or POST.\n")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
}

func (p *httpStuff) hsmGroups(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")
	log.Printf("Sender: '%s'",agent)

	if (r.Method == "GET") {
		ba,baerr := json.Marshal(&Groups)
		if (baerr != nil) {
			log.Printf("ERROR: problem marshalling component list: '%v'\n",baerr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(ba)
	} else if (r.Method == "POST") {
		body,err := ioutil.ReadAll(r.Body)
		if (err != nil) {
			log.Printf("ERROR: problem reading request body: '%v'\n",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body,&Groups)
		if (err != nil) {
			log.Printf("ERROR: problem unmarshalling request body: '%v'\n",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
        log.Printf("ERROR: request is not a GET or POST.\n")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
}



func (p *httpStuff) subs_rcv(w http.ResponseWriter, r *http.Request) {
    agent := r.Header.Get("User-Agent")
    log.Printf("Sender: '%s'",agent)

    if (r.Method != "POST") {
        log.Printf("ERROR: request is not a POST.\n")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var jdata ScnSubscribe
    body,err := ioutil.ReadAll(r.Body)
    err = json.Unmarshal(body,&jdata)
    if (err != nil) {
        log.Println("ERROR unmarshaling data:",err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    log.Printf("=================================================\n")
    log.Printf("Received an SCN subscription:\n")
    log.Printf("    Subscriber: %s\n",jdata.Subscriber)
    log.Printf("    Url:        %s\n",jdata.Url)
    if (len(jdata.States) > 0) {
        log.Printf("    States:     '%s'\n",jdata.States[0])
        for ix := 1; ix < len(jdata.States); ix++ {
            log.Printf("                '%s'\n",jdata.States[ix])
        }
    }
    if (len(jdata.SoftwareStatus) > 0) {
        log.Printf("    SWStatus:   '%s'\n",jdata.SoftwareStatus[0])
        for ix := 1; ix < len(jdata.SoftwareStatus); ix++ {
            log.Printf("                '%s'\n",jdata.SoftwareStatus[ix])
        }
    }
    if (len(jdata.Roles) > 0) {
        log.Printf("    Roles:      '%s'\n",jdata.Roles[0])
        for ix := 1; ix < len(jdata.Roles); ix++ {
            log.Printf("                '%s'\n",jdata.Roles[ix])
        }
    }
    if (jdata.Enabled != nil) {
        log.Printf("    Enabled:    %t\n",*jdata.Enabled)
    }
    log.Printf("\n")
    log.Printf("=================================================\n")
    w.WriteHeader(http.StatusOK)
}

func main() {
    var envstr string
    port := ":27999"

    envstr = os.Getenv("PORT")
    if (envstr != "") {
        port = envstr
		if (!strings.Contains(port,":")) {
			port = ":"+envstr
		}
    }

    urlep := "/hsm/v1/Subscriptions/SCN"
    urlcomps := "/hsm/v1/State/Components"
    urlgroups := "/hsm/v1/groups"
    hstuff := new(httpStuff)
    http.HandleFunc(urlep,hstuff.subs_rcv)
    http.HandleFunc(urlcomps,hstuff.hsmComponents)
    http.HandleFunc(urlgroups,hstuff.hsmGroups)
    log.Printf("==> Listening on endpoint '%s', port '%s'\n",urlep,port)
    log.Printf("==> Listening on endpoint '%s', port '%s'\n",urlcomps,port)
    log.Printf("==> Listening on endpoint '%s', port '%s'\n",urlgroups,port)

    err := http.ListenAndServe(port,nil)
    if (err != nil) {
        log.Println("ERROR firing up HTTP:",err)
        os.Exit(1)
    }

    os.Exit(0)
}

