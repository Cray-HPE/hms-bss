#!/bin/bash -l
#
# Copyright 2020 Hewlett Packard Enterprise Development LP
#
###############################################################
#
#     CASM Test - Cray Inc.
#
#     TEST IDENTIFIER   : bss_tavern_api_test
#
#     DESCRIPTION       : Automated test for verifying the HMS Boot 
#                         Script Service (BSS) API on Cray Shasta systems.
#                         
#     AUTHOR            : Mitch Schooler
#
#     DATE STARTED      : 01/21/2020
#
#     LAST MODIFIED     : 09/14/2020
#
#     SYNOPSIS
#       This is a test wrapper for HMS Boot Script Service (BSS) API
#       tests implemented in Tavern that launch via pytest within the
#       Continuous Test (CT) framework. All Tavern tests packaged in
#       the target CT test directory for BSS are executed.
#
#     INPUT SPECIFICATIONS
#       Usage: bss_tavern_api_test
#       
#       Arguments: None
#
#     OUTPUT SPECIFICATIONS
#       Plaintext is printed to stdout and/or stderr. The script exits
#       with a status of '0' on success and '1' on failure.
#
#     DESIGN DESCRIPTION
#       This test wrapper generates a Tavern configuration file based
#       on the target test system it is running on and then executes all
#       BSS Tavern API CT tests using HMS's hms-pytest container which
#       includes pytest and other dependencies required to run Tavern.
#
#     SPECIAL REQUIREMENTS
#       Must be executed on the target test system on a fully-installed
#       NCN with the Continuous Test infrastructure in place.
#
#     UPDATE HISTORY
#       user       date         description
#       -------------------------------------------------------
#       schooler   01/21/2020   initial implementation
#       schooler   09/14/2020   use latest hms_common_file_generator
#
#     DEPENDENCIES
#       - hms-pytest wrapper script which is expected to be packaged
#         in /usr/bin on the NCN.
#       - hms_smoke_test_lib_ncn-resources_remote-resources.sh which
#         is expected to be packaged in
#         /opt/cray/tests/ncn-resources/hms/hms-test on the NCN.
#       - hms_pytest_ini_file_generator_ncn-resources_remote-resources.py
#         which is expected to be packaged in
#         /opt/cray/tests/ncn-resources/hms/hms-test on the NCN.
#       - hms_common_file_generator_ncn-resources_remote-resources.py
#         which is expected to be packaged in
#         /opt/cray/tests/ncn-resources/hms/hms-test on the NCN.
#       - BSS Tavern API tests with names of the form test_*.tavern.yaml
#         which are expected to be packaged in
#         /opt/cray/tests/ncn-functional/hms/hms-bss on the NCN.
#
#     BUGS/LIMITATIONS
#       None
#
###############################################################

# cleanup
function cleanup()
{
    echo "cleaning up..."
    if [[ -d ${BSS_TEST_DIR}/.pytest_cache ]] ; then
        rm -rf ${BSS_TEST_DIR}/.pytest_cache
    fi
    rm -f ${PYTEST_INI_PATH}
    rm -f ${COMMON_FILE_PATH}
}

# HMS path declarations
HMS_TEST_LIB="/opt/cray/tests/ncn-resources/hms/hms-test/hms_smoke_test_lib_ncn-resources_remote-resources.sh"
PYTEST_INI_GENERATOR="/opt/cray/tests/ncn-resources/hms/hms-test/hms_pytest_ini_file_generator_ncn-resources_remote-resources.py"
PYTEST_INI_PATH="/opt/cray/tests/ncn-functional/hms/hms-bss/pytest.ini"
COMMON_FILE_GENERATOR="/opt/cray/tests/ncn-resources/hms/hms-test/hms_common_file_generator_ncn-resources_remote-resources.py"
COMMON_FILE_PATH="/opt/cray/tests/ncn-functional/hms/hms-bss/common.yaml"
BSS_TEST_DIR="/opt/cray/tests/ncn-functional/hms/hms-bss"
API_TARGET="https://api-gw-service-nmn.local/apis"

# set SSL certificate checking to True for test execution from the NCN
VERIFY="True"
echo "VERIFY=${VERIFY}"

# set up signal handling
trap ">&2 echo \"recieved kill signal, exiting with status of '1'...\" ; \
    cleanup ; \
    exit 1" SIGHUP SIGINT SIGTERM

