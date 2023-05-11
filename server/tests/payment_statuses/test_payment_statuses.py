from http import HTTPStatus
import pytest

class TestPaymentStatuses:
  def test_get_all_response_status(self, payment_statuses_get_all):
    assert payment_statuses_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, payment_statuses_create_one):
    assert payment_statuses_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, payment_statuses_get_one_by_id):
    assert payment_statuses_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, payment_statuses_create_one_get_data, payment_statuses_payload_create_one):
    for key, value in payment_statuses_payload_create_one.items():
      assert value == payment_statuses_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, payment_statuses_delete_one_by_id):
    assert payment_statuses_delete_one_by_id.status_code == HTTPStatus.OK
