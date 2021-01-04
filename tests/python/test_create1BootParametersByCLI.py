#!/usr/bin/python3
# Copyright 2019 Cray Inc. All Rights Reserved
"""
Test case for bootparameters create 1 item
"""
import sys
import bsslib
from bsslib import run_bss, TESTDATA

################################################################################
#
#   create1bootparameters
#
################################################################################
def create1bootparameters():
    "Test bootparameters create with 1 item"
    test = "create1bootparameters"
    ret = True
    testname = "["+test+"]"
    run_bss("bootparameters delete --hosts "+TESTDATA["host1"])
    print(testname+" Create boot parameter objects")
    excode, result, errstr = run_bss(["bootparameters", "create",
                                      "--hosts", TESTDATA["host1"],
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
                                          "--hosts", TESTDATA["host1"]])
        if excode != 0:
            print(testname+" FAIL: "+errstr)
            ret = False
        elif result is None:
            print(testname+" FAIL: No output retrieving new object")
            ret = False
        elif not isinstance(result, list) or len(result) != 1 or not isinstance(result[0], dict):
            print(testname+" FAIL: Unexpected output: %s", result)
            ret = False
        else:
            obj = result[0]
            for k in ["hosts", "kernel", "initrd", "params"]:
                if k not in obj.keys():
                    print(testname+" FAIL: No %s entry found in result: %s" % (k, obj["hosts"]))
                    ret = False
                    break
        if ret and (not isinstance(obj["hosts"], list) or len(obj["hosts"]) != 1 \
                         or obj["hosts"][0] != TESTDATA["host1"]):
            print(testname+" FAIL: hosts entry incorrect: %s, expected %s"
                  % (obj["hosts"], TESTDATA["host1"]))
            ret = False
        ret = ret and bsslib.check(testname, obj, "kernel")
        ret = ret and bsslib.check(testname, obj, "initrd")
        ret = ret and bsslib.check(testname, obj, "params")

    if ret:
        print(testname+" PASS: Create 1 item successful")
        bsslib.cleanup(["host1", "kernel", "initrd"])
    return ret

def test_create1bootparameters():
    "test for bootparameters create 1 item"
    assert create1bootparameters()

if __name__ == "__main__":
    sys.exit(create1bootparameters())
