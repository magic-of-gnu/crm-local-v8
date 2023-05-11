from urllib.parse import urljoin

import pytest
import requests

@pytest.fixture(scope="module")
def users_routes(env_routes):
  return env_routes["users_routes"]

@pytest.fixture(scope="module")
def url_users_create_one(users_routes, base_url):
  return urljoin(base_url, users_routes["base"])

@pytest.fixture
def url_users_get_all(users_routes, base_url):
  return urljoin(base_url, users_routes["base"])

@pytest.fixture(scope="module")
def url_users_get_one_by_id(users_routes, base_url):
  return urljoin(base_url, users_routes["baseID"])

@pytest.fixture(scope="module")
def url_users_delete_one_by_id(users_routes, base_url):
  return urljoin(base_url, users_routes["baseID"])

@pytest.fixture(scope="module")
def url_users_patch_one_by_id(users_routes, base_url):
  return urljoin(base_url, users_routes["baseID"])

# all payloads for users
@pytest.fixture(scope="module")
def users_payloads(env_payloads):
  return env_payloads["users"]

# payloads for create_one user
@pytest.fixture(scope="module")
def users_payload_create_one(users_payloads):
  return users_payloads["create_one"]

# payloads for create_one user
@pytest.fixture(scope="module")
def users_payload_create_one_check_data(users_payloads):
  return users_payloads["create_one_check"]

# GET users all
@pytest.fixture
def users_get_all(url_users_get_all):
  with requests.Session() as sess:
    response = sess.get(url_users_get_all)

  return response

# POST users create_one
@pytest.fixture(scope="module")
def users_create_one(url_users_create_one, users_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_users_create_one, json=users_payload_create_one)

  return response

@pytest.fixture(scope="module")
def users_create_one_get_data(users_create_one):
  data = users_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def users_create_one_get_id(users_create_one_get_data):
  return users_create_one_get_data["id"]

# GET one user by id
@pytest.fixture(scope="module")
def users_get_one_by_id(url_users_get_one_by_id, users_create_one_get_id):
  user_id = users_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_users_get_one_by_id.format(id=user_id))

  return response

# DELETE users delete_one_by_id 
@pytest.fixture
def users_delete_one_by_id(users_create_one_get_id, url_users_delete_one_by_id):
  item_id = users_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_users_delete_one_by_id.format(id=item_id))

  return response