from http import HTTPStatus
import pytest

class TestEmployees:
  def test_get_all_response_status(self, employees_get_all):
    assert employees_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, employees_create_one):
    assert employees_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, employees_get_one_by_id):
    assert employees_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, employees_create_one_get_data, employees_payload_create_one):
    for key, value in employees_payload_create_one.items():
      assert value == employees_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, employees_delete_one_by_id):
    assert employees_delete_one_by_id.status_code == HTTPStatus.OK
