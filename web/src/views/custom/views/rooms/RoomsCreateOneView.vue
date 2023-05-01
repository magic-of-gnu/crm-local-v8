<template>
<ToastComponent
v-if="toasts"
:toasts="toasts"
/>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="centre_id">Centre Name</CFormLabel>
        <CFormSelect
          v-model="selectedCentreID"
          aria-label="Default select example"
          :options="optionsCentres"
          name="centre_id"
          aria-placeholder="Select"
        >
        </CFormSelect>
      </div>
      <div class="mb-3">
        <CFormLabel for="room_name">Room Name</CFormLabel>
        <CFormInput
          v-model="name"
          type="text"
          id="name"
          name="name"
          placeholder="Enter Room Name"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="num_seats">Number of Seats</CFormLabel>
        <CFormInput
          v-model.number="num_seats"
          type="number"
          id="num_seats"
          name="num_seats"
          placeholder="Enter Number of Seats"
        />
      </div>
      <div class="mb-3">
        <CFormLabel for="info">Room Information</CFormLabel>
        <CFormTextarea
          v-model.trim="info"
          id="info"
          type="string"
          name="info"
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
import m from '@/views/custom/hooks/rooms/methods.js'
import centresMethods from '@/views/custom/hooks/centres/methods.js'
import ToastComponent from '@/components/ToastComponent.vue'
import { ref, onBeforeMount } from 'vue'
import centres from '@/router/custom/centres';

const optionsCentres = ref([{
  label: "Select Centre",
  value: "",
  disabled: true,
  hidden: true,
  selected: true
}])

const selectedCentreID = ref(null)
const num_seats = ref(null)
const name = ref("")
const info = ref("")

const toasts = ref([])

function getAllCentresList() {

  centresMethods.getCentresList()
    .then((d) => {
      let fetched_data = d.data.data

      for (const item of fetched_data) {
        optionsCentres.value.push({
          label: item.name,
          value: item.id
        })
      }
    })
}

async function run_this_method(event) {

  const request_data = {
    centre_id: selectedCentreID.value,
    name: name.value,
    num_seats: num_seats.value,
    info: info.value,
  }

  const response = await m.postCreateOne(request_data)
  populateToastsFromResponse(response, toasts)

  selectedCentreID.value = null
  num_seats.value = null
  name.value = null
  info.value = null

}

onBeforeMount(async () => {
  const response_centres = await centresMethods.getAllList()

  for (const item of response_centres.data.data.data) {
    optionsCentres.value.push({
      label: item.name,
      value: item.id
    })
  }

})
</script>
