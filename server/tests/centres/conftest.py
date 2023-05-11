from urllib.parse import urljoin

import pytest
import requests


@pytest.fixture(scope="module")
def centres_routes(env_routes):
  return env_routes["centres_routes"]

@pytest.fixture(scope="module")
def url_centres_create_one(centres_routes, base_url):
  return urljoin(base_url, centres_routes["base"])

@pytest.fixture(scope="module")
def url_centres_get_all(centres_routes, base_url):
  return urljoin(base_url, centres_routes["base"])

@pytest.fixture(scope="module")
def url_centres_get_one_by_id(centres_routes, base_url):
  return urljoin(base_url, centres_routes["baseID"])

@pytest.fixture
def url_centres_delete_one_by_id(centres_routes, base_url):
  return urljoin(base_url, centres_routes["baseID"])

@pytest.fixture(scope="session")
def url_centres_patch_one_by_id(centres_routes, base_url):
  return urljoin(base_url, centres_routes["baseID"])



# all payloads for centres
@pytest.fixture(scope="module")
def centres_payloads(env_payloads):
  return env_payloads["centres"]

# payload for create_one centre
@pytest.fixture(scope="module")
def centres_payload_create_one(centres_payloads):
  return centres_payloads["create_one"]

# payload for patch1
@pytest.fixture(scope="module")
def centres_payload_patch1(centres_payloads):
  return centres_payloads["patch1"]


# GET all centres
@pytest.fixture
def centres_get_all(url_centres_get_all):
  with requests.Session() as sess:
    response = sess.get(url_centres_get_all)

  return response


# POST centres create_one
@pytest.fixture(scope="module")
def centres_create_one(url_centres_create_one, centres_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_centres_create_one, json=centres_payload_create_one)

  return response


@pytest.fixture(scope="module")
def centres_create_one_get_data(centres_create_one):
  data = centres_create_one.json()["data"]
  return data


@pytest.fixture(scope="module")
def centres_create_one_get_id(centres_create_one_get_data):
  centre_id = centres_create_one_get_data["id"]
  return centre_id


# GET one centre by id
@pytest.fixture
def centres_get_one_by_id(url_centres_get_one_by_id, centres_create_one_get_id):
  centre_id = centres_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_centres_get_one_by_id.format(id=centre_id))

  return response


# DELETE centres delete_one_by_id 
@pytest.fixture
def centres_delete_one_by_id(centres_create_one_get_id, url_centres_delete_one_by_id):
  centre_id = centres_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_centres_delete_one_by_id.format(id=centre_id))

  return response
