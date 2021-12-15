bonnie++
==========

Build
-------

::

  docker build -t tools/bonnie .

Usage
------

::

  docker run --rm -it tools/bonnie .
  docker run --rm -v $PWD:/data -it tools/bonnie -d /data -u 0
