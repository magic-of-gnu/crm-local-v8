from http import HTTPStatus
import pytest

class TestCourses:
  def test_get_all_response_status(self, courses_get_all):
    assert courses_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, courses_create_one):
    assert courses_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, courses_get_one_by_id):
    assert courses_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, courses_create_one_get_data, courses_payload_create_one):
    for key, value in courses_payload_create_one.items():
      assert value == courses_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, courses_delete_one_by_id):
    assert courses_delete_one_by_id.status_code == HTTPStatus.OK
