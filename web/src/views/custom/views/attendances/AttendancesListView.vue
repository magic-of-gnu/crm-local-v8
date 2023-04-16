
<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-4">
        <CCardHeader>
          <CRow class="justify-content-between">
            <CCol>
              <strong>Vue Progress</strong> <small>Basic example</small>
            </CCol>
            <CCol class="align-self-end">
              <router-link
                :to="{ name: 'attendances_create_by_course_name' }"
                component="CButton"
                color="primary"
                disabled
                >Create By Course Name
              </router-link>
            </CCol>
          </CRow>
        </CCardHeader>
        <CCardBody>
          <CTable hover>
            <CTableHead>
              <CTableRow>
                <CTableHeaderCell scope="col">ID</CTableHeaderCell>
                <CTableHeaderCell scope="col">Room Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Course Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Employee Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Employee Username</CTableHeaderCell>
                <CTableHeaderCell scope="col">Student Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Student Username</CTableHeaderCell>
                <CTableHeaderCell scope="col">Attendance Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Description</CTableHeaderCell>
                <CTableHeaderCell scope="col">Created At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Updated At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
                <CTableDataCell>{ item.id }</CTableDataCell>
                <CTableDataCell>{ item.room_name }</CTableDataCell>
                <CTableDataCell>{ item.course_name }</CTableDataCell>
                <CTableDataCell>{ item.employee_first_name } { item.employee_last_name}</CTableDataCell>
                <CTableDataCell>{ item.employee_username }</CTableDataCell>
                <CTableDataCell>{ item.student_first_name } { item.student_last_name }</CTableDataCell>
                <CTableDataCell>{ item.student_username }</CTableDataCell>
                <CTableDataCell>{ item.attendance_name }</CTableDataCell>
                <CTableDataCell>{ item.description }</CTableDataCell>
                <CTableDataCell>{ item.created_at }</CTableDataCell>
                <CTableDataCell>{ item.updated_at }</CTableDataCell>
              </CTableRow>
            </CTableBody>
          </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script setup>
import m from '@/views/custom/hooks/attendances/methods.js'
import { ref, onMounted } from 'vue'

const itemsList = ref('')

onMounted(() => {
  m.getAllList().then((d) => {
    itemsList.value = d.data.data
  })
})
</script>
