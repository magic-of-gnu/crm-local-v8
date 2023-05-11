from urllib.parse import urljoin

import pytest
import requests

from centres.conftest import *


@pytest.fixture(scope="module")
def students_routes(env_routes):
  return env_routes["students_routes"]

@pytest.fixture(scope="module")
def url_students_create_one(students_routes, base_url):
  return urljoin(base_url, students_routes["base"])

@pytest.fixture
def url_students_get_all(students_routes, base_url):
  return urljoin(base_url, students_routes["base"])

@pytest.fixture(scope="module")
def url_students_get_one_by_id(students_routes, base_url):
  return urljoin(base_url, students_routes["baseID"])

@pytest.fixture(scope="module")
def url_students_delete_one_by_id(students_routes, base_url):
  return urljoin(base_url, students_routes["baseID"])

@pytest.fixture(scope="module")
def url_students_patch_one_by_id(students_routes, base_url):
  return urljoin(base_url, students_routes["baseID"])

# all payloads for students
@pytest.fixture(scope="module")
def students_payloads(env_payloads):
  return env_payloads["students"]

# payloads for create_one student
@pytest.fixture(scope="module")
def students_payload_create_one(students_payloads):
  return students_payloads["create_one"]

# GET students all
@pytest.fixture
def students_get_all(url_students_get_all):
  with requests.Session() as sess:
    response = sess.get(url_students_get_all)

  return response

# POST students create_one
@pytest.fixture(scope="module")
def students_create_one(url_students_create_one, students_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_students_create_one, json=students_payload_create_one)

  return response

@pytest.fixture(scope="module")
def students_create_one_get_data(students_create_one):
  data = students_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def students_create_one_get_id(students_create_one_get_data):
  return students_create_one_get_data["id"]

# GET one student by id
@pytest.fixture(scope="module")
def students_get_one_by_id(url_students_get_one_by_id, students_create_one_get_id):
  student_id = students_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_students_get_one_by_id.format(id=student_id))

  return response

# DELETE students delete_one_by_id 
@pytest.fixture
def students_delete_one_by_id(students_create_one_get_id, url_students_delete_one_by_id):
  item_id = students_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_students_delete_one_by_id.format(id=item_id))

  return response