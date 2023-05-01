
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
import { ref, onBeforeMount, computed, watch } from 'vue'
import { epochToDatetime, utcTimeToDate, utcTimeToLocaleTime } from '@/utils/utils.js'

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

watch (selected_course_id, async (newValue, oldValue) => {
  if (!newValue) {
    return
  }

  lecturesCalendar.value = []
  const response_lc = await lecturesCalendarMethods.getManyByCourseID({
    course_id: newValue
  })
  
  if (!response_lc.data.data) {
    return
  }
  response_lc.data.data.forEach((item, ii) => {
      lecturesCalendar.value.push(item)
    })
})


onBeforeMount(async () => {
  const response_courses = await coursesMethods.getAllList()
  for (const item of response_courses.data.data.data) {
    coursesOptions.value.push({
      label: item.name,
      value: item.id
    })
  }

  const response_av = await attendanceValuesMethods.getAllList()
  for (const item of response_av.data.data.data) {
    attendanceValuesOptions.value.push({
      label: item.name,
      value: item.id
    })
  }
})

</script>



