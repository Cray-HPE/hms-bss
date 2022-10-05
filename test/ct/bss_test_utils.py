import socket
from box import Box

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
