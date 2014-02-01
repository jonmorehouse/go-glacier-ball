Amazon Glacier File Backup
==========================

Functionality
-------------

* accepts line delimited file paths from stdin
* uploads to amazon glacier in 1gb tarball archives
* cleans up after itself (each upload / archive is a go worker so we can manage memory overhead)

Example
-------

```
  find . -type f | glaciar-ball
```

Output
------

* 01-30-2014-01.tgz
* 01-30-2014-02.tgz
* ...