# verify that the hms-pytest wrapper script exists
PYTEST_PATH=$(which hms-pytest)
if [[ -z ${PYTEST_PATH} ]] ; then
    >&2 echo "ERROR: failed to locate command: hms-pytest"
    cleanup
    exit 1
fi

# source the HMS smoke test library file
if [[ -r ${HMS_TEST_LIB} ]] ; then
    . ${HMS_TEST_LIB}
else
    >&2 echo "ERROR: failed to source HMS smoke test library: ${HMS_TEST_LIB}"
    cleanup
    exit 1
fi

# verify that the Tavern test directory exists
if [[ ! -d ${BSS_TEST_DIR} ]] ; then
    >&2 echo "ERROR: failed to locate Tavern test directory: ${BSS_TEST_DIR}"
    cleanup
    exit 1
else
    TEST_DIR_FILES=$(ls ${BSS_TEST_DIR})
    TEST_DIR_TAVERN_FILES=$(echo "${TEST_DIR_FILES}" | grep -E "^test_.*\.tavern\.yaml")
    if [[ -z "${TEST_DIR_TAVERN_FILES}" ]] ; then
        >&2 echo "ERROR: no Tavern tests in CT test directory: ${BSS_TEST_DIR}"
        >&2 echo "${TEST_DIR_FILES}"
        cleanup
        exit 1
    fi
fi

# verify that the pytest.ini generator tool exists
if [[ ! -x ${PYTEST_INI_GENERATOR} ]] ; then
    >&2 echo "ERROR: failed to locate executable pytest.ini file generator: ${PYTEST_INI_GENERATOR}"
    cleanup
    exit 1
fi

# verify that the common file generator tool exists
if [[ ! -x ${COMMON_FILE_GENERATOR} ]] ; then
    >&2 echo "ERROR: failed to locate executable common file generator: ${COMMON_FILE_GENERATOR}"
    cleanup
    exit 1
fi

echo "Running bss_tavern_api_test..."

# retrieve Keycloak authentication token for session
TOKEN=$(get_auth_access_token)
TOKEN_RET=$?
if [[ ${TOKEN_RET} -ne 0 ]] ; then
    cleanup
    exit 1
fi

# generate pytest.ini configuration file
GENERATE_PYTEST_INI_CMD="${PYTEST_INI_GENERATOR} --file ${PYTEST_INI_PATH}"
timestamp_print "Running '${GENERATE_PYTEST_INI_CMD}'..."
eval "${GENERATE_PYTEST_INI_CMD}"
GENERATE_PYTEST_INI_RET=$?
if [[ ${GENERATE_PYTEST_INI_RET} -ne 0 ]] ; then
    >&2 echo "ERROR: pytest.ini file generator failed with error code: ${GENERATE_PYTEST_INI_RET}"
    cleanup
    exit 1
else
    if [[ ! -r ${PYTEST_INI_PATH} ]] ; then
        >&2 echo "ERROR: failed to generate readable pytest.ini file"
        cleanup
        exit 1
    fi
fi

# generate Tavern common.yaml configuration file
GENERATE_COMMON_FILE_CMD="${COMMON_FILE_GENERATOR} --base_url ${API_TARGET} --file ${COMMON_FILE_PATH} --access_token ${TOKEN} --verify ${VERIFY}"
timestamp_print "Running '${GENERATE_COMMON_FILE_CMD}'..."
eval "${GENERATE_COMMON_FILE_CMD}"
GENERATE_COMMON_FILE_RET=$?
if [[ ${GENERATE_COMMON_FILE_RET} -ne 0 ]] ; then
    >&2 echo "ERROR: common file generator failed with error code: ${GENERATE_COMMON_FILE_RET}"
    cleanup
    exit 1
else
    if [[ ! -r ${COMMON_FILE_PATH} ]] ; then
        >&2 echo "ERROR: failed to generate readable Tavern common.yaml file"
        cleanup
        exit 1
    fi
fi

# execute Tavern tests in the hms-pytest container with pytest
PYTEST_CMD="${PYTEST_PATH} --tavern-global-cfg=${COMMON_FILE_PATH} ${BSS_TEST_DIR}"
timestamp_print "Running '${PYTEST_CMD}'..."
eval "${PYTEST_CMD}"
TAVERN_RET=$?
if [[ ${TAVERN_RET} -ne 0 ]] ; then
    echo "FAIL: bss_tavern_api_test ran with failures"
    cleanup
    exit 1
else
    echo "PASS: bss_tavern_api_test passed!"
    cleanup
    exit 0
fi
