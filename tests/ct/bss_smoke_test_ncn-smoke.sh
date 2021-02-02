#!/bin/bash -l
# MIT License
#
# (C) Copyright [2021] Hewlett Packard Enterprise Development LP
#
# Permission is hereby granted, free of charge, to any person obtaining a
# copy of this software and associated documentation files (the "Software"),
# to deal in the Software without restriction, including without limitation
# the rights to use, copy, modify, merge, publish, distribute, sublicense,
# and/or sell copies of the Software, and to permit persons to whom the
# Software is furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included
# in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
# THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
# OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
# ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.
###############################################################
#
#     CASM Test - Cray Inc.
#
#     TEST IDENTIFIER   : bss_smoke_test
#
#     DESCRIPTION       : Automated test for verifying basic BSS API
#                         infrastructure and installation on Cray Shasta
#                         systems.
#
#     AUTHOR            : Mitch Schooler
#
#     DATE STARTED      : 04/29/2019
#
#     LAST MODIFIED     : 09/14/2020
#
#     SYNOPSIS
#       This is a smoke test for the HMS BSS API that makes basic HTTP
#       requests using curl to verify that the service's API endpoints
#       respond and function as expected after an installation.
#
#     INPUT SPECIFICATIONS
#       Usage: bss_smoke_test
#
#       Arguments: None
#
#     OUTPUT SPECIFICATIONS
#       Plaintext is printed to stdout and/or stderr. The script exits
#       with a status of '0' on success and '1' on failure.
#
#     DESIGN DESCRIPTION
#       This smoke test is based on the Shasta health check srv_check.sh
#       script in the CrayTest repository that verifies the basic health of
#       various microservices but instead focuses exclusively on the BSS
#       API. It was implemented to run on the NCN of the system under test
#       within the DST group's Continuous Testing (CT) framework as part of
#       the ncn-smoke test suite.
#
#     SPECIAL REQUIREMENTS
#       Must be executed from the NCN.
#
#     UPDATE HISTORY
#       user       date         description
#       -------------------------------------------------------
#       schooler   04/29/2019   initial implementation
#       schooler   07/10/2019   add AuthN support for API calls
#       schooler   07/10/2019   update smoke test library location
#                               from hms-services to hms-common
#       schooler   08/19/2019   add initial check_pod_status test
#       schooler   09/06/2019   add test case documentation
#       schooler   09/09/2019   update smoke test library location
#                               from hms-common to hms-test
#       schooler   09/10/2019   update Cray copyright header
#       schooler   10/07/2019   switch from SMS to NCN naming convention
#       schooler   06/23/2020   add service version and status API tests
#       schooler   09/14/2020   use latest hms_smoke_test_lib
#
#     DEPENDENCIES
#       - hms_smoke_test_lib_ncn-resources_remote-resources.sh which is
#         expected to be packaged in /opt/cray/tests/ncn-resources/hms/hms-test
#         on the NCN.
#
#     BUGS/LIMITATIONS
#       None
#
###############################################################

# HMS test metrics test cases: 7
# 1. Check cray-bss pod statuses
# 2. GET /service/version API response code
# 3. GET /service/status API response code
# 4. GET /bootparameters API response code
# 5. GET /dumpstate API response code
# 6. GET /bootscript?nid=<nid> API response code
# 7. GET /hosts API response code

# initialize test variables
TEST_RUN_TIMESTAMP=$(date +"%Y%m%dT%H%M%S")
TEST_RUN_SEED=${RANDOM}
OUTPUT_FILES_PATH="/tmp/bss_smoke_test_out-${TEST_RUN_TIMESTAMP}.${TEST_RUN_SEED}"
SMOKE_TEST_LIB="/opt/cray/tests/ncn-resources/hms/hms-test/hms_smoke_test_lib_ncn-resources_remote-resources.sh"
TARGET="api-gw-service-nmn.local"
CURL_ARGS="-i -s -S"
MAIN_ERRORS=0
CURL_COUNT=0

# cleanup
function cleanup()
{
    echo "cleaning up..."
    rm -f ${OUTPUT_FILES_PATH}*
}

# main
function main()
{
    # retrieve Keycloak authentication token for session
    TOKEN=$(get_auth_access_token)
    TOKEN_RET=$?
    if [[ ${TOKEN_RET} -ne 0 ]] ; then
        cleanup
        exit 1
    fi
    AUTH_ARG="-H \"Authorization: Bearer $TOKEN\""

    # GET tests
    for URL_ARGS in \
        "apis/bss/boot/v1/service/version" \
        "apis/bss/boot/v1/service/status" \
        "apis/bss/boot/v1/bootparameters" \
        "apis/bss/boot/v1/dumpstate" \
        "apis/bss/boot/v1/bootscript?nid=0" \
        "apis/bss/boot/v1/hosts"
    do
        URL=$(url "${URL_ARGS}")
        URL_RET=$?
        if [[ ${URL_RET} -ne 0 ]] ; then
            cleanup
            exit 1
        fi
        run_curl "GET ${AUTH_ARG} ${URL}"
        if [[ $? -ne 0 ]] ; then
            ((MAIN_ERRORS++))
        fi
    done

    echo "MAIN_ERRORS=${MAIN_ERRORS}"
    return ${MAIN_ERRORS}
}

# check_pod_status
function check_pod_status()
{
    run_check_pod_status "cray-bss"
    return $?
}

trap ">&2 echo \"recieved kill signal, exiting with status of '1'...\" ; \
    cleanup ; \
    exit 1" SIGHUP SIGINT SIGTERM

# source HMS smoke test library file
if [[ -r ${SMOKE_TEST_LIB} ]] ; then
    . ${SMOKE_TEST_LIB}
else
    >&2 echo "ERROR: failed to source HMS smoke test library: ${SMOKE_TEST_LIB}"
    exit 1
fi

# make sure filesystem is writable for output files
touch ${OUTPUT_FILES_PATH}
if [[ $? -ne 0 ]] ; then
    >&2 echo "ERROR: output file location not writable: ${OUTPUT_FILES_PATH}"
    cleanup
    exit 1
fi

echo "Running bss_smoke_test..."

# run initial pod status test
check_pod_status
if [[ $? -ne 0 ]] ; then
    echo "FAIL: bss_smoke_test ran with failures"
    cleanup
    exit 1
fi

# run main API tests
main
if [[ $? -ne 0 ]] ; then
    echo "FAIL: bss_smoke_test ran with failures"
    cleanup
    exit 1
else
    echo "PASS: bss_smoke_test passed!"
    cleanup
    exit 0
fi
