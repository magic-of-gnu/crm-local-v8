from http import HTTPStatus
import pytest

class TestUsers:
  def test_get_all_response_status(self, users_get_all):
    assert users_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, users_create_one):
    assert users_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, users_get_one_by_id):
    assert users_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, users_create_one_get_data, users_payload_create_one_check_data):
    print("data: ", users_create_one_get_data)
    print("check: ", users_payload_create_one_check_data)
    for key, value in users_payload_create_one_check_data.items():
      assert value == users_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, users_delete_one_by_id):
    assert users_delete_one_by_id.status_code == HTTPStatus.OK
