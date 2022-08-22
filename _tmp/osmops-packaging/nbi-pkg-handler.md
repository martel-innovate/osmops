Uploaded corrupt tgz stream to `vnf_packages_content`. This way I was
able to get a clue of what NBI does by looking at the Python exception
stack trace in the logs.

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content \
  -H 'Authorization: Bearer PxtGMSVAy1LJ2COGAFiYm7ctXI22CG7i' \
  -H 'Accept: application/json' -H 'Content-Type: application/gzip' \
  -H 'Content-Filename: openldap_knf.tar.gz' \
  -H 'Content-File-MD5: 2a7d74587151e9fd0c1fd727003b8a1b' \
  -d @../osm-pkgs/openldap_knf.tar.gz
```

Notice `curl` switch should've been `--data-binary`, not `-d` which
probably treats the file as text.

```console
$ multipass shell osm2
$ kubectl -n osm logs nbi-6f5fd9ff89-8xpkw
```

```log
2022-06-16T16:12:22 INFO nbi.server _cplogging.py:213 [16/Jun/2022:16:12:22]  CRITICAL: Exception Compressed file ended before the end-of-stream marker was reached
Traceback (most recent call last):
  File "/usr/lib/python3/dist-packages/osm_nbi/nbi.py", line 1417, in default
    completed = self.engine.upload_content(
  File "/usr/lib/python3/dist-packages/osm_nbi/engine.py", line 277, in upload_content
    return self.map_topic[topic].upload_content(
  File "/usr/lib/python3/dist-packages/osm_nbi/descriptor_topics.py", line 324, in upload_content
    tar = tarfile.open(mode="r", fileobj=file_pkg)
  File "/usr/lib/python3.8/tarfile.py", line 1603, in open
    return func(name, "r", fileobj, **kwargs)
  File "/usr/lib/python3.8/tarfile.py", line 1674, in gzopen
    t = cls.taropen(name, mode, fileobj, **kwargs)
  File "/usr/lib/python3.8/tarfile.py", line 1651, in taropen
    return cls(name, mode, fileobj, **kwargs)
  File "/usr/lib/python3.8/tarfile.py", line 1514, in __init__
    self.firstmember = self.next()
  File "/usr/lib/python3.8/tarfile.py", line 2318, in next
    tarinfo = self.tarinfo.fromtarfile(self)
  File "/usr/lib/python3.8/tarfile.py", line 1104, in fromtarfile
    buf = tarfile.fileobj.read(BLOCKSIZE)
  File "/usr/lib/python3.8/gzip.py", line 292, in read
    return self._buffer.read(size)
  File "/usr/lib/python3.8/_compression.py", line 68, in readinto
    data = self.read(len(byte_view))
  File "/usr/lib/python3.8/gzip.py", line 479, in read
    if not self._read_gzip_header():
  File "/usr/lib/python3.8/gzip.py", line 437, in _read_gzip_header
    self._read_exact(extra_len)
  File "/usr/lib/python3.8/gzip.py", line 416, in _read_exact
    raise EOFError("Compressed file ended before the "
EOFError: Compressed file ended before the end-of-stream marker was reached
```
