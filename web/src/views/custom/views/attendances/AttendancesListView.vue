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
        opName="courseName_op"
        opTitle="Operation"
        v-model:opModel="searchOps[1]"
        :opOptions="courseNameSearchOpOptions"
      />

      <StringFilter 
        inputName="lecture_date"
        inputTitle="Lecture Date"
        v-model:inputModel="searchValues[2]"
        inputPlaceholder="Enter Lecture Date"
        opName="lectureDate_op"
        opTitle="Operation"
        v-model:opModel="searchOps[2]"
        :opOptions="lectureDateSearchOpOptions"
      />

      <StringFilter 
        inputName="employee_first_name"
        inputTitle="Employee First Name"
        v-model:inputModel="searchValues[3]"
        inputPlaceholder="Enter Employee First Name"
        opName="employeeFirstName_op"
        opTitle="Operation"
        v-model:opModel="searchOps[3]"
        :opOptions="employeeFirstNameSearchOpOptions"
      />

      <StringFilter 
        inputName="employee_last_name"
        inputTitle="Employee Last Name"
        v-model:inputModel="searchValues[4]"
        inputPlaceholder="Enter Employee Last Name"
        opName="employeeLastName_op"
        opTitle="Operation"
        v-model:opModel="searchOps[4]"
        :opOptions="employeeLastNameSearchOpOptions"
      />

      <StringFilter 
        inputName="employee_username"
        inputTitle="Employee User Name"
        v-model:inputModel="searchValues[5]"
        inputPlaceholder="Enter Employee User Name"
        opName="employeeUsername_op"
        opTitle="Operation"
        v-model:opModel="searchOps[5]"
        :opOptions="employeeUsernameSearchOpOptions"
      />

      <StringFilter 
        inputName="student_first_name"
        inputTitle="Student First Name"
        v-model:inputModel="searchValues[6]"
        inputPlaceholder="Enter Student First Name"
        opName="studentFirstName_op"
        opTitle="Operation"
        v-model:opModel="searchOps[6]"
        :opOptions="studentFirstNameSearchOpOptions"
      />

      <StringFilter 
        inputName="student_last_name"
        inputTitle="Student Last Name"
        v-model:inputModel="searchValues[7]"
        inputPlaceholder="Enter Student Last Name"
        opName="studentLastName_op"
        opTitle="Operation"
        v-model:opModel="searchOps[7]"
        :opOptions="studentLastNameSearchOpOptions"
      />

      <StringFilter 
        inputName="student_username"
        inputTitle="Student User Name"
        v-model:inputModel="searchValues[8]"
        inputPlaceholder="Enter Student User Name"
        opName="studentUsername_op"
        opTitle="Operation"
        v-model:opModel="searchOps[8]"
        :opOptions="studentUsernameSearchOpOptions"
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
                <CTableHeaderCell scope="col">
                  Room Name
                  <CIcon
                    v-if="sort_asc.room_name"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'room_name', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'room_name', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Course Name
                  <CIcon
                    v-if="sort_asc.course_name"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'course_name', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'course_name', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Lecture Date
                  <CIcon
                    v-if="sort_asc.lecture_date"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'lecture_date', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'lecture_date', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Employee Name
                  <CIcon
                    v-if="sort_asc.employee_name"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'employee_name', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'employee_name', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Employee Username
                  <CIcon
                    v-if="sort_asc.employee_username"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'employee_username', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'employee_username', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Student Name
                  <CIcon
                    v-if="sort_asc.student_name"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'student_name', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'student_name', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Student Username
                  <CIcon
                    v-if="sort_asc.student_username"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'student_username', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'student_username', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Attendance Name
                  <CIcon
                    v-if="sort_asc.attendance_name"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'attendance_name', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'attendance_name', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Payment Status Name
                  <CIcon
                    v-if="sort_asc.payment_status_name"
                    class="mx-2"
                    icon="cilSortAlphaDown"
                    size="sm" 
                    @click="sortList(itemsList, 'payment_status_name', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortAlphaUp"
                    size="sm" 
                    @click="sortList(itemsList, 'payment_status_name', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Created At
                  <CIcon
                    v-if="sort_asc.created_at"
                    class="mx-2"
                    icon="cilSortNumericUp"
                    size="sm" 
                    @click="sortList(itemsList, 'created_at', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortNumericDown"
                    size="sm" 
                    @click="sortList(itemsList, 'created_at', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">
                  Updated At
                  <CIcon
                    v-if="sort_asc.updated_at"
                    class="mx-2"
                    icon="cilSortNumericUp"
                    size="sm" 
                    @click="sortList(itemsList, 'updated_at', sort_asc)"
                  />
                  <CIcon
                    v-else
                    class="mx-2"
                    icon="cilSortNumericDown"
                    size="sm" 
                    @click="sortList(itemsList, 'updated_at', sort_asc)"
                  />
                </CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
                <CTableDataCell>{{ showUUID(item.id) }}</CTableDataCell>
                <CTableDataCell>{{ item.room_name }}</CTableDataCell>
                <CTableDataCell>{{ item.course_name }}</CTableDataCell>
                <CTableDataCell>{{ epochToDatetime(item.lecture_date) }}</CTableDataCell>
                <CTableDataCell>{{ item.employee_first_name }} {{ item.employee_last_name}}</CTableDataCell>
                <CTableDataCell>{{ item.employee_username }}</CTableDataCell>
                <CTableDataCell>{{ item.student_first_name }} {{ item.student_last_name }}</CTableDataCell>
                <CTableDataCell>{{ item.student_username }}</CTableDataCell>
                <CTableDataCell>{{ item.attendance_name }}</CTableDataCell>
                <CTableDataCell>{{ item.payment_status_name }}</CTableDataCell>
                <CTableDataCell>{{ utcTimeToUpdateTime(item.attendance_created_at) }}</CTableDataCell>
                <CTableDataCell>{{ utcTimeToUpdateTime(item.attendance_updated_at) }}</CTableDataCell>
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
import { ref, onBeforeMount, reactive } from 'vue'
import { sortList, stringOperators, numberOperators, showUUID } from '@/utils/utils.js'
import { utcTimeToUpdateTime } from '@/utils/utils.js'
import StringFilter from '@/components/custom_filters/StringFilter.vue'
import NumFilter from '@/components/custom_filters/NumFilter.vue'
import ToastComponent from '@/components/ToastComponent.vue'
import { populateToastsFromResponse } from '@/utils/toasts';
import ConfirmModal from '@/components/ModalConfirmComponent.vue'
import { epochToDatetime } from '@/utils/utils.js'
import SortIconComponent from '@/components/SortIconComponent.vue'

