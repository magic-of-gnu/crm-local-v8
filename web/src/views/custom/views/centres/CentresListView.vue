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
                :to="{ name: 'centre_create_one' }"
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
                <CTableHeaderCell scope="col">Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Info</CTableHeaderCell>
                <CTableHeaderCell scope="col">Created At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Updated At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in centresList" :key="item.id">
                <CTableDataCell>{{ item.id }}</CTableDataCell>
                <CTableDataCell>{{ item.name }}</CTableDataCell>
                <CTableDataCell>{{ item.description }}</CTableDataCell>
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

<script>
import m from '@/views/custom/hooks/centres/methods.js'
import { ref } from 'vue'

export default {
  name: 'CentresListView',
  setup() {
    const centresList = ref('')

    m.getCentresList().then((d) => {
      centresList.value = d.data.data
      console.log('centresList: ', centresList)
      console.log('d.data: ', d.data.data)
    })

    return {
      centresList,
    }
  },
}
</script>
