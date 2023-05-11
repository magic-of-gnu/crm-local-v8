from urllib.parse import urljoin

import pytest
import requests

from centres.conftest import *


@pytest.fixture(scope="module")
def courses_routes(env_routes):
  return env_routes["courses_routes"]

@pytest.fixture(scope="module")
def url_courses_create_one(courses_routes, base_url):
  return urljoin(base_url, courses_routes["base"])

@pytest.fixture
def url_courses_get_all(courses_routes, base_url):
  return urljoin(base_url, courses_routes["base"])

@pytest.fixture(scope="module")
def url_courses_get_one_by_id(courses_routes, base_url):
  return urljoin(base_url, courses_routes["baseID"])

@pytest.fixture(scope="module")
def url_courses_delete_one_by_id(courses_routes, base_url):
  return urljoin(base_url, courses_routes["baseID"])

@pytest.fixture(scope="module")
def url_courses_patch_one_by_id(courses_routes, base_url):
  return urljoin(base_url, courses_routes["baseID"])

# all payloads for courses
@pytest.fixture(scope="module")
def courses_payloads(env_payloads):
  return env_payloads["courses"]

# payloads for create_one course
@pytest.fixture(scope="module")
def courses_payload_create_one(courses_payloads):
  return courses_payloads["create_one"]

# GET courses all
@pytest.fixture
def courses_get_all(url_courses_get_all):
  with requests.Session() as sess:
    response = sess.get(url_courses_get_all)

  return response

# POST courses create_one
@pytest.fixture(scope="module")
def courses_create_one(url_courses_create_one, courses_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_courses_create_one, json=courses_payload_create_one)

  return response

@pytest.fixture(scope="module")
def courses_create_one_get_data(courses_create_one):
  data = courses_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def courses_create_one_get_id(courses_create_one_get_data):
  return courses_create_one_get_data["id"]

# GET one course by id
@pytest.fixture(scope="module")
def courses_get_one_by_id(url_courses_get_one_by_id, courses_create_one_get_id):
  course_id = courses_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_courses_get_one_by_id.format(id=course_id))

  return response

# DELETE courses delete_one_by_id 
@pytest.fixture
def courses_delete_one_by_id(courses_create_one_get_id, url_courses_delete_one_by_id):
  item_id = courses_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_courses_delete_one_by_id.format(id=item_id))

  return response