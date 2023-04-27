import json
from http import HTTPStatus

import requests

req_headers = {
  "Content-type": "application/json",
  "Accept": "application/json"
}

def get_json_config(fname):
  with open(fname) as f:
    cfg = json.load(f)

  return cfg


def check_dict_eql(d1: dict, d2: dict, keys: [str]):
  for key in keys:
    assert d1.get(key) == d2.get(key), f"different values on key: {key}; d1.key: {d1.get(key)}; d2.key: {d2.get(key)}"

def check_dict_neql(d1: dict, d2: dict, keys: [str]):
  for key in keys:
    assert d1.get(key) != d2.get(key), f"identical values on key: {key}; d1.key: {d1.get(key)}; d2.key: {d2.get(key)}"



def test_getAll(url, req_headers, http_status=HTTPStatus.OK):

  with requests.Session() as sess:

    response = sess.get(url, headers=req_headers)

  if response.status_code != http_status:
    raise Exception(f"incorrect status code during get all {url}")

  return response.json()

def test_getOneByID(url, path_params, req_headers, http_status=HTTPStatus.OK):

  with requests.Session() as sess:

    response = sess.get(url.format(**path_params), headers=req_headers)

  if response.status_code != http_status:
    raise Exception(f"incorrect status code during get one {url}")

  return response.json()["data"]["data"]

def test_createOne(url, payload, req_headers, http_status=HTTPStatus.CREATED):

  with requests.Session() as sess:

    response = sess.post(url, headers=req_headers, json=payload)

  if response.status_code != http_status:
    raise Exception(f"incorrect status code during create one {url}")

  return response.json()

def test_patchOneByID(url, path_params, payload, req_headers, http_status=HTTPStatus.OK):

  with requests.Session() as sess:

    response = sess.patch(url.format(**path_params), headers=req_headers, json=payload)

  if response.status_code != http_status:
    raise Exception(f"incorrect status code during patch {url}")

  return response.json()

def test_deleteOneByID(url, path_params, req_headers, http_status=HTTPStatus.OK):

  with requests.Session() as sess:

    response = sess.delete(url.format(**path_params), headers=req_headers)

  if response.status_code != http_status:
    raise Exception(f"incorrect status code during delete {url}")

  return response.json()
