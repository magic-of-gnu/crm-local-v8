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
        <CFormLabel for="password">Password</CFormLabel>
        <CFormInput
          v-model="password"
          type="password"
          id="password"
          name="password"
          placeholder="Enter Password"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="password">Repeat Password</CFormLabel>
        <CFormInput
          v-model="password2"
          type="password"
          id="password"
          name="password"
          placeholder="Repeat Password"
        />
      </div>
      <div class="mb-3">
        <CFormSwitch
          v-model="is_admin"
          id="is_admin"
          label="Is Admin"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="num_seats">User Type</CFormLabel>
        <CFormInput
          v-model.number="user_type"
          type="number"
          id="user_type"
          name="user_type"
          placeholder="Enter User Type (must be integer)"
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
import m from '@/views/custom/hooks/users/methods.js'
import { CFormSwitch } from '@coreui/vue';
import { ref } from 'vue'

const first_name = ref("")
const last_name = ref("")
const username = ref("")
const password = ref("")
const password2 = ref("")
const is_admin = ref(false)
const user_type = ref("")
let is_admin_int = -1 

function run_this_method(event) {
  console.log("is_admin: ", is_admin)
  if (is_admin.value === true) {
    is_admin_int = 1;
  }
  m.postCreateOne({
    first_name: first_name.value,
    last_name: last_name.value,
    username: username.value,
    password: password.value,
    password2: password2.value,
    is_admin: is_admin_int,
    // is_admin: is_admin.value,
    user_type: user_type.value,
  })
}

</script>
