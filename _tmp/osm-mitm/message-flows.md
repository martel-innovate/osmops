OSM client HTTP message flows
-----------------------------
> Or what the heck OSM client does under the bonnet.


### Getting an auth token

This happens every time you run an `osm` command, i.e. tokens aren't cached!
Example flow

```http
POST /osm/admin/v1/tokens HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Content-Length: 65

{"username": "admin", "password": "admin", "project_id": "admin"}
```

```http
HTTP/1.1 200 OK
Server: nginx/1.14.0 (Ubuntu)
Date: Wed, 08 Sep 2021 17:52:11 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 549
Connection: keep-alive
Www-Authenticate: Bearer realm="Needed a token or Authorization http header"
Location: /osm/admin/v1/tokens/TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2
Set-Cookie: session_id=072faf1c629771cdad9133c133fe8bee1202f258; expires=Wed, 08 Sep 2021 18:52:11 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "issued_at": 1631123531.1251214,
    "expires": 1631127131.1251214,
    "_id": "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2",
    "id": "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2",
    "project_id": "fada443a-905c-4241-8a33-4dcdbdac55e7",
    "project_name": "admin",
    "username": "admin",
    "user_id": "5c6f2d64-9c23-4718-806a-c74c3fc3c98f",
    "admin": true,
    "roles": [
        {
            "name": "system_admin",
            "id": "cb545e44-cd2b-4c0b-93aa-7e2cee79afc3"
        }
    ],
...
```

Notice the token is valid for an hour:

    issued_at = Wednesday, 8 September 2021 17:52:11.125 (GMT)
    expires   = Wednesday, 8 September 2021 18:52:11.125 (GMT)


### Getting the history of operations on an NS instance

OSM client command

```bash
$ osm ns-op-list ldap
ERROR: ns 'ldap' not found
```

HTTP request

```http
GET /osm/nslcm/v1/ns_instances_content HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer qIFJhw2JkGbgBToJiuKgYNSKuFgnQlYX
```

HTTP response

```http
.HTTP/1.1 200 OK
Server: nginx/1.14.0 (Ubuntu)
Date: Thu, 09 Sep 2021 14:19:53 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 3
Connection: keep-alive
Set-Cookie: session_id=321df9a60ac919141432e830cfcd8cb306f31877; expires=Thu, 09 Sep 2021 15:19:53 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

[]
```


### Creating a VIM account

OSM client command

```bash
$ osm vim-create --name openvim-site \
    --auth_url http://10.10.10.10:9080/openvim \
    --account_type openvim --description "Openvim site" \
    --tenant osm --user dummy --password dummy
59b92c04-29fa-42a7-923e-63322240b80e
```

HTTP request

```http
POST /osm/admin/v1/vim_accounts HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2
Content-Length: 196

{"name": "openvim-site", "vim_type": "openvim", "description": "Openvim site", "vim_url": "http://10.10.10.10:9080/openvim", "vim_user": "dummy", "vim_password": "dummy", "vim_tenant_name": "osm"}
```

HTTP response

```http
HTTP/1.1 202 Accepted
Server: nginx/1.14.0 (Ubuntu)
Date: Wed, 08 Sep 2021 17:52:11 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 108
Connection: keep-alive
Location: /osm/admin/v1/vim_accounts/59b92c04-29fa-42a7-923e-63322240b80e
Set-Cookie: session_id=4cd3ace1f2635ca888bbbb6d24a5905540345809; expires=Wed, 08 Sep 2021 18:52:11 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "id": "59b92c04-29fa-42a7-923e-63322240b80e",
    "op_id": "59b92c04-29fa-42a7-923e-63322240b80e:0"
}
```

Notice VIM account names have to be unique. In fact, OSM NBI enforces that.
If you try creating another VIM account with the same name, you get an error:

```bash
$ curl localhost/osm/admin/v1/vim_accounts \
    -v -X POST \
    -H 'Authorization: Bearer TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2' \
    -H 'Content-Type: application/yaml' \
    -d'{"name": "openvim-site", "vim_type": "openvim", "description": "Openvim site", "vim_url": "http://10.10.10.10:9080/openvim", "vim_user": "dummy", "vim_password": "dummy", "vim_tenant_name": "osm"}'

...
HTTP/1.1 409 Conflict
...
---
code: CONFLICT
detail: name 'openvim-site' already exists for vim_accounts
status: 409
```


