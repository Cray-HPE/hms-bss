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
Test case for bootparameters create 2 items
"""
import sys
import bsslib
from bsslib import run_bss, TESTDATA

################################################################################
#
#   create2bootparameters
#
################################################################################
def create2bootparameters():
    "Test bootparameters create with 2 items"
    test = "create2bootparameters"
    ret = True
    testname = "["+test+"]"
    run_bss("bootparameters delete --hosts "+TESTDATA["hosts"])
    print(testname+" Create boot parameter objects")
    excode, result, errstr = run_bss(["bootparameters", "create",
                                      "--hosts", TESTDATA["hosts"],
                                      "--kernel", TESTDATA["kernel"],
                                      "--initrd", TESTDATA["initrd"],
                                      "--params", TESTDATA["params"]])
    if excode != 0:
        print(testname+" FAIL: "+errstr)
        ret = False
    elif result is not None:
        print(testname+" FAIL: Unexpected output creating new object: %s" % result)
        ret = False
    else:
        excode, result, errstr = run_bss(["bootparameters", "list",
                                          "--hosts", TESTDATA["hosts"]])
        if excode != 0:
            print(testname+" FAIL: "+errstr)
            ret = False
        elif result is None:
            print(testname+" FAIL: No output retrieving new object")
            ret = False
        elif not isinstance(result, list) or len(result) != len(TESTDATA["hosts"].split(",")):
            print(testname+" FAIL: Unexpected output: %s" % result)
            ret = False
        else:
            for obj in result:
                if not isinstance(obj, dict):
                    print(testname+" FAIL: Unexpected format for result object: " % obj)
                    ret = False
                    break
                for k in ["hosts", "kernel", "initrd", "params"]:
                    if k not in obj.keys():
                        print(testname+" FAIL: No %s entry found in result: %s" % (k, obj["hosts"]))
                        ret = False
                        beak
                if ret == 0 and (not isinstance(obj["hosts"], list) or len(obj["hosts"]) != 1 \
                                 or obj["hosts"][0] not in hosts):
                    print(testname+" FAIL: hosts entry incorrect: %s, expected one of %s"
                          % (obj["hosts"], TESTDATA["hosts"]))
                    ret = False
                ret = ret and bsslib.check(testname, obj, "kernel")
                ret = ret and bsslib.check(testname, obj, "initrd")
                ret = ret and bsslib.check(testname, obj, "params")
                if not ret:
                    break

    if ret:
        print(testname+" PASS: Create 2 items successful")
        bsslib.cleanup(["hosts", "kernel", "initrd"])
    return ret

def test_create2bootparameters():
    "test for bootparameters create 2 items"
    assert create2bootparameters()

if __name__ == "__main__":
    sys.exit(create2bootparameters())
