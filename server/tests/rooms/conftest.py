from urllib.parse import urljoin

import pytest
import requests

from centres.conftest import *


@pytest.fixture(scope="module")
def rooms_routes(env_routes):
  return env_routes["rooms_routes"]

@pytest.fixture(scope="module")
def url_rooms_create_one(rooms_routes, base_url):
  return urljoin(base_url, rooms_routes["base"])

@pytest.fixture
def url_rooms_get_all(rooms_routes, base_url):
  return urljoin(base_url, rooms_routes["base"])

@pytest.fixture(scope="module")
def url_rooms_get_one_by_id(rooms_routes, base_url):
  return urljoin(base_url, rooms_routes["baseID"])

@pytest.fixture(scope="module")
def url_rooms_delete_one_by_id(rooms_routes, base_url):
  return urljoin(base_url, rooms_routes["baseID"])

@pytest.fixture(scope="module")
def url_rooms_patch_one_by_id(rooms_routes, base_url):
  return urljoin(base_url, rooms_routes["baseID"])

# all payloads for rooms
@pytest.fixture(scope="module")
def rooms_payloads(env_payloads):
  return env_payloads["rooms"]

# payloads for create_one room
@pytest.fixture(scope="module")
def rooms_payload_create_one(rooms_payloads, centres_create_one_get_id):

  # retrieve centre id and add it to the payload
  centre_id = centres_create_one_get_id

  payload = rooms_payloads["create_one"]
  payload["centre_id"] = centre_id
  return  payload

# GET rooms all
@pytest.fixture
def rooms_get_all(url_rooms_get_all):
  with requests.Session() as sess:
    response = sess.get(url_rooms_get_all)

  return response

# POST rooms create_one
@pytest.fixture(scope="module")
def rooms_create_one(url_rooms_create_one, rooms_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_rooms_create_one, json=rooms_payload_create_one)

  return response

@pytest.fixture(scope="module")
def rooms_create_one_get_data(rooms_create_one):
  data = rooms_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def rooms_create_one_get_id(rooms_create_one_get_data):
  print("rooms_create_one_get_data", rooms_create_one_get_data)
  return rooms_create_one_get_data["id"]

# GET one room by id
@pytest.fixture(scope="module")
def rooms_get_one_by_id(url_rooms_get_one_by_id, rooms_create_one_get_id):
  room_id = rooms_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_rooms_get_one_by_id.format(id=room_id))
    print("response: ", response.json())

  return response

# DELETE rooms delete_one_by_id 
@pytest.fixture
def rooms_delete_one_by_id(rooms_create_one_get_id, url_rooms_delete_one_by_id):
  item_id = rooms_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_rooms_delete_one_by_id.format(id=item_id))

  return response