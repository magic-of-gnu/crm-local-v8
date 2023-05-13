from http import HTTPStatus
import pytest

class TestLecturesCalendar:
  def test_get_all_response_status(self, lectures_calendars_get_all):
    assert lectures_calendars_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, lectures_calendars_create_many):
    assert lectures_calendars_create_many.status_code == HTTPStatus.CREATED

  # def test_get_one_by_id_response_status(self, lectures_calendars_get_one_by_id):
  #   assert lectures_calendars_get_one_by_id.status_code == HTTPStatus.OK

  # def test_create_one_data_correctness(self, lectures_calendars_create_one_get_data, lectures_calendars_payload_create_one):
  #   for key, value in lectures_calendars_payload_create_one.items():
  #     assert value == lectures_calendars_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, lectures_calendars_delete_many_one_by_one_by_id):
    for response in lectures_calendars_delete_many_one_by_one_by_id:
      assert response.status_code == HTTPStatus.OK
