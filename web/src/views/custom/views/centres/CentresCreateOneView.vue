<template>
<ToastComponent
v-if="toasts && toasts.length > 0"
:toasts="toasts"
/>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="centre_name">Centre Name</CFormLabel>
        <CFormInput
          v-model="name"
          type="string"
          id="centre_name"
          name="name"
          placeholder="Enter Centre Name Here"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="centre_description">Description</CFormLabel>
        <CFormTextarea
          v-model="description"
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
import m from '@/views/custom/hooks/centres/methods.js'
import { ref } from 'vue';
import ToastComponent from '@/components/ToastComponent.vue'

const name = ref(null)
const description = ref(null)
const toasts = ref([])

async function run_this_method(event) {

  const response = await m.postCreateOne({
    name: name.value,
    description: description.value,
  })

  if (response.data.hasOwnProperty("toasts")) {
    response.data.toasts.forEach((item) => {
      toasts.value.push(item)
    });
  }

  name.value = null
  description.value = null
}
</script>
