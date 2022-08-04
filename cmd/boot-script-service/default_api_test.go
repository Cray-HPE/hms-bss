// MIT License
//
// (C) Copyright [2022] Hewlett Packard Enterprise Development LP
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
	"regexp"
	"testing"
)

func mockGetSignedS3Url(s3Url string) (string, error) {
	return s3Url + "_signed", nil
}

func mockGetSignedS3UrlError(s3Url string) (string, error) {
	return s3Url, fmt.Errorf("error")
}

func TestReplaceS3Params_regexCompile(t *testing.T) {
	_, err := regexp.Compile(s3ParamsRegex)
	if err != nil {
		t.Errorf("Failed to compile the regex: %s, error: %v\n", s3ParamsRegex, err)
	}
}

func TestReplaceS3Params_replace(t *testing.T) {
	params := "bond=bond0 metal.server=s3://bucket/path root=craycps-s3:s3://boot-images m=s3://b/p"
	expected_params := "bond=bond0 metal.server=s3://bucket/path_signed root=craycps-s3:s3://boot-images m=s3://b/p_signed"

	newParams, err := replaceS3Params(params, mockGetSignedS3Url)
	if err != nil {
		t.Errorf("replaceS3Params returned an error for params: %s, error: %v\n", params, err)
	}
	if newParams != expected_params {
		t.Errorf("replaceS3Params failed.\n  expected: %s\n  actual: %s\n", expected_params, newParams)
	}
}

func TestReplaceS3Params_no_replace(t *testing.T) {
	// This test expects the string to remain unchanged
	// It is a limitation of replaceS3Params that it does not replace anything for the first
	// key value pair in the params string.
	// If replaceS3Params is improved to handle this remove the initial s3 path in this params
	params := "metal.server=s3://bucket/path root=craycps-s3:s3://boot-images"
	expected_params := params

	newParams, err := replaceS3Params(params, mockGetSignedS3Url)
	if err != nil {
		t.Errorf("replaceS3Params returned an error for params: %s, error: %v\n", params, err)
	}
	if newParams != expected_params {
		t.Errorf("replaceS3Params failed.\n  expected: %s\n  actual: %s\n", expected_params, newParams)
	}
}

func TestReplaceS3Params_error(t *testing.T) {
	params := "bond=bond0 metal.server=s3://bucket/path"
	expected_params := params
	newParams, err := replaceS3Params(params, mockGetSignedS3UrlError)
	if err == nil {
		t.Errorf("replaceS3Params error for params when using mock that returns an error. params: %s\n", params)
	}
	if newParams != expected_params {
		t.Errorf("replaceS3Params failed.\n  expected: %s\n  actual: %s\n", expected_params, newParams)
	}
}
