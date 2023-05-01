<template>
<CRow>
  <CCol class="mb-3" xs="11">
    <CFormLabel
      :for="props.yearName">
      {{ props.yearTitle }}
    </CFormLabel>
    <CFormInput
      v:model.number="props.yearModel"
      type="text"
      :id="props.yearName"
      :placeholder="props.yearPlaceholder"
      @input="handleInput"
    />
  </CCol>
  <CCol class="mb-3" xs="11">
    <CFormLabel
      :for="props.monthName">
      {{ props.monthTitle }}
    </CFormLabel>
    <CFormSelect
      v:model="props.selectModel"
      :options="props.selectOptions"
      :name="props.select"
      :id="props.select"
      @change="handleSelectOp"
    ></CFormSelect>
  </CCol>
  <CCol class="mb-3" xs="11">
    <CFormLabel
      :for="props.input">
      {{ props.inputTitle }}
    </CFormLabel>
    <CFormInput
      v:model="props.inputModel"
      type="text"
      :id="props.input"
      :placeholder="props.placeholder"
      @input="handleInput"
    />
  </CCol>
  <CCol xs="1">
    <CFormLabel :for="props.select">{{ props.selectTitle }}</CFormLabel>
    <CFormSelect
      v:model="props.selectModel"
      :options="props.selectOptions"
      :name="props.select"
      :id="props.select"
      @change="handleSelectOp"
    ></CFormSelect>
  </CCol>
</Crow>
</template>

<script setup>

const props = defineProps([
  "yearName",
  "yearTitle",
  "yearModel",
  "monthName",
  "monthTitle",
  "monthModel",
  "dayName",
  "dayTitle",
  "dayModel",
  "opName",
  "opTitle",
  "opModel",
  "opOptions",
  "placeholder",
])

const emit = defineEmits([
  "update:inputModel",
  "update:selectModel",
])

const handleInput = (event) => {
  emit("update:inputModel", event.target.value)
}

const handleSelectOp = (event) => {
  emit("update:selectModel", event.target.value)
}

// data.input
// data.inputTitle
// data.inputModel
// data.select
// data.selectTitle
// data.selectOptions
// data.selectlModel
// data.placeholder

</script>