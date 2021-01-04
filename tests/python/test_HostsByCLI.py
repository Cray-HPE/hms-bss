#!/usr/bin/python3
# Copyright 2019 Cray Inc. All Rights Reserved
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
