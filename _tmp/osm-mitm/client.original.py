# Copyright 2018 Telefonica
#
# All Rights Reserved.
#
#    Licensed under the Apache License, Version 2.0 (the "License"); you may
#    not use this file except in compliance with the License. You may obtain
#    a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
#    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
#    License for the specific language governing permissions and limitations
#    under the License.

"""
OSM SOL005 client API
"""

# from osmclient.v1 import vca
from osmclient.sol005 import vnfd
from osmclient.sol005 import nsd
from osmclient.sol005 import nst
from osmclient.sol005 import nsi
from osmclient.sol005 import ns
from osmclient.sol005 import vnf
from osmclient.sol005 import vim
from osmclient.sol005 import wim
from osmclient.sol005 import package
from osmclient.sol005 import http
from osmclient.sol005 import sdncontroller
from osmclient.sol005 import project as projectmodule
from osmclient.sol005 import user as usermodule
from osmclient.sol005 import role
from osmclient.sol005 import pdud
from osmclient.sol005 import k8scluster
from osmclient.sol005 import vca
from osmclient.sol005 import repo
from osmclient.sol005 import osmrepo
from osmclient.sol005 import subscription
from osmclient.common import package_tool
import json
import logging


class Client(object):
    def __init__(
        self,
        host=None,
        so_port=9999,
        user="admin",
        password="admin",
        project="admin",
        **kwargs
    ):

        self._user = user
        self._password = password
        self._project = project
        self._project_domain_name = kwargs.get("project_domain_name")
        self._user_domain_name = kwargs.get("user_domain_name")
        self._logger = logging.getLogger("osmclient")
        self._auth_endpoint = "/admin/v1/tokens"
        self._headers = {}
        self._token = None
        if len(host.split(":")) > 1:
            # backwards compatible, port provided as part of host
            self._host = host.split(":")[0]
            self._so_port = host.split(":")[1]
        else:
            self._host = host
            self._so_port = so_port

        self._http_client = http.Http(
            "https://{}:{}/osm".format(self._host, self._so_port), **kwargs
        )
        self._headers["Accept"] = "application/json"
        self._headers["Content-Type"] = "application/yaml"
        http_header = [
            "{}: {}".format(key, val) for (key, val) in list(self._headers.items())
        ]
        self._http_client.set_http_header(http_header)

        self.vnfd = vnfd.Vnfd(self._http_client, client=self)
        self.nsd = nsd.Nsd(self._http_client, client=self)
        self.nst = nst.Nst(self._http_client, client=self)
        self.package = package.Package(self._http_client, client=self)
        self.ns = ns.Ns(self._http_client, client=self)
        self.nsi = nsi.Nsi(self._http_client, client=self)
        self.vim = vim.Vim(self._http_client, client=self)
        self.wim = wim.Wim(self._http_client, client=self)
        self.sdnc = sdncontroller.SdnController(self._http_client, client=self)
        self.vnf = vnf.Vnf(self._http_client, client=self)
        self.project = projectmodule.Project(self._http_client, client=self)
        self.user = usermodule.User(self._http_client, client=self)
        self.role = role.Role(self._http_client, client=self)
        self.pdu = pdud.Pdu(self._http_client, client=self)
        self.k8scluster = k8scluster.K8scluster(self._http_client, client=self)
        self.vca = vca.VCA(self._http_client, client=self)
        self.repo = repo.Repo(self._http_client, client=self)
        self.osmrepo = osmrepo.OSMRepo(self._http_client, client=self)
        self.package_tool = package_tool.PackageTool(client=self)
        self.subscription = subscription.Subscription(self._http_client, client=self)
        """
        self.vca = vca.Vca(http_client, client=self, **kwargs)
        self.utils = utils.Utils(http_client, **kwargs)
        """

    def get_token(self):
        self._logger.debug("")
        if self._token is None:
            postfields_dict = {
                "username": self._user,
                "password": self._password,
                "project_id": self._project,
            }
            if self._project_domain_name:
                postfields_dict["project_domain_name"] = self._project_domain_name
            if self._user_domain_name:
                postfields_dict["user_domain_name"] = self._user_domain_name
            http_code, resp = self._http_client.post_cmd(
                endpoint=self._auth_endpoint,
                postfields_dict=postfields_dict,
                skip_query_admin=True,
            )
            #            if http_code not in (200, 201, 202, 204):
            #                message ='Authentication error: not possible to get auth token\nresp:\n{}'.format(resp)
            #                raise ClientException(message)

            token = json.loads(resp) if resp else None
            self._token = token["id"]

            if self._token is not None:
                self._headers["Authorization"] = "Bearer {}".format(self._token)
                http_header = [
                    "{}: {}".format(key, val)
                    for (key, val) in list(self._headers.items())
                ]
                self._http_client.set_http_header(http_header)

    def get_version(self):
        _, resp = self._http_client.get2_cmd(endpoint="/version", skip_query_admin=True)
        # print(http_code, resp)
        try:
            resp = json.loads(resp)
            version = resp.get("version")
            date = resp.get("date")
        except ValueError:
            version = resp.split()[2]
            date = resp.split()[4]
        return "{} {}".format(version, date)

    def set_default_params(self, **kwargs):
        host = kwargs.pop("host", None)
        if host is not None:
            self._host = host
        port = kwargs.pop("port", None)
        if port is not None:
            self._so_port = port
        self._http_client.set_query_admin(**kwargs)
