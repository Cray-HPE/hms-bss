This is the location for the HMS Level 2 boot script service code.

It should include a Swagger.yaml or .json file for the service REST API,
along with all of the code to implement the stateless service itself.

This service should contain just what is needed to provide boot arguments (initrd, kargs, etc) Level 2 boot services for static 
images.

This code will be refactored from the old hms-netboot code for bootargsd and associated components created for the Q4 Redfish and Q1 
systems management deep dive demos.

### BSS CT Testing

This repository builds and publishes hms-bss-ct-test RPMs along with the service itself containing tests that verify BSS on the
NCNs of live Shasta systems. The tests require the hms-ct-test-base RPM to also be installed on the NCNs in order to execute.
