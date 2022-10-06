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
        raise exceptions.RegexAccessError(f"No match for regex: {expression}")

    return {"regex": Box(match.groupdict())}


def validate_yaml_simple(response, expected):
    data = yaml.safe_load(response)

    assert check_keys_match_recursive(expected, expected, [])

    return {}