# A sample Dockerfile
#
# VERSION 0.1.7
#
# List of things installed by this Dockerfile
# foo
# bar
# baz
#
# To build this, you must provide my-cool-debian-package.deb in the same
# directory as the Dockerfile, or, use the build script which will download it
# for you

FROM docker.epeters.com/base_images/eds-ubuntu:1.1.0
MAINTAINER Edward Petersen edward.petersen@gmail.com

# install that cool debian package and remove the installation artifact
ADD ./my-cool-debian-package.deb /tmp/
RUN dpkg -i /tmp/my-cool-debian-package.deb && rm -f /tmp/my-cool-debian-package.deb

# create /etc/alternatives entries
# THIS SHOULD REALLY BE DONE AS AN "&&" WITH THE ABOVE RUN TO CONSOLIDATE COMMANDS, BUT BROKEN OUT HERE FOR CLARITY
RUN update-alternatives --install /usr/bin/cool cool /usr/local/cool/bin/cool 1000
RUN update-alternatives --auto cool

# install some apt-get stuff
RUN apt-get update && apt-get install -y python=2.7.5-5ubuntu3 python-dev=2.7.5-5ubuntu3 python-virtualenv=1.11.4-1 python-pip=1.5.4-1

# add pip configuration file and pypirc file
ADD /pip.conf /root/.pip/pip.conf
ADD /pypirc /root/.pypirc

# bash command, as this image is merely a base for further work
CMD ["/bin/bash"]
