# MIT License
#
# (C) Copyright [2022] Hewlett Packard Enterprise Development LP
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

import socket
from box import Box

import json
import jmespath
import logging
import re

from tavern.util import exceptions
from tavern.util.dict_util import recurse_access_key,check_keys_match_recursive

import yaml

logger = logging.getLogger(__name__)

def build_hsm_ethernet_interface_for_test_container():
    # Determine the IP of the container
    hostname = socket.gethostname()
    ip_address = socket.gethostbyname(hostname)

    # Build up a HSM EthernetInterface with this containers IP address
    return {
        "MACAddress": "0e:ff:ff:ff:ff:fe",
        "ComponentID": "x9999c0s1b0n0",
        "Description": "Created by BSS CT Tavern Tests",
        "IPAddresses": [{
            "IPAddress": ip_address
        }]
    }

def save_ip_address_of_test_container(response):
    # Determine the IP of the container
    hostname = socket.gethostname()
    ip_address = socket.gethostbyname(hostname)

    return Box({"test_container_ip_address": ip_address})

def better_validate_regex(response, expression):
    logger.debug("Matching %s with %s", response.text, expression)

    match = re.search(expression, response.text)
    if match is None:
        raise exceptions.RegexAccessError(f"No match for regex '{expression}' in response:\n{response.text}")

    return {"regex": Box(match.groupdict())}


def validate_yaml_simple(response, expected):
    response_data = yaml.safe_load(response.text)

    check_keys_match_recursive(response_data, expected, [])

    return {}
