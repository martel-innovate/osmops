Intercepting OSM client messages
--------------------------------

OSM client only uses HTTPS. Since the code doesn't verify the server
identity, it's easy to set up a man-in-the-middle attack to observe
HTTP traffic between the client and the server---i.e. the NBI, OSM's
north bound interface. To do that, you can install `stunnel` and tweak
the config in this dir---have a look at `stunnel-mitm-proxy.conf`.

An easier way to catch HTTP messages is to tweak the client code to
make it use plain HTTP without TLS. To see how diff the two Python
files in this dir.

To monitor HTTP traffic in the OSM host, you have to replace

    /usr/lib/python3/dist-packages/osmclient/sol005/client.py

with `client.py` in this dir, then

```bash
    $ multipass shell osm
    [osm]$ sudo tcpdump -i cni0 -s 1024 -A port 80
```

and use the OSM client, e.g.

```bash
    $ multipass shell osm
    [osm]$ osm ns-op-list ldap
```
