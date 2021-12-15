ltp
====

**Note**: it is not quite recommended to run ltp from container since the kernel header installed within the centos container might be different from the os you are using.

Build
-------

::

  docker build -t tools/ltp .

Usage
------

::

  docker run --rm -it tools/ltp
  docker run --rm -it -v /mnt:/data tools/ltp -p -q -l /data/result.log -o /data/output.log -C /data/err.log -d /data

