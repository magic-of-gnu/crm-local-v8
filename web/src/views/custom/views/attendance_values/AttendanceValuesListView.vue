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

      <NumFilter 
        dataType="number"
        inputName="value"
        inputTitle="Value"
        v-model:inputModel="searchValues[0]"
        inputPlaceholder="Enter Value"
        opName="value_op"
        opTitle="Operation"
        v-model:opModel="searchOps[0]"
        :opOptions="valueSearchOpOptions"
      />

      <StringFilter 
        inputName="name"
        inputTitle="Name"
        v-model:inputModel="searchValues[0]"
        inputPlaceholder="Enter Name"
        opName="name_op"
        opTitle="Operation"
        v-model:opModel="searchOps[0]"
        :opOptions="nameSearchOpOptions"
      />

      <StringFilter 
        inputName="description"
        inputTitle="Description"
        v-model:inputModel="searchValues[1]"
        inputPlaceholder="Enter Description Filter"
        opName="description_op"
        opTitle="Operation"
        v-model:opModel="searchOps[1]"
        :opOptions="descriptionSearchOpOptions"
      />

      <CRow class="mb-3">
        <CCol>
          <CButton
            @click.prevent="localApplyFilter"
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
                :to="{ name: 'attendance_values_create_one' }"
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
                <CTableHeaderCell scope="col">Value</CTableHeaderCell>
                <CTableHeaderCell scope="col">Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Description</CTableHeaderCell>
                <CTableHeaderCell scope="col">Created At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Updated At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
                <CTableDataCell>{{ showUUID(item.id) }}</CTableDataCell>
                <CTableDataCell>{{ item.value }}</CTableDataCell>                
                <CTableDataCell>{{ item.name }}</CTableDataCell>
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
import m from '@/views/custom/hooks/attendanceValues/methods.js'
import { ref, onBeforeMount, reactive } from 'vue'
import { sortList, stringOperators, numberOperators, showUUID } from '@/utils/utils.js'
import { utcTimeToUpdateTime } from '@/utils/utils.js'
import StringFilter from '@/components/custom_filters/StringFilter.vue'
import NumFilter from '@/components/custom_filters/StringFilter.vue'
import ToastComponent from '@/components/ToastComponent.vue'
import { populateToastsFromResponse } from '@/utils/toasts';
import ConfirmModal from '@/components/ModalConfirmComponent.vue'
import { applyFilter } from '@/utils/filter.js'

const showModal = ref(false)
const toasts = ref([])

const valueSearchOpOptions = numberOperators().map((x) => {
  return {label: x, value: x}
})
const nameSearchOpOptions = stringOperators().map((x) => {
  return {label: x, value: x}
})
const descriptionSearchOpOptions = stringOperators().map((x) => {
  return {label: x, value: x}
})
const createdAtSearchOpOptions = numberOperators().map((x) => {
  return {label: x, value: x}
})
const updatedAtSearchOpOptions = numberOperators().map((x) => {
  return {label: x, value: x}
})
const originalData = ref(null)
const columnNames = ["value", "name", "description", "created_at", "updated_at"]
const selectedItems = ref({})

const searchValues = ref(new Array(columnNames.length).fill(null))
const searchOps = ref([
  valueSearchOpOptions[0].value,
  nameSearchOpOptions[0].value,
  descriptionSearchOpOptions[0].value,
  createdAtSearchOpOptions[0].value,
  updatedAtSearchOpOptions[0].value,
])

const itemsList = ref('')
let sort_asc = reactive({
  value: true,
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
