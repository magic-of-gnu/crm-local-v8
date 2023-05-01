<template>
<ToastComponent 
  v-if="toasts"
  :toasts="toasts"
/>

<ConfirmModal 
  content="Please, Confirm that You Want to Delete"
  closeBtnText="Cancel"
  confirmBtnText="Confirm"
  :showModal="showModal"
  :confirmDeleteClick="confirmDeleteClick"
/>

  <CRow>
    <CCol :xs="12">

  <CRow>
    <CCol xs="12">
    <CForm @submit.prevent>

      <StringFilter 
        inputName="room_name"
        inputTitle="Room Name"
        v-model:inputModel="searchValues[0]"
        inputPlaceholder="Enter Room Name"
        opName="room_name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[0]"
        :opOptions="roomNameSearchOpOptions"
      />

      <StringFilter 
        inputName="course_name"
        inputTitle="Course Name"
        v-model:inputModel="searchValues[1]"
        inputPlaceholder="Enter Course Name"
        opName="course_name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[1]"
        :opOptions="courseNameSearchOpOptions"
      />

      <StringFilter 
        inputName="employee_first_name"
        inputTitle="Employee First Name"
        v-model:inputModel="searchValues[2]"
        inputPlaceholder="Enter Employee First Name"
        opName="employee_name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[2]"
        :opOptions="employeeFirstNameSearchOpOptions"
      />

      <StringFilter 
        inputName="employee_last_name"
        inputTitle="Employee Last Name"
        v-model:inputModel="searchValues[3]"
        inputPlaceholder="Enter Employee Last Name"
        opName="employee_last_name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[3]"
        :opOptions="employeeLastNameSearchOpOptions"
      />


      <CRow class="mb-3">
        <CCol>
          <CButton
            @click.prevent="localApplyFilter()"
            component="button"
            type="button"
            color="primary"
          >Search</CButton>
        </CCol>
      </CRow>

    </CForm>
    </CCol>
  </CRow>




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
                <CTableDataCell>{{ showUUID(item.id) }}</CTableDataCell>
                <CTableDataCell>{{ item.room_name }}</CTableDataCell>                
                <CTableDataCell>{{ item.course_name }}</CTableDataCell>                
                <CTableDataCell>{{ item.employee_first_name }} {{ item.employee_last_name }}</CTableDataCell>
                <CTableDataCell>{{ epochToDatetime(item.lecture_date) }}</CTableDataCell>
                <CTableDataCell>{{ item.lecture_duration }}</CTableDataCell>
                <CTableDataCell>{{ utcTimeToUpdateTime(item.created_at) }}</CTableDataCell>
                <CTableDataCell>{{ utcTimeToUpdateTime(item.updated_at) }}</CTableDataCell>
                <CTableDataCell>
                  <CIcon 
                    icon="cilPencil"
                  /> 
                  <CIcon 
                    @click="confirmDeleteOneByID(item.id)" 
                    icon="cilTrash"
                  />
                </CTableDataCell>
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
import { ref, onBeforeMount, computed, reactive } from 'vue'
import { epochToDatetime } from '@/utils/utils.js'
import { sortList, stringOperators, numberOperators, showUUID } from '@/utils/utils.js'
import { utcTimeToUpdateTime } from '@/utils/utils.js'
import StringFilter from '@/components/custom_filters/StringFilter.vue'
import ToastComponent from '@/components/ToastComponent.vue'
import { populateToastsFromResponse } from '@/utils/toasts';
import ConfirmModal from '@/components/ModalConfirmComponent.vue'
import { applyFilter } from '@/utils/filter.js'


const showModal = ref(false)
const toasts = ref([])

const roomNameSearchOpOptions = stringOperators().map((x) => {
  return {label: x, value: x}
})
const courseNameSearchOpOptions = stringOperators().map((x) => {
  return {label: x, value: x}
})
const employeeFirstNameSearchOpOptions = stringOperators().map((x) => {
  return {label: x, value: x}
})
const employeeLastNameSearchOpOptions = stringOperators().map((x) => {
  return {label: x, value: x}
})


const originalData = ref(null)
const columnNames = ["room_name", "course_name", "employee_first_name", "employee_last_name", "lecture_date", "lecture_duration"]
const selectedItems = ref({})

const searchValues = ref(new Array(columnNames.length).fill(null))
const searchOps = ref([
  roomNameSearchOpOptions[0].value,
  courseNameSearchOpOptions[0].value,
  employeeFirstNameSearchOpOptions[0].value,
  employeeLastNameSearchOpOptions[0].value,
])

const itemsList = ref('')
let sort_asc = reactive({
  name: true,
  description: true,
  created_at: true,
  updated_at: true,
  id: true,
})

function confirmDeleteOneByID(id) {
  showModal.value = true
  selectedItems.value[id] = true
}

async function confirmDeleteClick(btnValue) {
  if (btnValue === true) {
    let response = null
    for (const key in selectedItems.value) {
        response = await m.postDeleteOneByID(key)
      }
    populateToastsFromResponse(response, toasts)
    await reset()
  }
  selectedItems.value = {}
  showModal.value = false
}

function localApplyFilter() {
  applyFilter(originalData.value, columnNames, searchValues.value, searchOps.value, itemsList)
}

async function reset() {
  const response = await m.getAllList()
  originalData.value = response.data.data.data
  itemsList.value = originalData.value.slice()
}

onBeforeMount(async () => {
  await reset()
})

</script>
