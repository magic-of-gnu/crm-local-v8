from http import HTTPStatus
import pytest

class TestRooms:
  def test_get_all_response_status(self, rooms_get_all):
    assert rooms_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, rooms_create_one):
    assert rooms_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, rooms_get_one_by_id):
    assert rooms_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, rooms_create_one_get_data, rooms_payload_create_one):
    for key, value in rooms_payload_create_one.items():
      assert value == rooms_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, rooms_delete_one_by_id):
    assert rooms_delete_one_by_id.status_code == HTTPStatus.OK
