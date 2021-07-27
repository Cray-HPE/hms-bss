# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.10.2] - 2021-07-21

### Changed

- Moved all stash references to GitHub.

## [1.10.1] - 2021-07-19

### Changed

- Add support for building within the CSM Jenkins.

## [1.10.0] - 2021-06-28

### Security

- CASMHMS-4898 - Updated base container images for security updates.

## [1.9.1] - 2021-04-14

### Changed

- Updated Dockerfiles to pull base images from Artifactory instead of DTR.

## [1.9.0] - 2021-01-26

### Changed

- Update copyrights.

## [1.8.0] - 2021-01-14

### Changed

- Updated license file.

## [1.7.1] - 2020-10-29

### Security

- CASMHMS-4148 - Update grpc go module for grpc/etcd incompatibility issue.

## [1.7.0] - 2020-10-20

### Security

- CASMHMS-4105 - Updated base Golang Alpine image to resolve libcrypto vulnerability.

## [1.6.1] - 2020-04-27

### Changed

- CASMHMS-2975 - Updated hms-hmetcd to use trusted baseOS.

## [1.6.0] - 2019-12-11

### Changed

- Split this module into a separate package from hms-common

## [1.5.5] - 2019-12-02

### Added

- The SNMPAuthPass and SNMPPrivPass fields to the CompCredentials struct

## [1.5.4] - 2019-11-22

### Added

- Definitions for subroles

## [1.5.3] - 2019-10-04

### Added

- Extended securestorage mock Vault adapter to also function as a more
  generalized storage mechanism for complex unit test case scenarios.  All
  existing functionality is preserved. Use as a generalized store requires
  initializing InputLookup.Key (or InputLookupKeys.KeyPath) and setting
  LookupNum (or LookupKeysNum) to -1.

## [1.5.2] - 2019-10-03

### Fixed

- Synced up with the HMS Component Naming Convention.  Note that this introduces
some incompatibilties with previous versions.

## [1.5.1] - 2019-09-18

### Added

- Added the "Locked" component flag to base.

## [1.5.0] - 2019-08-13

### Added

- Added SMNetManager already in use in REDS/MEDS to common library.

## [1.4.2] - 2019-08-07

### Fixed

- Segmentation fault in decode logic of secure store when a nil structure (i.e., no results) are returned from Vault.

## [1.4.1] - 2019-08-01

### Added

- Management role to base

## [1.4.0] - 2019-07-30

### Added

- Added the securestorage package that performs basic actions (Store, Lookup, etc) on a chosen secure backing store. The initial list of backing stores only includes Vault.
- Added the compcredentials package that performs common component credential operations with the securestorage package.

## [1.3.0] - 2019-07-08

### Added

- Added HTTP library that utilizes retryablehttp to perform HTTP operations and optionally unmarshal the returned value into an interface.

## [1.2.0] - 2019-05-18

### Changed

- Added changes for CabinetPDU support
- Tweak to state change table

## [1.1.0] - 2019-05-13

### Removed

- Removed `hmsds`, `sharedtest`, `sm`, and `redfish` packages from this repo as they are actually SMD specific and therefore belong in that repo.

## [1.0.0] - 2019-05-13

### Added

- This is the initial release of the `hms-common` repo. It contains everything that was in `hms-services` at the time with the major exception of being `go mod` based now.

### Added

### Changed

### Deprecated

### Removed

### Fixed

### Security
