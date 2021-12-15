iperf3
========

Build
-------

::

  docker build -t tools/iperf3 .

Usage
------

**Server**:

::

  docker run --rm -it --network host tools/iperf3 -s -B 192.168.1.10

**Client**:

::

  docker run --rm -it --network host tools/iperf3 -c 192.168.1.10 -p 5201 -B 192.168.1.11

