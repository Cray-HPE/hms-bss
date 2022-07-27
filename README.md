This is the repository for the HMS Boot Script Service (BSS) code.

It includes a swagger.yaml file for the service REST API specification, along with all of the code to implement the stateless
service itself.

This service should contain just what is needed to provide boot arguments (initrd, kargs, etc) and Level 2 boot services for
static images.

This code has been refactored from the old hms-netboot code for bootargsd and associated components created for the Q4 Redfish
and Q1 systems management deep dive demos.

### BSS CT Testing

In addition to the service itself, this repository builds and publishes cray-bss-test images containing tests that verify BSS
on live Shasta systems. The tests are invoked via helm test as part of the Continuous Test (CT) framework during CSM installs
and upgrades. The version of the cray-bss-test image (vX.Y.Z) should match the version of the cray-bss image being tested, both
of which are specified in the helm chart for the service.