### KNF service onboarding and instantiation

Example flow from OSM manual section: 5.6.5.1 KNF Helm Chart

- https://osm.etsi.org/docs/user-guide/05-osm-usage.html#knf-helm-chart

**NB**. Download package tarballs from:

- https://osm-download.etsi.org/ftp/Packages/examples/

the repo in section 5.6.5.1 is outdated.


#### Onboarding

OSM client command to upload a package with a VNFD for an Open LDAP KNF:

```bash
$ osm nfpkg-create openldap_knf.tar.gz
```

HTTP request

```http
POST /osm/vnfpkgm/v1/vnf_packages_content HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/gzip
Authorization: Bearer nOcHehp8wJcxSze8lJFzEKUBTI9iOdgk
Content-Filename: openldap_knf.tar.gz
Content-File-MD5: 6f10bac4462725413f4e14f185619ead
Content-Length: 449

......._..openldap_knf.tar...
```

HTTP response

```http
HTTP/1.1 201 Created
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 10:07:25 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 53
Connection: keep-alive
Location: /osm/vnfpkgm/v1/vnf_packages_content/d506d18f-0738-42ab-8b45-cfa98da38e7a
Set-Cookie: session_id=78799fcfb8463bfb1da410066fefa6c89e9ed1ec; expires=Fri, 10 Sep 2021 11:07:24 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "id": "d506d18f-0738-42ab-8b45-cfa98da38e7a"
}
```

Notice OSM NBI enforces uniqueness of VNFD IDs. If you try uploading another
package with a VNFD having the same ID as the one we've just uploaded, OSM
NBI will complain loudly:

```http
HTTP/1.1 409 Conflict
...
{
    "code": "CONFLICT",
    "status": 409,
    "detail": "vnfd with id 'openldap_knf' already exists for this project"
}
```

OSM client command to upload a package with a NSD for the Open LDAP KNF
defined by the previous package:

```bash
$ osm nspkg-create openldap_ns.tar.gz
```

HTTP request

```http
POST /osm/nsd/v1/ns_descriptors_content HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/gzip
Authorization: Bearer G3zXf7lFs91YmewaUQnU5yLc0hOMUyBD
Content-Filename: openldap_ns.tar.gz
Content-File-MD5: 38f617220c88c1c32a8c9be55d781041
Content-Length: 977

......._..openldap_ns.tar..
```

HTTP response

```http
HTTP/1.1 201 Created
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 10:10:37 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 53
Connection: keep-alive
Location: /osm/nsd/v1/ns_descriptors_content/aba58e40-d65f-4f4e-be0a-e248c14d3e03
Set-Cookie: session_id=e9ee44a81f693d768ffe4b7265ab8cfbcef078c0; expires=Fri, 10 Sep 2021 11:10:35 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "id": "aba58e40-d65f-4f4e-be0a-e248c14d3e03"
}
```

Notice OSM NBI enforces uniqueness of NSD IDs. If you try uploading another
package with a NSD having the same ID as the one we've just uploaded, OSM
NBI will complain loudly:

```http
HTTP/1.1 409 Conflict
...
{
    "code": "CONFLICT",
    "status": 409,
    "detail": "nsd with id 'openldap_ns' already exists for this project"
}
```

#### NS instantiation

OSM client command to create an NS instance using the OpenLDAP chart uploaded
by the previous commands. Notice we use the VIM account name and OSM client
looks up the corresponding ID for us. Notice the name-ID lookup works because
OSM NBI enforces VIM name uniqueness---see earlier note about it.

```bash
$ osm ns-create --ns_name ldap --nsd_name openldap_ns --vim_account mylocation1
0335c32c-d28c-4d79-9b94-0ffa36326932
```

[HTTP message flow](./message-flow.ns-create.md)

