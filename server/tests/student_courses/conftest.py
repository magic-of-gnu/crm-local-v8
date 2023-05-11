from urllib.parse import urljoin

import pytest
import requests

import courses.conftest as courses_fixtures
import students.conftest as students_fixtures

from students.conftest import *
from courses.conftest import *

from courses.conftest import courses_get_one_by_id
from students.conftest import students_get_one_by_id

@pytest.fixture(scope="module")
def student_courses_routes(env_routes):
  return env_routes["student_courses_routes"]

@pytest.fixture(scope="module")
def url_student_courses_create_one(student_courses_routes, base_url):
  return urljoin(base_url, student_courses_routes["base"])

@pytest.fixture
def url_student_courses_get_all(student_courses_routes, base_url):
  return urljoin(base_url, student_courses_routes["base"])

@pytest.fixture(scope="module")
def url_student_courses_get_one_by_id(student_courses_routes, base_url):
  return urljoin(base_url, student_courses_routes["baseID"])

@pytest.fixture(scope="module")
def url_student_courses_delete_one_by_id(student_courses_routes, base_url):
  return urljoin(base_url, student_courses_routes["baseID"])

@pytest.fixture(scope="module")
def url_student_courses_patch_one_by_id(student_courses_routes, base_url):
  return urljoin(base_url, student_courses_routes["baseID"])

# all payloads for student_courses
@pytest.fixture(scope="module")
def student_courses_payloads(env_payloads):
  return env_payloads["student_courses"]

# payloads for create_one student_course
@pytest.fixture(scope="module")
def student_courses_payload_create_one(student_courses_payloads, courses_create_one_get_id, students_create_one_get_id):

  # retrieve centre id and add it to the payload
  course_id = courses_create_one_get_id
  student_id = students_create_one_get_id

  payload = student_courses_payloads["create_one"]
  payload["student_id"] = student_id
  payload["course_id"] = course_id
  return  payload

# GET student_courses all
@pytest.fixture
def student_courses_get_all(url_student_courses_get_all):
  with requests.Session() as sess:
    response = sess.get(url_student_courses_get_all)
    print(response.json())

  return response

# POST student_courses create_one
@pytest.fixture(scope="module")
def student_courses_create_one(url_student_courses_create_one, student_courses_payload_create_one):
  with requests.Session() as sess:
    response = sess.post(url_student_courses_create_one, json=student_courses_payload_create_one)

  return response

@pytest.fixture(scope="module")
def student_courses_create_one_get_data(student_courses_create_one):
  data = student_courses_create_one.json()["data"]
  return data

@pytest.fixture(scope="module")
def student_courses_create_one_get_id(student_courses_create_one_get_data):
  return student_courses_create_one_get_data["id"]

# GET one student_course by id
@pytest.fixture(scope="module")
def student_courses_get_one_by_id(url_student_courses_get_one_by_id, student_courses_create_one_get_id):
  student_course_id = student_courses_create_one_get_id
  with requests.Session() as sess:
    response = sess.get(url_student_courses_get_one_by_id.format(id=student_course_id))

  return response

# DELETE student_courses delete_one_by_id 
@pytest.fixture
def student_courses_delete_one_by_id(student_courses_create_one_get_id, url_student_courses_delete_one_by_id):
  item_id = student_courses_create_one_get_id
  with requests.Session() as sess:
    response = sess.delete(url_student_courses_delete_one_by_id.format(id=item_id))

  return response