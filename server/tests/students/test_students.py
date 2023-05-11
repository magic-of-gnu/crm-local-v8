from http import HTTPStatus
import pytest

class TestStudents:
  def test_get_all_response_status(self, students_get_all):
    assert students_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, students_create_one):
    assert students_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, students_get_one_by_id):
    assert students_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, students_create_one_get_data, students_payload_create_one):
    for key, value in students_payload_create_one.items():
      assert value == students_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, students_delete_one_by_id):
    assert students_delete_one_by_id.status_code == HTTPStatus.OK