OSM client command to create a second NS instance from the same chart but
this time with two replicas. Notice we use the VIM account ID this time and
OSM will use that ID as is. (**Question**: what's the algo to determine if
a string is a name or an ID?! Possibly another sore point here...)

```bash
$ osm ns-create --ns_name ldap2 --nsd_name openldap_ns \
    --vim_account 4a4425f7-3e72-4d45-a4ec-4241186f3547 \
    --config '{additionalParamsForVnf: [{"member-vnf-index": "openldap", additionalParamsForKdu: [{ kdu_name: "ldap", "additionalParams": {"replicaCount": "2"}}]}]}'
136fcc46-c363-4d74-af14-c115fff7d80a
```

[HTTP message flow](./message-flow.ns-create2.md)

Notice OSM NBI doesn't enforce uniqueness of NS names. In fact, it lets you
happily duplicate e.g. the `ldap` name we created earlier:

```bash
$ curl localhost/osm/nslcm/v1/ns_instances_content \
    -v -X POST \
    -H 'Authorization: Bearer 0WhgBufy1Wt82NbF9OsmftwpRfcsV4sU' \
    -H 'Content-Type: application/yaml' \
    -d'{"nsdId": "aba58e40-d65f-4f4e-be0a-e248c14d3e03", "nsName": "ldap", "nsDescription": "default description", "vimAccountId": "4a4425f7-3e72-4d45-a4ec-4241186f3547"}'
...
HTTP/1.1 201 Created
...
---
id: 794ef9a2-8bbb-42c1-869a-bab6422982ec
nslcmop_id: 0fdfaa6a-b742-480c-9701-122b3f732e4f
```

#### NS upgrade

OSM client command to upgrade the first LDAP NS we created earlier. Notice
OSM client looks up the instance ID from the name we specify in the command
line (`ldap`), but this is **not** a good idea since, as noted earlier,
NS instance names aren't unique. To avoid wreaking havoc we should always
use NS instance IDs but I don't think OSM client actually supports that?

```bash
$ osm ns-action ldap --vnf_name openldap --kdu_name ldap --action_name upgrade --params '{kdu_model: "stable/openldap:1.2.2"}'
5c6e4a0d-6238-4aa8-9147-e4e738bf16f4
```

[HTTP message flow](./message-flow.ns-action.upgrade.md)

This upgrade op eventually failed---OSM client always returns 0 since the
actual op gets executed server-side asynchronously. In fact, here's the
instance history after the op completed server-side:

```
 ID                                     action          start                   end            status
------------------------------------  -----------  --------------------  --------------------  ------
0c5464da-df42-498e-b306-76d470b76a0d  instantiate  Sep-10-2021 14:40:26  Sep-10-2021 14:41:53  OK
5c6e4a0d-6238-4aa8-9147-e4e738bf16f4  action       Sep-10-2021 16:53:57  Sep-10-2021 16:54:27  Failed
```

The error:

```
FAILED Executing kdu upgrade: Error executing command:
/usr/local/bin/helm3 upgrade stable-openldap-1-2-3-0098084071 stable/openldap --namespace fada443a-905c-4241-8a33-4dcdbdac55e7 --atomic --output yaml  --timeout 1800s  --version 1.2.2

Output: Error: UPGRADE FAILED: an error occurred while rolling back the release.
original upgrade error:
cannot patch "stable-openldap-1-2-3-0098084071" with kind Service: Service "stable-openldap-1-2-3-0098084071" is invalid: spec.clusterIP: Invalid value: "": field is immutable
```

OSM client command to up the number of replicas of the second LDAP NS we
created earlier. Notice this is slightly different than the example in the
OSM manual since I got rid of the `kdu_model` param that doesn't work as you
can see from the outcome of running the previous upgrade action.

```bash
$ osm ns-action ldap2 --vnf_name openldap --kdu_name ldap --action_name upgrade --params '{"replicaCount": "3"}'
f92f746f-c10a-448e-84e1-3acfd8b684cb
```

[HTTP message flow](./message-flow.ns-action.upgrade2.md)

But in the end this upgrade action didn't work either. Exactly the same error
as before got recorded in the NS instance action history.
