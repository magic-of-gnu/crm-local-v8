<template>
<ToastComponent 
  v-if="toasts"
  :toasts="toasts"
/>

  <CRow>
    <CCol xs="12">
    <CForm action="some-post-function" method="POST">

      <CRow>
        <CCol class="mb-3">
          <CFormLabel for="student_id">Student Name</CFormLabel>
          <CFormSelect
            v-model="selectedStudentID"
            aria-label="Default select example"
            :options="optionsStudents"
            name="students_id"
            aria-placeholder="Select"
          >
          </CFormSelect>
        </CCol>
      </CRow>

      <CRow>
        <CCol class="mb-3">
          <CFormLabel for="course_id">Course Name</CFormLabel>
          <CFormSelect
            v-model="selectedCourseID"
            aria-label="Default select example"
            :options="optionsCourses"
            name="course_id"
            aria-placeholder="Select"
          >
          </CFormSelect>
        </CCol>
      </CRow>

      <CRow>
        <CCol class="mb-3">
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
        </CCol>
      </CRow>

      <CRow>
        <CCol class="mb-3">
          <CFormLabel for="price">Price</CFormLabel>
          <CFormInput
            v-model.number="price"
            type="number"
            id="price"
            name="price"
            placeholder="Enter Price"
          />
        </CCol>
      </CRow>

      <CRow>
        <CCol class="mb-3">
          <CFormLabel for="payment_status_id">Payment Status Name</CFormLabel>
          <CFormSelect
            v-model="selectedPaymentStatusID"
            aria-label="Default select example"
            :options="optionsPaymentStatuses"
            name="payment_status_id"
            aria-placeholder="Select"
          >
          </CFormSelect>
        </CCol>
      </CRow>

      <CRow>
        <CCol class="mb-3">
          <CFormLabel for="lectures_number">Number of Lectures</CFormLabel>
          <CFormInput
            v-model.number="lectures_number"
            type="number"
            id="lectures_number"
            name="lectures_number"
            placeholder="Enter Lectures Number"
          />
        </CCol>
      </CRow>

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
    </CCol>
  </CRow>
</template>

<script setup>
import m from '@/views/custom/hooks/invoices/methods.js'
import { ref, onBeforeMount, watch } from 'vue'
import ToastComponent from '@/components/ToastComponent.vue'
import { populateToastsFromResponse } from '@/utils/toasts';
import coursesMethods from '@/views/custom/hooks/courses/methods.js'
import studentsMethods from '@/views/custom/hooks/students/methods.js'
import studentCoursesMethods from '@/views/custom/hooks/studentCourses/methods.js'
import paymentStatusesMethods from '@/views/custom/hooks/paymentStatuses/methods.js'
import { jsEpochToGoEpoch } from '@/utils/utils.js'

import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'

const toasts = ref([])

const optionsCourses = ref([{
  label: "Select Course",
  value: "",
  disabled: true,
  hidden: true,
  selected: true
}])

const optionsStudents = ref([{
  label: "Select Student",
  value: "",
  disabled: true,
  hidden: true,
  selected: true
}])

const optionsPaymentStatuses = ref([{
  label: "Select Payment Status",
  value: "",
  disabled: true,
  hidden: true,
  selected: true
}])

const selectedCourseID = ref(null)
const selectedStudentID = ref(null)
const start_date = ref(null)
const price = ref(null)
const selectedPaymentStatusID = ref(null)
const lectures_number = ref(null)

async function run_this_method(event) {
  const start_date_epoch = jsEpochToGoEpoch(start_date.value)
  const request_data = {
    course_id: selectedCourseID.value,
    student_id: selectedStudentID.value,
    start_date: start_date_epoch,
    price: price.value,
    payment_status_id: selectedPaymentStatusID.value,
    lectures_number: lectures_number.value,
  }

  console.log("request_data: ", request_data)

  const response = await m.postCreateOne(request_data)
  populateToastsFromResponse(response, toasts)

  // selectedCourseID.value = null
  // selectedPaymentStatusID.value = null
  // start_date.value = null
  // price.value = null
  // selectedPaymentStatusID.value = null
  // lectures_number.value = null

}

  // const studentCoursesParams = {
  //   student_id: 
  // }
  // const response_courses = await studentCoursesMethods.getManyByCustomID()

  // for (const item of response_courses.data.data.data) {
  //   optionsCourses.value.push({
  //     label: item.name,
  //     value: item.id
  //   })
  // }

watch(selectedStudentID, async (newValue, oldValue) => {
  const response_student_courses = await studentCoursesMethods.getManyByCustomID({
    student_id: newValue,
  })

  console.log("data: ", response_student_courses.data.data)
  optionsCourses.value = [{
    label: "Select Course",
    value: "",
    disabled: true,
    hidden: true,
    selected: true
  }]

  for (const item of response_student_courses.data.data) {
    optionsCourses.value.push({
      label: item.course.name,
      value: item.course.id
    })
  }
})


onBeforeMount(async () => {

  const response_students = await studentsMethods.getAllList()

  for (const item of response_students.data.data.data) {
    optionsStudents.value.push({
      label: item.first_name + " " + item.last_name,
      value: item.id
    })
  }

  const response_payment_statuses = await paymentStatusesMethods.getAllList()

  for (const item of response_payment_statuses.data.data.data) {
    optionsPaymentStatuses.value.push({
      label: item.name,
      value: item.id
    })
  }

})

</script>
