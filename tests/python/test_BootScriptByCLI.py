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
Test case for bootscript
"""
import sys
import bsslib
from bsslib import run_bss, TESTDATA

################################################################################
#
#   getbootscript
#
################################################################################
def getbootscript():
    "Test bootscript list (limited)"
    test = "getbootscript"
    ret = True
    testname = "["+test+"]"
    excode, result, errstr = run_bss(["bootparameters", "create",
                                      "--hosts", TESTDATA["unknown"],
                                      "--kernel", TESTDATA["kernel"],
                                      "--initrd", TESTDATA["initrd"],
                                      "--params", TESTDATA["params"]])

    if excode == 0:
        excode, result, errstr = run_bss(["bootscript", "list",
                                          "--name", TESTDATA["host1"]])
    if excode != 0:
        #if "Not Found" not in errstr:
        print(testname+" FAIL: "+errstr)
        ret = False
    elif result is None:
        print(testname+" FAIL: No boot script produced")
        ret = False
    elif not isinstance(result, str):
        print(testname+" FAIL: Invalid output: %s" % result)
        ret = False
    else:
        rlines = result.split("\n")
        if len(rlines) < 3:
            print(testname+" FAIL: Output appears to incorrect:\n"+result)
            ret = False
        ipxe = "#!ipxe"
        if ret and rlines[0] != ipxe:
            print(testname+" FAIL: Output does not contain ipxe indicator %s:\n%s" % (ipxe, rlines[0]))
            ret = False
        else:
            # Every returned boot script has to have a chain
            # command somewhere, so we'll make sure it's there.
            foundchain = False
            for line in rlines[1:]:
                if line.startswith("chain "):
                    foundchain = True
            if not foundchain:
                print(testname+" FAIL: No chain command found in boot script:\n"+result)
                ret = False
    if ret:
        print(testname+" PASS: bootscript request successful")
        bsslib.cleanup(["unknown", "kernel", "initrd"])
    return ret

def test_getbootscript():
    "test for bootscript"
    assert getbootscript()

if __name__ == "__main__":
    sys.exit(getbootscript())
