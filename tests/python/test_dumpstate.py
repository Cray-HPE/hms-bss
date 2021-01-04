#!/usr/bin/python3
# Copyright 2019 Cray Inc. All Rights Reserved
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
