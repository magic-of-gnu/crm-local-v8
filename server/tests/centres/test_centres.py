from http import HTTPStatus
import pytest

class TestCentres:
  def test_get_all_response_status(self, centres_get_all):
    assert centres_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, centres_create_one):
    assert centres_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, centres_get_one_by_id):
    assert centres_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, centres_create_one_get_data, centres_payload_create_one):
    for key, value in centres_payload_create_one.items():
      assert value == centres_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, centres_delete_one_by_id):
    assert centres_delete_one_by_id.status_code == HTTPStatus.OK
