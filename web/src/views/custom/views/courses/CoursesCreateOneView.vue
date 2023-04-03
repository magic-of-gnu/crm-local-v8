<template>
  <CRow>
    <CForm action="some-post-function" method="POST">
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
/* eslint-disable */

import m from '@/views/custom/hooks/courses/methods.js'
import { ref } from 'vue'

const name = ref("")
const description = ref("")

function run_this_method(event) {
  m.postCreateOne({
    name: name.value,
    description: description.value,
  })
}

</script>
