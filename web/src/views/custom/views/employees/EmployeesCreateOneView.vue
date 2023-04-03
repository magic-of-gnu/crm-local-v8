<template>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="first_name">First Name</CFormLabel>
        <CFormInput
          v-model.trim="first_name"
          type="text"
          id="first_name"
          name="first_name"
          placeholder="Enter First Name"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="last_name">Last Name</CFormLabel>
        <CFormInput
          v-model.trim="last_name"
          type="text"
          id="last_name"
          name="last_name"
          placeholder="Enter Last Name"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="username">Username</CFormLabel>
        <CFormInput
          v-model.trim="username"
          type="text"
          id="username"
          name="username"
          placeholder="Enter Username"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="info">Employee Information</CFormLabel>
        <CFormTextarea
          v-model.trim="info"
          id="info"
          type="string"
          name="dinfo"
          placeholder="Enter Room Information Here"
        ></CFormTextarea>
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
/* eslint-disable */

import m from '@/views/custom/hooks/employees/methods.js'
import { ref } from 'vue'

let first_name = ref("")
let last_name = ref("")
let username = ref("")
let info = ref("")

function run_this_method(event) {
  m.postEmployeesCreateOne({
    first_name: first_name.value,
    last_name: last_name.value,
    username: username.value,
    info: info.value,
  })
}

</script>
