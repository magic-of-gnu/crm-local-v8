<template>
  <CRow>
    <CForm action="some-post-function" method="POST">

      <div class="mb-3">
        <CFormLabel for="centre_id">Student Name</CFormLabel>
        <CFormSelect
          v-model="selected_student_id"
          aria-label="Default select example"
          :options="studentsOptions"
          name="student_id"
          aria-placeholder="Select"
        >
        </CFormSelect>
      </div>
      <div class="mb-3">
        <CFormLabel for="centre_id">Course Name</CFormLabel>
        <CFormSelect
          v-model="selected_course_id"
          aria-label="Default select example"
          :options="coursesOptions"
          name="course_id"
          aria-placeholder="Select"
        >
        </CFormSelect>
      </div>
      <div class="mb-3">
        <CFormLabel for="payment_amount">Payment Amount</CFormLabel>
        <CFormInput
          v-model.number="payment_amount"
          type="number"
          id="payment_amount"
          name="payment_amount"
          placeholder="Enter Course Price for Student"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="description">Description</CFormLabel>
        <CFormInput
          v-model.trim="description"
          type="text"
          id="description"
          name="description"
          placeholder="Enter Description"
        />
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
import m from '@/views/custom/hooks/studentCourses/methods.js'
import studentsMethods from '@/views/custom/hooks/students/methods.js'
import coursesMethods from '@/views/custom/hooks/courses/methods.js'
import { ref, onMounted } from 'vue'

const payment_amount = ref(null)
const description = ref("")

const selected_student_id = ref(null)
const selected_course_id = ref(null)

const studentsOptions = ref([{}])
const coursesOptions = ref([{}])

function run_this_method(event) {
  m.postCreateOne({
    student_id: selected_student_id.value,
    course_id: selected_course_id.value,
    payment_amount: payment_amount.value,
    description: description.value,
  })
}

onMounted(() => {
  studentsMethods.getAllList()
    .then( (d) => {
      for (const item of d.data.data) {
        studentsOptions.value.push({
          label: item.first_name + " " + item.last_name,
          value: item.id
      })
    }
  })

  coursesMethods.getAllList()
    .then( (d) => {
      for (const item of d.data.data) {
        coursesOptions.value.push({
          label: item.name,
          value: item.id
      })
    }
  })

})

</script>