import { stringOperatorsAsOptions, numberOperatorsAsOptions } from '@/utils/utils.js'
import { applyAllFilters } from '@/utils/filter.js'

const showModal = ref(false)
const toasts = ref([])

const roomNameSearchOpOptions = stringOperatorsAsOptions()
const courseNameSearchOpOptions = stringOperatorsAsOptions()
const lectureDateSearchOpOptions = stringOperatorsAsOptions()
const employeeFirstNameSearchOpOptions = stringOperatorsAsOptions()
const employeeLastNameSearchOpOptions = stringOperatorsAsOptions()
const employeeUsernameSearchOpOptions = stringOperatorsAsOptions()
const studentFirstNameSearchOpOptions = stringOperatorsAsOptions()
const studentLastNameSearchOpOptions = stringOperatorsAsOptions()
const studentUsernameSearchOpOptions = stringOperatorsAsOptions()

const originalData = ref(null)
const columnNames = [
  "room_name",
  "course_name",
  "lecture_date",
  "employee_first_name",
  "employee_last_name",
  "employee_username",
  "student_first_name",
  "student_last_name",
  "student_username",
]
const selectedItems = ref({})

const searchValues = ref(new Array(columnNames.length).fill(null))
const searchOps = ref([
  roomNameSearchOpOptions[0].value,
  courseNameSearchOpOptions[0].value,
  lectureDateSearchOpOptions[0].value,
  employeeFirstNameSearchOpOptions[0].value,
  employeeLastNameSearchOpOptions[0].value,
  employeeUsernameSearchOpOptions[0].value,
  studentFirstNameSearchOpOptions[0].value,
  studentLastNameSearchOpOptions[0].value,
  studentUsernameSearchOpOptions[0].value,
])


const itemsList = ref('')
let sort_asc = reactive({
  room_name: true,
  course_name: true,
  lecture_date: true,
  employee_first_name: true,
  employee_last_name: true,
  employee_username: true,
  student_first_name: true,
  student_last_name: true,
  student_username: true,
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
  applyAllFilters(originalData.value, columnNames, searchValues.value, searchOps.value, itemsList)
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
