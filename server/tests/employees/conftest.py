from urllib.parse import urljoin

import pytest
import requests

@pytest.fixture(scope="module")
def employees_routes(env_routes):
  return env_routes["employees_routes"]

@pytest.fixture(scope="module")
def url_employees_create_one(employees_routes, base_url):
  return urljoin(base_url, employees_routes["base"])

@pytest.fixture
def url_employees_get_all(employees_routes, base_url):
  return urljoin(base_url, employees_routes["base"])

@pytest.fixture(scope="module")
def url_employees_get_one_by_id(employees_routes, base_url):
  return urljoin(base_url, employees_routes["baseID"])

@pytest.fixture(scope="module")
def url_employees_delete_one_by_id(employees_routes, base_url):
  return urljoin(base_url, employees_routes["baseID"])

@pytest.fixture(scope="module")
def url_employees_patch_one_by_id(employees_routes, base_url):
  return urljoin(base_url, employees_routes["baseID"])

# all payloads for employees
@pytest.fixture(scope="module")
def employees_payloads(env_payloads):
  return env_payloads["employees"]

# payloads for create_one employee
@pytest.fixture(scope="module")
def employees_payload_create_one(employees_payloads):
  return employees_payloads["create_one"]

# GET employees all
@pytest.fixture
def employees_get_all(url_employees_get_all):
  with requests.Session() as sess:
    response = sess.get(url_employees_get_all)

  return response

# POST employees create_one
@pytest.fixture(scope="module")
def employees_create_one(url_employees_create_one, employees_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_employees_create_one, json=employees_payload_create_one)

  return response

@pytest.fixture(scope="module")
def employees_create_one_get_data(employees_create_one):
  data = employees_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def employees_create_one_get_id(employees_create_one_get_data):
  return employees_create_one_get_data["id"]

# GET one employee by id
@pytest.fixture(scope="module")
def employees_get_one_by_id(url_employees_get_one_by_id, employees_create_one_get_id):
  employee_id = employees_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_employees_get_one_by_id.format(id=employee_id))

  return response

# DELETE employees delete_one_by_id 
@pytest.fixture
def employees_delete_one_by_id(employees_create_one_get_id, url_employees_delete_one_by_id):
  item_id = employees_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_employees_delete_one_by_id.format(id=item_id))

  return response