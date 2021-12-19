genprofile
============

Generate profiles based on a defined template and an option file.

Build
-------

::

  go build .

Usage
-------

How to create a profile template
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

A profile template should be defined by followeing the golang text/template syntax. Below is an example for fio:

::

  [global]
  filename={{ .filename }}
  size={{ .size }}

  ioengine={{ .ioengine }}

  iodepth={{ .iodepth }}

  time_based=1
  runtime={{ .runtime }}

  [{{ .ioengine }}-{{ .rw }}-{{ .bs }}]
  rw={{ .rw }}
  bs={{ .bs }}

How to create an option file
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

An option file define parameters which are populated into a profile template. It uses yaml syntax and supports single yaml document. In the meanwhile, only literal variable (like int, float, string, boolean, etc.) and list are supported, advanced expressions like dict, list of list, dict of dict, etc. are all not supported - please use multiple option files for such scenarios.

Below is an option file example for fio:

::

  ---
  filename: /mnt/fio.out
  size: 1g
  iodepth: 1
  runtime: 300

  rw:
  - read
  - write

  ioengine:
  - sync
  - libaio

  bs:
  - 4k
  - 128k
  - 4m

How to generate test profiles based on a template and an option file
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The above fio profile template and option file are used:

::

  # ./genprofile --help
  # ./genprofile -t fio.profile.tmpl -o fio.options.yaml -p fio
  # ls -l
  -rw-r--r-- 1 vagrant vagrant 120 Dec 19 14:12 fio-test-0.profile
  -rw-r--r-- 1 vagrant vagrant 122 Dec 19 14:12 fio-test-10.profile
  -rw-r--r-- 1 vagrant vagrant 126 Dec 19 14:12 fio-test-11.profile
  -rw-r--r-- 1 vagrant vagrant 124 Dec 19 14:12 fio-test-1.profile
  -rw-r--r-- 1 vagrant vagrant 122 Dec 19 14:12 fio-test-2.profile
  -rw-r--r-- 1 vagrant vagrant 126 Dec 19 14:12 fio-test-3.profile
  -rw-r--r-- 1 vagrant vagrant 124 Dec 19 14:12 fio-test-4.profile
  -rw-r--r-- 1 vagrant vagrant 128 Dec 19 14:12 fio-test-5.profile
  -rw-r--r-- 1 vagrant vagrant 126 Dec 19 14:12 fio-test-6.profile
  -rw-r--r-- 1 vagrant vagrant 130 Dec 19 14:12 fio-test-7.profile
  -rw-r--r-- 1 vagrant vagrant 120 Dec 19 14:12 fio-test-8.profile
  -rw-r--r-- 1 vagrant vagrant 124 Dec 19 14:12 fio-test-9.profile
  # rm -rf fio*.profile
  # ./genprofile -t fio.go.tmpl.sample -o fio.options.yaml.sample -f ioengine -f rw -f bs -p fio
  # ls -l
  -rw-r--r-- 1 vagrant vagrant 128 Dec 19 14:14 fio-libaio-read-128k.profile
  -rw-r--r-- 1 vagrant vagrant 124 Dec 19 14:14 fio-libaio-read-4k.profile
  -rw-r--r-- 1 vagrant vagrant 124 Dec 19 14:14 fio-libaio-read-4m.profile
  -rw-r--r-- 1 vagrant vagrant 130 Dec 19 14:14 fio-libaio-write-128k.profile
  -rw-r--r-- 1 vagrant vagrant 126 Dec 19 14:14 fio-libaio-write-4k.profile
  -rw-r--r-- 1 vagrant vagrant 126 Dec 19 14:14 fio-libaio-write-4m.profile
  -rw-r--r-- 1 vagrant vagrant 124 Dec 19 14:14 fio-sync-read-128k.profile
  -rw-r--r-- 1 vagrant vagrant 120 Dec 19 14:14 fio-sync-read-4k.profile
  -rw-r--r-- 1 vagrant vagrant 120 Dec 19 14:14 fio-sync-read-4m.profile
  -rw-r--r-- 1 vagrant vagrant 126 Dec 19 14:14 fio-sync-write-128k.profile
  -rw-r--r-- 1 vagrant vagrant 122 Dec 19 14:14 fio-sync-write-4k.profile
  -rw-r--r-- 1 vagrant vagrant 122 Dec 19 14:14 fio-sync-write-4m.profile