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

#  MIT License
#
#
#  Permission is hereby granted, free of charge, to any person obtaining a
#  copy of this software and associated documentation files (the "Software"),
#  to deal in the Software without restriction, including without limitation
#  the rights to use, copy, modify, merge, publish, distribute, sublicense,
#  and/or sell copies of the Software, and to permit persons to whom the
#  Software is furnished to do so, subject to the following conditions:
#
#
"""
TOOL IDENTIFIER : bssLib.py
TOOL TITLE      : Common Library for testing BSS
AUTHOR          : Jim Nowicki
DATE STARTED    : 04/2019
"""

from subprocess import Popen, PIPE
from json import loads
from shlex import split

TESTDATA = {
    "host1"     : "testHost1",
    "host2"     : "testHost2",
    "hosts"     : "testHost1,testHost2",
    "unknown"   : "Unknown-test_arch",
    "kernel"    : "/test/kernel",
    "newkernel" : "/test/newkernel",
    "initrd"    : "/test/initrd",
    "newinitrd" : "/test/newinitrd",
    "params"    : "testParam1 testParam2=testVal2",
    "newparams" : "newTestParam1 newTestParam2=testVal2",
}
TESTDATA_DELOPT = {
    "host1"     : "--hosts",
    "host2"     : "--hosts",
    "hosts"     : "--hosts",
    "unknown"   : "--hosts",
    "kernel"    : "--kernel",
    "newkernel" : "--kernel",
    "initrd"    : "--initrd",
    "newinitrd" : "--initrd",
    "params"    : "--params",
    "newparams" : "--params",
}

CRAYCOMMAND = 'cray'
BSSSUBCOMMAND = 'bss'

def run_command(cmd):
    "Run a command, return the exit code, stdout and stderr."
    if isinstance(cmd, str):
        cmd = split(cmd)
    process = Popen(cmd, stdout=PIPE, stderr=PIPE)
    process.wait()
    stdout, stderr = process.communicate()
    stdout = stdout.decode("utf-8")
    result = stdout.strip()

    if result != "":
        try:
            result = loads(stdout)
        except Exception:
            result = stdout
    else:
        result = None
    errstr = None
    if stderr is not None:
        errstr = stderr.decode("utf-8")
        if errstr != "":
            err = "Error: "
            if errstr.startswith(err):
                errstr = errstr[len(err):]
        else:
            errstr = None
    return process.returncode, result, errstr

def run_bss(cmd):
    "Run a cray cli bss command"
    if isinstance(cmd, str):
        cmd = split(cmd)
    cmd.insert(0, BSSSUBCOMMAND)
    cmd.insert(0, CRAYCOMMAND)
    return run_command(cmd)

def check(testname, obj, key):
    ret = True
    if isinstance(key, tuple):
        objkey = key[0]
        testkey = key[1]
    else:
        objkey = key
        testkey = key
    if obj[objkey] != TESTDATA[testkey]:
        print(testname+" FAIL: %s entry incorrect: %s, expected %s"
              % (objkey, obj[objkey], TESTDATA[testkey]))
        ret = False
    return ret
        

def cleanup(items = None):
    "Try to clean up the items from the BSS server that we created."
    if items is None:
        items = TESTDATA.keys()
    for i in items:
        run_bss(["bootparameters", "delete", TESTDATA_DELOPT[i], TESTDATA[i]])
