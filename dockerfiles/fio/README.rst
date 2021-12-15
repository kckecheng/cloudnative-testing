fio
====

Build
-------

::

  docker build -t tools/fio .

Usage
------

::

  docker run --rm -it tools/fio
  docker run --rm -it -v $PWD:/data tools/fio -name=test1 -rw=randrw -filename=/data/a.out -size=100M -runtime=300s

