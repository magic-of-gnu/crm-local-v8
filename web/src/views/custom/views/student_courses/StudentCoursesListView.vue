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
        inputName="first_name"
        inputTitle="First Name"
        v-model:inputModel="searchValues[0]"
        inputPlaceholder="Enter First Name"
        opName="first_name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[0]"
        :opOptions="firstNameSearchOpOptions"
      />

      <StringFilter 
        inputName="last_name"
        inputTitle="Last Name"
        v-model:inputModel="searchValues[1]"
        inputPlaceholder="Enter Last Name"
        opName="last_name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[1]"
        :opOptions="lastNameSearchOpOptions"
      />

      <StringFilter 
        inputName="course_name"
        inputTitle="Course Name"
        v-model:inputModel="searchValues[2]"
        inputPlaceholder="Enter Course Name"
        opName="courseName_op"
        opTitle="Operation"
        v-model:opModel="searchOps[2]"
        :opOptions="courseNameSearchOpOptions"
      />

      <NumFilter 
        dataType="number"
        inputName="payment_amount"
        inputTitle="Payment Amount"
        v-model:inputModel="searchValues[3]"
        inputPlaceholder="Enter Payment Amount"
        opName="paymentAmount_op"
        opTitle="Operation"
        v-model:opModel="searchOps[3]"
        :opOptions="paymentAmountSearchOpOptions"
      />

      <StringFilter 
        inputName="description"
        inputTitle="description"
        v-model:inputModel="searchValues[4]"
        inputPlaceholder="Enter Description Filter"
        opName="description_op"
        opTitle="Operation"
        v-model:opModel="searchOps[4]"
        :opOptions="descriptionSearchOpOptions"
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
                :to="{ name: 'student_courses_create_one' }"
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
                <CTableHeaderCell scope="col">Student Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Course Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Payment Amount</CTableHeaderCell>
                <CTableHeaderCell scope="col">Description</CTableHeaderCell>
                <CTableHeaderCell scope="col">Created At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Updated At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
                <CTableDataCell>{{ showUUID(item.id) }}</CTableDataCell>
                <CTableDataCell>{{ item.student_first_name }} {{ item.student_last_name }}</CTableDataCell>
                <CTableDataCell>{{ item.course_name }}</CTableDataCell>                
                <CTableDataCell>{{ item.payment_amount }}</CTableDataCell>
                <CTableDataCell>{{ item.description }}</CTableDataCell>
                <CTableDataCell>{{ utcTimeToUpdateTime(item.created_at) }}</CTableDataCell>
                <CTableDataCell>{{ utcTimeToUpdateTime(item.updated_at) }}</CTableDataCell>
              </CTableRow>
            </CTableBody>
          </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script setup>
import m from '@/views/custom/hooks/studentCourses/methods.js'
import { ref, onBeforeMount, reactive } from 'vue'
import { sortList, stringOperators, numberOperators, showUUID } from '@/utils/utils.js'
import { utcTimeToUpdateTime } from '@/utils/utils.js'
import StringFilter from '@/components/custom_filters/StringFilter.vue'
import NumFilter from '@/components/custom_filters/NumFilter.vue'
import ToastComponent from '@/components/ToastComponent.vue'
import { populateToastsFromResponse } from '@/utils/toasts';
import ConfirmModal from '@/components/ModalConfirmComponent.vue'
import { applyFilter } from '@/utils/filter.js'

import { stringOperatorsAsOptions, numberOperatorsAsOptions } from '@/utils/utils.js'
import { applyAllFilters } from '@/utils/filter.js'

const showModal = ref(false)
const toasts = ref([])

const firstNameSearchOpOptions = stringOperatorsAsOptions()
const lastNameSearchOpOptions = stringOperatorsAsOptions()
const courseNameSearchOpOptions = stringOperatorsAsOptions()
const paymentAmountSearchOpOptions = numberOperatorsAsOptions()
const descriptionSearchOpOptions = stringOperatorsAsOptions()

const originalData = ref(null)
const columnNames = ["student_first_name", "student_last_name", "course_name", "payment_amount", "description", "created_at", "updated_at"]
const selectedItems = ref({})

const searchValues = ref(new Array(columnNames.length).fill(null))
const searchOps = ref([
  firstNameSearchOpOptions[0].value,
  lastNameSearchOpOptions[0].value,
  courseNameSearchOpOptions[0].value,
  paymentAmountSearchOpOptions[0].value,
  descriptionSearchOpOptions[0].value,
])


const itemsList = ref('')
let sort_asc = reactive({
  first_name: true,
  last_name: true,
  username: true,
  info: true,
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

// function localApplyFilter() {
//   applyFilter(originalData.value, columnNames, searchValues.value, searchOps.value, itemsList)
// }

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
