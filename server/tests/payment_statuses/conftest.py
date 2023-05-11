from urllib.parse import urljoin

import pytest
import requests


@pytest.fixture(scope="module")
def payment_statuses_routes(env_routes):
  return env_routes["payment_statuses_routes"]

@pytest.fixture(scope="module")
def url_payment_statuses_create_one(payment_statuses_routes, base_url):
  return urljoin(base_url, payment_statuses_routes["base"])

@pytest.fixture
def url_payment_statuses_get_all(payment_statuses_routes, base_url):
  return urljoin(base_url, payment_statuses_routes["base"])

@pytest.fixture(scope="module")
def url_payment_statuses_get_one_by_id(payment_statuses_routes, base_url):
  return urljoin(base_url, payment_statuses_routes["baseID"])

@pytest.fixture(scope="module")
def url_payment_statuses_delete_one_by_id(payment_statuses_routes, base_url):
  return urljoin(base_url, payment_statuses_routes["baseID"])

@pytest.fixture(scope="module")
def url_payment_statuses_patch_one_by_id(payment_statuses_routes, base_url):
  return urljoin(base_url, payment_statuses_routes["baseID"])

# all payloads for payment_statuses
@pytest.fixture(scope="module")
def payment_statuses_payloads(env_payloads):
  return env_payloads["payment_statuses"]

# payloads for create_one payment_statuse
@pytest.fixture(scope="module")
def payment_statuses_payload_create_one(payment_statuses_payloads):
  return payment_statuses_payloads["create_one"]

# GET payment_statuses all
@pytest.fixture
def payment_statuses_get_all(url_payment_statuses_get_all):
  with requests.Session() as sess:
    response = sess.get(url_payment_statuses_get_all)

  return response

# POST payment_statuses create_one
@pytest.fixture(scope="module")
def payment_statuses_create_one(url_payment_statuses_create_one, payment_statuses_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_payment_statuses_create_one, json=payment_statuses_payload_create_one)

  return response

@pytest.fixture(scope="module")
def payment_statuses_create_one_get_data(payment_statuses_create_one):
  data = payment_statuses_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def payment_statuses_create_one_get_id(payment_statuses_create_one_get_data):
  return payment_statuses_create_one_get_data["id"]

# GET one payment_statuse by id
@pytest.fixture(scope="module")
def payment_statuses_get_one_by_id(url_payment_statuses_get_one_by_id, payment_statuses_create_one_get_id):
  payment_statuse_id = payment_statuses_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_payment_statuses_get_one_by_id.format(id=payment_statuse_id))

  return response

# DELETE payment_statuses delete_one_by_id 
@pytest.fixture
def payment_statuses_delete_one_by_id(payment_statuses_create_one_get_id, url_payment_statuses_delete_one_by_id):
  item_id = payment_statuses_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_payment_statuses_delete_one_by_id.format(id=item_id))

  return response