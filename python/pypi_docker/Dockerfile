#
# PyPi Server
#

FROM ubuntu:trusty
MAINTAINER Edward Petersen edward.petersen@gmail.com

RUN apt-get update && apt-get install -y python-pip && pip install pypiserver passlib

VOLUME ["/pypi"]
EXPOSE 8000

# Default command
CMD ["pypi-server", "-p", "8000", "/pypi"]
