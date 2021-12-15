sysbench
==========

Build
-------

::

  docker build -t tools/sysbench .

Usage
------

::

  docker run --rm -it tools/sysbench
  docker run --rm -it tools/sysbench cpu run
  docker run --rm -it tools/sysbench memory run
  docker run --rm -it tools/sysbench fileio help
  docker run --rm -it -v /mnt:/data tools/sysbench fileio <options> prepare
  docker run --rm -it -v /mnt:/data tools/sysbench fileio <options> run
  docker run --rm -it -v /mnt:/data tools/sysbench fileio <options> cleanup

