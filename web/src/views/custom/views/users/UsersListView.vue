
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
                :to="{ name: 'users_create_one' }"
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
                <CTableHeaderCell scope="col">First Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Last Name</CTableHeaderCell>
                <CTableHeaderCell scope="col">Username</CTableHeaderCell>
                <CTableHeaderCell scope="col">Last Login</CTableHeaderCell>
                <CTableHeaderCell scope="col">Is Admin</CTableHeaderCell>
                <CTableHeaderCell scope="col">User Type</CTableHeaderCell>
                <CTableHeaderCell scope="col">Created At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Updated At</CTableHeaderCell>
                <CTableHeaderCell scope="col">Operations</CTableHeaderCell>
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
                <CTableDataCell>{{ item.id }}</CTableDataCell>
                <CTableDataCell>{{ item.first_name }}</CTableDataCell>
                <CTableDataCell>{{ item.last_name }}</CTableDataCell>
                <CTableDataCell>{{ item.username }}</CTableDataCell>
                <CTableDataCell>{{ item.last_login }}</CTableDataCell>
                <CTableDataCell>{{ item.is_admin }}</CTableDataCell>
                <CTableDataCell>{{ item.user_type }}</CTableDataCell>
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
import m from '@/views/custom/hooks/users/methods.js'
import { ref, onMounted } from 'vue'

const itemsList = ref('')

onMounted(() => {
  m.getAllList().then((d) => {
    itemsList.value = d.data.data
  })
})
</script>
