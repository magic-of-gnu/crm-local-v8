<template>
<ToastComponent
v-if="toasts"
:toasts="toasts"
/>
  <CRow>
    <CForm action="some-post-function" method="POST">

      <div class="mb-3">
        <CFormLabel for="centre_id">Room Name</CFormLabel>
        <CFormSelect
          v-model="selected_room_id"
          aria-label="Default select example"
          :options="roomsOptions"
          name="room_id"
          aria-placeholder="Select"
          required
        >
        </CFormSelect>
      </div>
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
      <div class="mb-3">
        <CFormLabel for="employee_id">Employee Name</CFormLabel>
        <CFormSelect
          v-model="selected_employee_id"
          aria-label="Default select example"
          :options="employeesOptions"
          name="employee_id"
          aria-placeholder="Select"
          required
        >
        </CFormSelect>
      </div>

      <div class="mb-3">
        <CFormLabel for="attendance_value_id">Default Attendance Value</CFormLabel>
        <CFormSelect
          id="attendance_value_id"
          v-model="selected_attendance_value_id"
          aria-label="Default select example"
          :options="attendanceValuesOptions"
          name="attendance_value_id"
          aria-placeholder="Select"
          required
        >
        </CFormSelect>
      </div>

      <div class="mb-3">
        <CFormLabel for="start_date">Start Date</CFormLabel>
        <VueDatePicker
          name="start_date"
          v-model="start_date"
          auto-apply
          placeholder="Select Start Date"
          close-on-scroll
          :highlight-week-days="[0, 7]"
          required
          model-type="timestamp"
          week-start="1"
          :enable-time-picker="false"
        ></VueDatePicker>
      </div>
      <div class="mb-3">
        <CFormLabel for="end_date">End Date</CFormLabel>
        <VueDatePicker
          name="end_date"
          v-model="end_date"
          auto-apply
          placeholder="Select End Date"
          close-on-scroll
          :highlight-week-days="[0, 7]"
          required
          model-type="timestamp"
          week-start="1"
          :enable-time-picker="false"
        ></VueDatePicker>
      </div>

      <div class="row mb-3">
        <div class="col">

          <CButton
            @click.prevent="($event) => addNewDateAndTime($event)"
            component="input"
            type="button"
            color="primary"
            value="Add New Day"
          />
        </div>

        <div class="col">

          <CButton
            @click.prevent="($event) => deleteLastDateAndTime(ii)"
            component="input"
            type="button"
            color="primary"
            value="Delete Last Option"
          />
        </div>
        
      </div>

      <div
        class="mb-3"
        v-for="(item, ii) in dates_and_times"
        :key="ii"
      >

        <CInputGroup class="row mb-3">

          <div class="col mb-3">

            <CFormLabel for="day">Select Day</CFormLabel>
            <CFormSelect
              v-model="item.day"
              :options="daysOptions"
              name="day"
              required
            >
            </CFormSelect>
          </div>

          <div class="col mb-3">
            <CFormLabel for="start_time">Start Time</CFormLabel>
            <VueDatePicker
              v-model="item.start_time"
              :is-24="true"
              auto-apply
              time-picker
              placeholder="Select Start Time"
              close-on-scroll
              required
            ></VueDatePicker>
          </div>

          <div class="col mb-3">
            <CFormLabel for="duration">Duration</CFormLabel>
            <VueDatePicker
              v-model="item.duration"
              auto-apply
              time-picker
              placeholder="Select Duration"
              close-on-scroll
              required
            ></VueDatePicker>
          </div>

          <CButton
            class="col mb-3"
            @click.prevent="($event) => deleteThisDateAndTime($event)"
            component="input"
            type="button"
            color="primary"
            value="x"
            size="sm"
            variant="ghost"
          />

        </CInputGroup>
      </div>
      <CButton
        @click.prevent="($event) => run_this_method($event)"
        component="input"
        type="submit"
        color="primary"
        value="Submit"
      />
    </CForm>
  </CRow>
</template>

