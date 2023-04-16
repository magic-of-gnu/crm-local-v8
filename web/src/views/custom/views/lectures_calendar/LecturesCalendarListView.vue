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
                :to="{ name: 'lectures_calendar_create_one' }"
                component="CButton"
                color="primary"
                disabled
                >Create One
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
                <CTableHeaderCell scope="col">Date</CTableHeaderCell>
                <CTableHeaderCell scope="col">Duration</CTableHeaderCell>
                <CTableHeaderCell scope="col">Created At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Updated At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
                <CTableDataCell>{{ item.id }}</CTableDataCell>
                <CTableDataCell>{{ item.room_name }}</CTableDataCell>                
                <CTableDataCell>{{ item.course_name }}</CTableDataCell>                
                <CTableDataCell>{{ item.employee_first_name }} {{ item.employee_last_name }}</CTableDataCell>
                <CTableDataCell>{{ epochToDatetime(item.lecture_date) }}</CTableDataCell>
                <CTableDataCell>{{ item.lecture_duration }}</CTableDataCell>
                <CTableDataCell>{{ item.created_at }}</CTableDataCell>
                <CTableDataCell>{{ item.updated_at }}</CTableDataCell>
              </CTableRow>
            </CTableBody>
          </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script setup>
import m from '@/views/custom/hooks/lecturesCalendar/methods.js'
// import moment from 'moment'
import { ref, onMounted, computed } from 'vue'
import { epochToDatetime } from '@/utils/utils.js'

const itemsList = ref('')

onMounted(() => {
  m.getAllList().then((d) => {
    itemsList.value = d.data.data
  })
})
</script>
