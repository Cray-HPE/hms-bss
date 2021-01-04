#!/usr/bin/env bash

# When it's time to update the version information, use this script as a person

echo "Detecting the version from the changelog..."
VERSION=$(grep -m1 \#\# ../CHANGELOG.md | sed -e "s/\].*$//" |sed -e "s/^.*\[//")
echo "Version = $VERSION"

echo "Add the version to the VERSION file"
echo ${VERSION} > ../.version

# tag it
git commit -am "Version $VERSION"
git tag -a "v$VERSION" -m "Version $VERSION"
git push
git push --tags
