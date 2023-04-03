<template>
  <CRow>
    <CForm action="some-post-function" method="POST">
      <div class="mb-3">
        <CFormLabel for="centre_id">Centre Name</CFormLabel>
        <CFormSelect
          v-model="centre_id"
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

import m from '@/views/custom/hooks/rooms/methods.js'
import centresMethods from '@/views/custom/hooks/centres/methods.js'
import { ref, onMounted } from 'vue'

const optionsCentres = ref([{}])
const selectedCentre = ref(null)

console.log('running...')

let centre_id = null
let name = ""
let num_seats = null
let info = ""

function getAllCentresList() {
  console.log("retrieving centres list")

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
console.log("m: ", m)


function run_this_method(event) {
  // console.log('running button click')
  // console.log("event: ", event)

  // centre_id = event.target.form.elements.centre_id.value
  // name = event.target.form.elements.name.value
  // num_seats = event.target.form.elements.num_seats.value
  // info = event.target.form.elements.info.value

  // console.log("centre_id: ", centre_id, "name: ", name, "num_seats: ", num_seats, "info: ", info)
  // console.log("selectedCentre: ", selectedCentre)

  m.postRoomsCreateOne({
    centre_id: centre_id,
    name: name,
    num_seats: num_seats,
    info: info,
  })
}

onMounted(() => {
  getAllCentresList()
})
</script>
