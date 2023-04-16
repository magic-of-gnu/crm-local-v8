
<template>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="course_id">Course Name</CFormLabel>
        <CFormInput type="text" id="course_id" :placeholder="$route.params.id" disabed/>
        <CFormText component="span" id="exampleFormControlInputHelpInline">
            {{ $route.params.id }}
        </CFormText>
      </div>
      <div class="row mb-3">
        <div class="col">
          <CFormLabel for="date">Lecture Date</CFormLabel>
          <CFormInput type="text" id="date" :placeholder="lecture_date" disabed/>
        </div>

        <div class="col">
          <CFormLabel for="start_time">Start Time</CFormLabel>
          <CFormInput type="text" id="start_time" :placeholder="start_time" disabed/>
        </div>

        <div class="col">
          <CFormLabel for="duration">Duration</CFormLabel>
          <CFormInput type="text" id="duration" :placeholder="duration" disabed/>
        </div>
      </div>

      <div class="row mb-3">
        <div class="col">
          <CFormLabel for="room_name">Class Room</CFormLabel>
          <CFormInput type="text" id="room_name" :placeholder="room_name" disabed/>
        </div>
      </div>

      <div class="row mb-3">
        <div class="col">
          <CFormLabel for="employee_name">Employee Room</CFormLabel>
          <CFormInput type="text" id="employee_name" :placeholder="employee_name" disabed/>
        </div>
      </div>

      <div class="row">
        <div class="col">
          <CListGroup v-if="courseStudents">
            <CListGroupItem
              v-for="(item, ii) in courseStudents"
              :key="item.id"
              > 
              <div class="row">
                <div class="col">
                  <p>Student Name: {{ item.student.first_name }} {{ item.student.last_name }}</p>
                  <p>Student Username: {{ item.student.username }}</p>
                </div>
                <div class="col">
                  <div class="row">
                    <div
                      class="col"
                      v-for="vitem in attendanceValues"
                      :key="vitem.id"
                    >
                      <input
                        type="radio"
                        :id="vitem.id + ii"
                        :value="vitem.id"
                        v-model="courseStudents[ii].attendance_value_id"
                      />
                      <label :for="vitem.id + ii">{{ vitem.name }}</label>
                    </div>
                  </div>
                </div>
                <div class="col">
                  <CFormTextarea
                    id="description"
                    type="string"
                    name="description"
                    rows="2"
                    placeholder="Enter Comment Here"
                    v-model="courseStudents[ii].attendance_description"
                  ></CFormTextarea>
                </div>
              </div>
            </CListGroupItem>
          </CListGroup>
          <CListGroup v-else>
            <CListGroupItem>No Lectures on Such Course</CListGroupItem>
          </CListGroup>
        </div>
      </div>

      <div class="row my-3">
        <div class="col">
          <CButton
            @click.prevent="($event) => run_this_method($event)"
            component="input"
            type="submit"
            color="primary"
            value="Submit"
          />
        </div>
      </div>
    </CForm>
  </CRow>
</template>

<script setup>
import m from '@/views/custom/hooks/attendances/methods.js'
import lecturesCalendarMethods from '@/views/custom/hooks/lecturesCalendar/methods.js'
import studentCoursesMethods from '@/views/custom/hooks/studentCourses/methods.js'
import attendanceValuesMethods from '@/views/custom/hooks/attendanceValues/methods.js'
import { ref, onBeforeMount, watch } from 'vue'
import { useRoute } from 'vue-router'

import { epochToDatetime, utcTimeToDate, utcTimeToLocaleTime, utcTimeToTimeHoursMinutes } from '@/utils/utils.js'

const route = useRoute()

const currentLectureCalendar = ref(null)

const room_name = ref("")
const employee_name = ref("")

const attendanceValues = ref([])

const course_id = ref(null)
const courseStudents = ref([])

const lecture_date = ref("")
const start_time = ref("")
const duration = ref("")



function run_this_method(event) {
  let request_data = []

  console.log("courseStudents: ", courseStudents.value)
  courseStudents.value.forEach((item) => {
    request_data.push({
      lectures_calendar_id: route.params.id,
      student_id: item.student_id,
      attendance_value_id: item.attendance_value_id,
      description: item.attendance_description || ""
    })
  })

  m.postCreateMany(request_data)
}

function getDate(time_string) {
  return utcTimeToDate(time_string)
}

function getTimeHM(time_string) {
  return utcTimeToTimeHoursMinutes(time_string)
}

onBeforeMount(async () => {
  await lecturesCalendarMethods.getOneByID(route.params.id)
    .then( (d) => {
      currentLectureCalendar.value = d.data
      lecture_date.value = getDate(d.data.date)
      start_time.value = getTimeHM(d.data.date)
      duration.value = d.data.duration
      room_name.value = d.data.room.name
      course_id.value = d.data.course_id
      employee_name.value = d.data.employee.first_name + " " + d.data.employee.last_name + " (" + d.data.employee.username + ")"
    })

  attendanceValuesMethods.getAllList()
    .then( (d) => {
      attendanceValues.value = d.data.data
  })
  
  await studentCoursesMethods.getManyByCustomID({
    "course_id": course_id.value
  }).then( (d) => {
    courseStudents.value = d.data
  })


})

</script>



