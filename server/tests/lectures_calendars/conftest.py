from urllib.parse import urljoin

import pytest
import requests

from courses.conftest import *
from rooms.conftest import *
from employees.conftest import *
from attendance_values.conftest import *

@pytest.fixture(scope="module")
def lectures_calendars_routes(env_routes):
  return env_routes["lectures_calendars_routes"]

@pytest.fixture(scope="module")
def url_lectures_calendars_create_many(lectures_calendars_routes, base_url):
  return urljoin(base_url, lectures_calendars_routes["base"])

@pytest.fixture
def url_lectures_calendars_get_all(lectures_calendars_routes, base_url):
  return urljoin(base_url, lectures_calendars_routes["base"])

@pytest.fixture(scope="module")
def url_lectures_calendars_get_one_by_id(lectures_calendars_routes, base_url):
  return urljoin(base_url, lectures_calendars_routes["baseID"])

@pytest.fixture(scope="module")
def url_lectures_calendars_delete_one_by_id(lectures_calendars_routes, base_url):
  return urljoin(base_url, lectures_calendars_routes["baseID"])

@pytest.fixture(scope="module")
def url_lectures_calendars_patch_one_by_id(lectures_calendars_routes, base_url):
  return urljoin(base_url, lectures_calendars_routes["baseID"])

# all payloads for lectures_calendars
@pytest.fixture(scope="module")
def lectures_calendars_payloads(env_payloads):
  return env_payloads["lectures_calendars"]

# payloads for create_one lectures_calendar
@pytest.fixture(scope="module")
# def lectures_calendars_payload_create_one(lectures_calendars_payloads, courses_create_one_get_id, students_create_one_get_id):
def lectures_calendars_payload_create_many(
  lectures_calendars_payloads,
  rooms_create_one_get_id,
  courses_create_one_get_id,
  employees_create_one_get_id,
  attendance_value_id_class_did_not_start
  ):

  room_id = rooms_create_one_get_id
  course_id = courses_create_one_get_id
  employee_id = employees_create_one_get_id
  attendance_value_id = attendance_value_id_class_did_not_start

  # retrieve centre id and add it to the payload
  payload = lectures_calendars_payloads["create_many"]
  payload["room_id"] = room_id
  payload["course_id"] = course_id
  payload["employee_id"] = employee_id
  payload["default_attendance_value_id"] = attendance_value_id_class_did_not_start

  return  payload

# GET lectures_calendars all
@pytest.fixture
def lectures_calendars_get_all(url_lectures_calendars_get_all):
  with requests.Session() as sess:
    response = sess.get(url_lectures_calendars_get_all)

  return response

# POST lectures_calendars create_one
@pytest.fixture(scope="module")
def lectures_calendars_create_many(url_lectures_calendars_create_many, lectures_calendars_payload_create_many):
  with requests.Session() as sess:
    response = sess.post(url_lectures_calendars_create_many, json=lectures_calendars_payload_create_many)

  return response

@pytest.fixture(scope="module")
def lectures_calendars_create_many_get_data(lectures_calendars_create_many):
  data = lectures_calendars_create_many.json()["data"]
  return data

@pytest.fixture(scope="module")
def lectures_calendars_create_many_get_id(lectures_calendars_create_many_get_data):
  lectures_calendars_id = [item["id"] for item in lectures_calendars_create_many_get_data]
  return lectures_calendars_id

@pytest.fixture(scope="module")
def lectures_calendars_create_many_get_one_id(lectures_calendars_create_many_get_id):
  return lectures_calendars_create_many_get_id[0]

# GET one lectures_calendar by id
@pytest.fixture(scope="module")
def lectures_calendars_get_one_by_id(url_lectures_calendars_get_one_by_id, lectures_calendars_create_many_get_one_id):
  lectures_calendar_id = lectures_calendars_create_many_get_one_id
  with requests.Session() as sess:
    response = sess.get(url_lectures_calendars_get_one_by_id.format(id=lectures_calendar_id))

  return response

# DELETE lectures_calendars delete_one_by_id 
@pytest.fixture
def lectures_calendars_delete_many_one_by_one_by_id(lectures_calendars_create_many_get_id, url_lectures_calendars_delete_one_by_id):
  responses = []
  with requests.Session() as sess:
    for item in lectures_calendars_create_many_get_id:
      responses.append(sess.delete(url_lectures_calendars_delete_one_by_id.format(id=item)))

  return responses