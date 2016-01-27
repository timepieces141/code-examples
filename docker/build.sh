#!/bin/bash

set -e

# name and version - these could be parametized with sane defaults, depening on
# where the script lives
NAME=docker.epeters.com/myrepo/my-image
VERSION=0.1.7

# this docker image requires a debian package
[ -f my-cool-debian-package.deb ] || wget http://www.awesome.com/my-cool-debian-package.deb

# build
docker build --force-rm=true --no-cache -t $NAME:$VERSION .

# echo the commands to tag and push this
echo "To push this image:"
echo "  docker tag $NAME:$VERSION $NAME:latest"
echo "  docker push $NAME:$VERSION $NAME:latest"
