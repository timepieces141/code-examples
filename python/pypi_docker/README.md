Docker PyPi Mirror
==================

This PyPi server runs by default expecting the packages to be (inside the container) at the path /pypi. You should mount a static volume to this in order to maintain the packages upon reboot. Here's the command to run to start the container:

```
docker run -d -h pypi.local -v /home/epeters/packages:/pypi -p 80:8000 --name pypi pypi
```

Of course, adding your own hostname (-h), your own port selections, and your own path on the local system where the volume should be mounted (persistent store of the packages).
