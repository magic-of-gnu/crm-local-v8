<template>
<ToastComponent
v-if="toasts"
:toasts="toasts"
/>
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
import ToastComponent from '@/components/ToastComponent.vue'
import m from '@/views/custom/hooks/students/methods.js'
import { ref } from 'vue'

const first_name = ref("")
const last_name = ref("")
const username = ref("")
const toasts = ref([])

async function run_this_method(event) {
  const response = await m.postStudentsCreateOne({
    first_name: first_name.value,
    last_name: last_name.value,
    username: username.value,
  })

  if (response.data.hasOwnProperty("toasts")) {
    response.data.toasts.forEach((item) => {
      toasts.value.push(item)
    });
  }
}

</script>
