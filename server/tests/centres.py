import sys

import utils

from http import HTTPStatus
from urllib.parse import urljoin

from utils import get_json_config
from utils import req_headers
from utils import test_getAll, test_getOneByID, test_createOne, test_patchOneByID, test_deleteOneByID

def run_tests():
  fname_config = "envs.json"

  cfg = get_json_config(fname_config)

  payload_createOne = {
    "name": "some centre name",
    "description": "some centre description too"
  }

  base_route_name = 'centres_routes'

  url_getAll = urljoin(cfg["baseURL"], cfg[base_route_name]["base"])
  url_getOneByID = urljoin(cfg["baseURL"], cfg[base_route_name]["baseID"])
  url_createOne = urljoin(cfg["baseURL"], cfg[base_route_name]["base"])
  url_deleteOneByID = urljoin(cfg["baseURL"], cfg[base_route_name]["baseID"])
  url_patchOneByID = urljoin(cfg["baseURL"], cfg[base_route_name]["baseID"])

  data_getAll = test_getAll(url_getAll, req_headers=req_headers)
  data_createOne = test_createOne(url_createOne, payload_createOne, req_headers=req_headers)

  params_id = {
    "id": data_createOne["data"]["id"]
  }

  payload_patchOne1 = {
    "id": data_createOne["data"]["id"],
    "name": "some new name",
  }

  payload_patchOne2 = {
    "id": data_createOne["data"]["id"],
    "name": "some new name2",
  }

  data_getOneByID = test_getOneByID(url_getOneByID, params_id, req_headers=req_headers)

  assert data_getOneByID["name"] == payload_createOne["name"], f"error on {url_patchOneByID}"
  assert data_getOneByID["description"] == payload_createOne["description"], f"error on {url_patchOneByID}"

  data_patchOne1 = test_patchOneByID(url_patchOneByID, params_id, payload_patchOne1, req_headers=req_headers)
  data_getOneByID1 = test_getOneByID(url_getOneByID, params_id, req_headers=req_headers)
  data_patchOne2 = test_patchOneByID(url_patchOneByID, params_id, payload_patchOne2, req_headers=req_headers)
  data_getOneByID2 = test_getOneByID(url_getOneByID, params_id, req_headers=req_headers)

  utils.check_dict_eql(payload_patchOne1, data_getOneByID1, ["name"])
  utils.check_dict_eql(payload_patchOne2, data_getOneByID2, ["name"])

  assert data_getOneByID1["name"] == payload_patchOne1["name"], f"error on {url_patchOneByID}"
  assert data_getOneByID2["name"] == payload_patchOne2["name"], f"error on {url_patchOneByID}"

  data_deleteOneByID = test_deleteOneByID(url_deleteOneByID, params_id, req_headers=req_headers)


if __name__ == "__main__":

  run_tests()