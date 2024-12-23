// This file contains the code to enable pprof profiling. It is only
// included in the build when the 'pprof' build tag is set in the Dockerfile.
//
//go:build pprof

/*
 * (C) Copyright [2025] Hewlett Packard Enterprise Development LP
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func RegisterPProfHandlers() {
	// Main profiling entry point
	http.HandleFunc("/v1/debug/pprof/", pprof.Index) // Index listing all pprof endpoints

	// Specific profiling handlers
	http.HandleFunc("/v1/debug/pprof/cmdline", pprof.Cmdline) // Command-line arguments
	http.HandleFunc("/v1/debug/pprof/profile", pprof.Profile) // CPU profile (default: 30 seconds)
	http.HandleFunc("/v1/debug/pprof/symbol", pprof.Symbol)   // Symbol resolution for addresses
	http.HandleFunc("/v1/debug/pprof/trace", pprof.Trace)     // Execution trace (default: 1 second)

	// Additional profiling endpoints
	http.Handle("/v1/debug/pprof/allocs", pprof.Handler("allocs"))             // Heap allocation samples
	http.Handle("/v1/debug/pprof/block", pprof.Handler("block"))               // Goroutine blocking events
	http.Handle("/v1/debug/pprof/goroutine", pprof.Handler("goroutine"))       // Stack traces of all goroutines
	http.Handle("/v1/debug/pprof/heap", pprof.Handler("heap"))                 // Memory heap profile
	http.Handle("/v1/debug/pprof/mutex", pprof.Handler("mutex"))               // Mutex contention profile
	http.Handle("/v1/debug/pprof/threadcreate", pprof.Handler("threadcreate")) // Stack traces of thread creation
}