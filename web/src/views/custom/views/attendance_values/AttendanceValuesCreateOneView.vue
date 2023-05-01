<template>
<ToastComponent
v-if="toasts"
:toasts="toasts"
/>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="value">Value</CFormLabel>
        <CFormInput
          v-model.number="value"
          type="text"
          id="value"
          name="value"
          placeholder="Enter Course value (has to be a number)"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="first_name">Name</CFormLabel>
        <CFormInput
          v-model.trim="name"
          type="text"
          id="name"
          name="name"
          placeholder="Enter Course Name"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="description">Description</CFormLabel>
        <CFormInput
          v-model.trim="description"
          type="text"
          id="description"
          name="description"
          placeholder="Enter Description"
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
import m from '@/views/custom/hooks/attendanceValues/methods.js'
import { ref } from 'vue'

const value = ref("")
const name = ref("")
const description = ref("")
const toasts = ref("")

async function run_this_method(event) {
  const response = await m.postCreateOne({
    value: value.value,
    name: name.value,
    description: description.value,
  })

  if (response.data.hasOwnProperty("toasts")) {
    response.data.toasts.forEach((item) => {
      toasts.value.push(item)
    });
  }
}

</script>
