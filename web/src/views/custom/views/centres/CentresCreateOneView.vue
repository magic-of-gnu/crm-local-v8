<template>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="centre_name">Centre Name</CFormLabel>
        <CFormInput
          type="string"
          id="centre_name"
          name="name"
          placeholder="Enter Centre Name Here"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="centre_description">Description</CFormLabel>
        <CFormTextarea
          id="exampleFcentre_description"
          type="string"
          name="description"
          rows="3"
          placeholder="Enter Description Here"
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

import m from '@/views/custom/hooks/centres/methods.js'
console.log('running...')
let name = ""
let description = ""

function run_this_method(event) {
  console.log('running button click')
  console.log("event: ", event)

  name = event.target.form.elements.name.value
  description = event.target.form.elements.description.value
  
  console.log("name: ", name, "description: ", description)

  m.postCentresCreateOne({
    name: name,
    description: description,
  })
}
</script>
