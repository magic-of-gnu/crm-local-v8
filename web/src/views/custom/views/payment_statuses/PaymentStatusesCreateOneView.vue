<template>
<ToastComponent 
  v-if="toasts"
  :toasts="toasts"
/>

  <CRow>
    <CCol xs="12">
    <CForm action="some-post-function" method="POST">
      <CRow>
        <CCol class="mb-3" xs="12">
          <CFormLabel for="name">Name</CFormLabel>
          <CFormInput
            type="text"
            id="name"
            placeholder="Enter Name"
            name="name"
            v-model="name"
          />
        </CCol>
      </CRow>

      <CRow>
        <CCol class="mb-3" xs="12">
          <CFormLabel for="description">Description</CFormLabel>
          <CFormTextarea 
            type="text"
            id="description"
            placeholder="Enter Description"
            name="description"
            v-model="description"
          />
        </CCol>
      </CRow>

      <div class="row my-3">
        <div class="col">
          <CButton
            @click.prevent="($event) => run_this_method($event)"
            component="input"
            type="submit"
            color="primary"
            value="Submit"
          />
        </div>
      </div>
    </CForm>
    </CCol>
  </CRow>
</template>

<script setup>
import { ref } from 'vue'
import m from '@/views/custom/hooks/paymentStatuses/methods.js'
import ToastComponent from '@/components/ToastComponent.vue'
import { populateToastsFromResponse } from '@/utils/toasts';

const toasts = ref([])

const name = ref(null)
const description = ref(null)

function run_this_method(event) {
  const request_data = {
    name: name.value,
    description: description.value
  }

  m.postCreateOne(request_data).then((response) => {
  populateToastsFromResponse(response, toasts)

  name.value = null
  description.value = null

  })

}

</script>



