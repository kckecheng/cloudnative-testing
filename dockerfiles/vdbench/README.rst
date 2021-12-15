vdbench
=========

Build
-------

::

  docker build -t tools/vdbench .

Usage
------

::

  docker run --rm -it tools/vdbench -tf
  docker run --rm -it -v $PWD:/config -v /mnt:/data tools/vdbench -f /config/vdbench_fs.sample

