from urllib.parse import urljoin

import pytest
import requests


@pytest.fixture(scope="module")
def attendance_values_routes(env_routes):
  return env_routes["attendance_values_routes"]

@pytest.fixture(scope="module")
def url_attendance_values_create_one(attendance_values_routes, base_url):
  return urljoin(base_url, attendance_values_routes["base"])

@pytest.fixture
def url_attendance_values_get_all(attendance_values_routes, base_url):
  return urljoin(base_url, attendance_values_routes["base"])

@pytest.fixture(scope="module")
def url_attendance_values_get_one_by_id(attendance_values_routes, base_url):
  return urljoin(base_url, attendance_values_routes["baseID"])

@pytest.fixture(scope="module")
def url_attendance_values_delete_one_by_id(attendance_values_routes, base_url):
  return urljoin(base_url, attendance_values_routes["baseID"])

@pytest.fixture(scope="module")
def url_attendance_values_patch_one_by_id(attendance_values_routes, base_url):
  return urljoin(base_url, attendance_values_routes["baseID"])

# all payloads for attendance_values
@pytest.fixture(scope="module")
def attendance_values_payloads(env_payloads):
  return env_payloads["attendance_values"]

# payloads for create_one attendance_value
@pytest.fixture(scope="module")
def attendance_values_payload_create_one(attendance_values_payloads):
  return attendance_values_payloads["create_one"]

# GET attendance_values all
@pytest.fixture
def attendance_values_get_all(url_attendance_values_get_all):
  with requests.Session() as sess:
    response = sess.get(url_attendance_values_get_all)

  return response

# POST attendance_values create_one
@pytest.fixture(scope="module")
def attendance_values_create_one(url_attendance_values_create_one, attendance_values_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_attendance_values_create_one, json=attendance_values_payload_create_one)

  return response

@pytest.fixture(scope="module")
def attendance_values_create_one_get_data(attendance_values_create_one):
  data = attendance_values_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def attendance_values_create_one_get_id(attendance_values_create_one_get_data):
  return attendance_values_create_one_get_data["id"]

# GET one attendance_value by id
@pytest.fixture(scope="module")
def attendance_values_get_one_by_id(url_attendance_values_get_one_by_id, attendance_values_create_one_get_id):
  attendance_value_id = attendance_values_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_attendance_values_get_one_by_id.format(id=attendance_value_id))

  return response

# DELETE attendance_values delete_one_by_id 
@pytest.fixture
def attendance_values_delete_one_by_id(attendance_values_create_one_get_id, url_attendance_values_delete_one_by_id):
  item_id = attendance_values_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_attendance_values_delete_one_by_id.format(id=item_id))

  return response