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
Test case for dumpstate
"""
import sys
import bsslib
from bsslib import run_bss, TESTDATA

################################################################################
#
#   hosts
#
################################################################################
def dumpstate():
    "Test dumpstate"
    test = "dumpstate"
    ret = True
    testname = "["+test+"]"
    print(testname+" dump state")
    excode, result, errstr = run_bss("dumpstate list")
    if excode != 0:
        print(testname+" FAIL: "+errstr)
        ret = False
    elif result is None:
        print(testname+" FAIL: No output received")
        ret = False
    elif not isinstance(result, dict):
        print(testname+" FAIL: Unexpected dumpstate output: %s" % result)
        ret = False
    elif "Params" not in result:
        print(testname+" FAIL: Missing \"Params\" entry: %s" % result)
        ret = False
    elif not isinstance(result["Params"], list):
        print(testname+" FAIL: \"Params\" entry not a list: %s" % result["Params"])
        ret = False
    elif "Components" not in result:
        print(testname+" FAIL: Missing \"Components\" entry: %s" % result)
        ret = False
    elif not isinstance(result["Components"], list):
        print(testname+" FAIL: \"Components\" entry not a list: %s" % result["Components"])
        ret = False
    else:
        for h in result["Components"]:
            if not isinstance(h, dict):
                print(testname+" FAIL: Component element not a map: %s" % h)
                ret = False
            for k in ["ID", "MAC", "NID", "FQDN"]:
                if k not in h:
                    print(testname+" FAIL: Component element %s missing %s element" % (h, k))
                    ret = False
                    break
            if not ret: break
        if ret:
            for h in result["Params"]:
                if not isinstance(h, dict):
                    print(testname+" FAIL: Param element not a map: %s" % h)
                    ret = False
    if ret:
        print(testname+" PASS: dumpstate")
    return ret

def test_dumpstate():
    "test dump state"
    assert dumpstate()

if __name__ == "__main__":
    sys.exit(dumpstate())
