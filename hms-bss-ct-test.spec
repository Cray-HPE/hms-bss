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

Name: hms-bss-ct-test
License: MIT
Summary: HMS CT tests for the Boot Script Service (BSS)
Group: System/Management
Version: %(cat .version) 
Release: %(echo ${BUILD_METADATA})
Source: %{name}-%{version}.tar.bz2
Vendor: Hewlett Packard Enterprise
Requires: hms-ct-test-base >= 1.9.0

# name of this repository
%define REPO hms-bss

# test installation location
%define TEST_DIR /opt/cray/tests

%description
This is a collection of post-install CT tests for BSS.

%prep
%setup -q

%build
# Categories of CT tests to install
TEST_BUCKETS=(
    ncn-smoke
    ncn-functional
    ncn-long
    ncn-destructive
    ncn-resources
    remote-smoke
    remote-functional
    remote-long
    remote-destructive
    remote-resources
)

echo "Current directory is: ${PWD}..."

echo "Searching for CT tests..."
for BUCKET in ${TEST_BUCKETS[@]} ; do
    find . -name "*${BUCKET}*" -exec mkdir -p %{buildroot}%{TEST_DIR}/${BUCKET}/hms/%{REPO}/ \; \
       -exec cp -v {} %{buildroot}%{TEST_DIR}/${BUCKET}/hms/%{REPO}/ \;
done

%files

# CT tests
%dir %{TEST_DIR}
%{TEST_DIR}/*

%changelog
* Wed Oct 27 2021 Mitch Schooler <mitchell.schooler@hpe.com>
- Moved CT test packaging from hms-test to individual service repos for their own RPM builds.
