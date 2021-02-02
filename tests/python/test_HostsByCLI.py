#!/usr/bin/python3
#  MIT License
#
#  (C) Copyright [2021] Hewlett Packard Enterprise Development LP
#
#  Permission is hereby granted, free of charge, to any person obtaining a
#  copy of this software and associated documentation files (the "Software"),
#  to deal in the Software without restriction, including without limitation
#  the rights to use, copy, modify, merge, publish, distribute, sublicense,
#  and/or sell copies of the Software, and to permit persons to whom the
#  Software is furnished to do so, subject to the following conditions:
#
#  The above copyright notice and this permission notice shall be included
#  in all copies or substantial portions of the Software.
#
#  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
#  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
#  THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
#  OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
#  ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
#  OTHER DEALINGS IN THE SOFTWARE.
"""
Test case for hosts
"""
import sys
import bsslib
from bsslib import run_bss, TESTDATA

################################################################################
#
#   hosts
#
################################################################################
def gethosts():
    "Test get hosts"
    test = "gethosts"
    ret = True
    testname = "["+test+"]"
    print(testname+" Get hosts")
    excode, result, errstr = run_bss("hosts list")
    if excode != 0:
        print(testname+" FAIL: "+errstr)
        ret = False
    elif result is not None and not isinstance(result, list):
        print(testname+" FAIL: Unexpected out retreiving hosts: %s" % result)
        ret = False
    elif result is not None:
        for h in result:
            if not isinstance(h, dict):
                print(testname+" FAIL: host element %s is not a map" % h)
                ret = False
            for k in ["ID", "MAC", "NID", "FQDN"]:
                if k not in h:
                    print(testname+" FAIL: host element %s missing %s element" % (h, k))
                    ret = False
                    break
            if not ret: break
    if ret:
        print(testname+" PASS: Get hosts")
    return ret

def test_gethosts():
    "test get hosts"
    assert gethosts()

if __name__ == "__main__":
    sys.exit(gethosts())
