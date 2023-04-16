
<template>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="course_id">Course Name</CFormLabel>
        <CFormSelect
          v-model="selected_course_id"
          aria-label="Default select example"
          :options="coursesOptions"
          name="course_id"
          aria-placeholder="Select"
          required
        >
        </CFormSelect>
      </div>

      <CListGroup v-if="lecturesCalendar">
        <router-link
          v-for="item in lecturesCalendar"
          :key="item.id"
          component="CButton"
          :to="{ name: 'attendances_create_one_id', params: { id: item.id }}"
          > 
          <div class="row">
            <div class="col">
              <p>Date: {{ getDate(item.date) }}</p>
              <p>Time: {{ getTime(item.date) }}</p>
              <p>Duration: {{ item.duration }} minutes</p>
            </div>
            <div class="col">
              <p>Teacher Name: {{ item.employee.first_name }} {{ item.employee.last_name }} </p>
              <p>Class Room Name: {{ item.room.name }} </p>
            </div>
          </div>
        </router-link>
      </CListGroup>
      <CListGroup v-else>
        <CListGroupItem>No Lectures on Such Course</CListGroupItem>
      </CListGroup>

    </CForm>
  </CRow>
</template>

<script setup>
import m from '@/views/custom/hooks/attendances/methods.js'
import coursesMethods from '@/views/custom/hooks/courses/methods.js'
import attendanceValuesMethods from '@/views/custom/hooks/attendanceValues/methods.js'
import lecturesCalendarMethods from '@/views/custom/hooks/lecturesCalendar/methods.js'
import { ref, onMounted, computed, watch } from 'vue'
import { epochToDatetime, utcTimeToDate, utcTimeToLocaleTime } from '@/utils/utils.js'

import VueDatePicker from '@vuepic/vue-datepicker';

const selected_course_id = ref(null)
const description = ref("")

const lecturesCalendar = ref(null)
const coursesOptions = ref([
  {
    label: "Select Course",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
  }
])

const attendanceValuesOptions = ref([
  {
    label: "Select Attendance Value",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
  }
])

const lecturesCalendarOptions = ref([
  {
    label: "Select Lecture",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
  }
])

function getDate(time_string) {
  return utcTimeToDate(time_string)
}

function getTime(time_string) {
  return utcTimeToLocaleTime(time_string)
}

function run_this_method(event) {
  m.postCreateOne({
    lectures_calendar_id: selected_lectures_calendar_id.value,
    student_id: selected_student_id.value,
    attendance_value_id: selected_attendance_value_id.value,
    description: description.value,
  })
}

watch (selected_course_id, (newValue, oldValue) => {
  if (!newValue) {
    return
  }

  lecturesCalendar.value = []
  lecturesCalendarMethods.getManyByCourseID({
    course_id: newValue
  }).then((d) => {
    if (!d.data) {
      return 
    }
    return d.data.forEach((item, ii) => {
      lecturesCalendar.value.push(item)
    })
  })
  console.log("lectureCalendar: ", lecturesCalendar.value)
})


onMounted(() => {
  coursesMethods.getAllList()
    .then( (d) => {
      console.log("d: ", d)
      for (const item of d.data.data) {
        coursesOptions.value.push({
          label: item.name,
          value: item.id
      })
    }
  })

  attendanceValuesMethods.getAllList()
    .then( (d) => {
      for (const item of d.data.data) {
        attendanceValuesOptions.value.push({
          label: item.name,
          value: item.id
      })
    }
  })
})

</script>



