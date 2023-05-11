from http import HTTPStatus
import pytest

class TestAttendanceValues:
  def test_get_all_response_status(self, attendance_values_get_all):
    assert attendance_values_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, attendance_values_create_one):
    assert attendance_values_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, attendance_values_get_one_by_id):
    assert attendance_values_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, attendance_values_create_one_get_data, attendance_values_payload_create_one):
    for key, value in attendance_values_payload_create_one.items():
      assert value == attendance_values_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, attendance_values_delete_one_by_id):
    assert attendance_values_delete_one_by_id.status_code == HTTPStatus.OK
