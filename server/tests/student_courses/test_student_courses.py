from http import HTTPStatus
import pytest

class TestStudentCourses:
  def test_get_all_response_status(self, student_courses_get_all):
    assert student_courses_get_all.status_code == HTTPStatus.OK

  def test_create_one_response_status(self, student_courses_create_one):
    assert student_courses_create_one.status_code == HTTPStatus.CREATED

  def test_get_one_by_id_response_status(self, student_courses_get_one_by_id):
    assert student_courses_get_one_by_id.status_code == HTTPStatus.OK

  def test_create_one_data_correctness(self, student_courses_create_one_get_data, student_courses_payload_create_one):
    for key, value in student_courses_payload_create_one.items():
      assert value == student_courses_create_one_get_data[key]

  def test_delete_one_by_id_response_status(self, student_courses_delete_one_by_id):
    assert student_courses_delete_one_by_id.status_code == HTTPStatus.OK