<script setup>
/* eslint-disable */

import m from '@/views/custom/hooks/lecturesCalendar/methods.js'
import employeesMethods from '@/views/custom/hooks/employees/methods.js'
import roomsMethods from '@/views/custom/hooks/rooms/methods.js'
import coursesMethods from '@/views/custom/hooks/courses/methods.js'
import attendanceValuesMethods from '@/views/custom/hooks/attendanceValues/methods.js'
import { ref, onBeforeMount } from 'vue'

import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import ToastComponent from '@/components/ToastComponent.vue'

const daysOptions = [
  {
    label: "Select Day",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
  },
  {
    label: "Monday",
    value: 1
  },
  {
    label: "Tuesday",
    value: 2
  },
  {
    label: "Wednesday",
    value: 3
  },
  {
    label: "Thursday",
    value: 4
  },
  {
    label: "Friday",
    value: 5
  },
  {
    label: "Satyrday",
    value: 6
  },
  {
    label: "Sunday",
    value: 7
  },
]

const dates_and_times = ref([])

const start_date = ref(new Date())
const end_date = ref(new Date())

const selected_course_id = ref(null)
const selected_room_id = ref(null)
const selected_employee_id = ref(null)
const selected_attendance_value_id = ref(null)

const employeesOptions = ref([{
  label: "Select Employee",
  value: "",
  disabled: true,
  hidden: true,
  selected: true
}])
const roomsOptions = ref([{
  label: "Select Room",
  value: "",
  disabled: true,
  hidden: true,
  selected: true
}])
const coursesOptions = ref([{
    label: "Select Course",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
}])
const attendanceValuesOptions = ref([
  {
    label: "Select Attendance Value",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
  }
])

async function run_this_method(event) {
  const start_date_epoch = Math.round(start_date.value / 1000) // milliseconds -> seconds
  const end_date_epoch = Math.round(end_date.value / 1000) // milliseconds -> seconds
  const times = []
  dates_and_times.value.forEach((item, ii) => {
    times.push({
      day: parseInt(item.day, 10),
      start_time: item.start_time.hours.toString().padStart(2, 0)+ item.start_time.minutes.toString().padStart(2, 0),
      duration: item.duration.hours * 60 + item.duration.minutes,
    })
  })
  const data = {
    room_id: selected_room_id.value,
    course_id: selected_course_id.value,
    employee_id: selected_employee_id.value,
    start_date: start_date_epoch,
    end_date: end_date_epoch,
    dates_and_times: times,
    default_attendance_value_id: selected_attendance_value_id.value,
  }

  const response = await m.postCreateOne(data)

  if (response.data.hasOwnProperty("toasts")) {
    response.data.toasts.forEach((item) => {
      toasts.value.push(item)
    });
  }

}

function addNewDateAndTime(event) {
  dates_and_times.value.push({
    day: null,
    start_time: null,
    duration: null
  })
}

function deleteLastDateAndTime() {
  dates_and_times.value.pop()
}

function deleteThisDateAndTime(ii) {
  dates_and_times.value.splice(ii, 1)
}

onBeforeMount(async () => {
  const response_av = await attendanceValuesMethods.getAllList()
  console.log("response_av: ", response_av.data)

  for (const item of response_av.data.data.data) {
    attendanceValuesOptions.value.push({
      label: item.name,
      value: item.id
    })
  }

  const response_rooms = await roomsMethods.getAllList()
  console.log(response_rooms)

  for (const item of response_rooms.data.data.data) {
    roomsOptions.value.push({
      label: item.name,
      value: item.id
    })
  }

  employeesMethods.getAllList()
    .then( (d) => {
      for (const item of d.data.data.data) {
        employeesOptions.value.push({
          label: item.first_name + " " + item.last_name,
          value: item.id
      })
    }
  })

  coursesMethods.getAllList()
    .then( (d) => {
      for (const item of d.data.data.data) {
        coursesOptions.value.push({
          label: item.name,
          value: item.id
      })
    }
  })

})

</script>
