<template>
<CRow>
  <CCol class="mb-3" xs="11">
    <CFormLabel
      :for="props.inputName">
      {{ props.inputTitle }}
    </CFormLabel>
    <CFormInput
      v:model="props.inputModel"
      type="text"
      :id="props.inputName"
      :placeholder="props.inputPlaceholder"
      @input="handleInput"
    />
  </CCol>
  <CCol xs="1">
    <CFormLabel :for="props.opName">{{ props.opTitle }}</CFormLabel>
    <CFormSelect
      v:model="props.opModel"
      :options="props.opOptions"
      :name="props.opName"
      :id="props.op"
      @change="handleOp"
    ></CFormSelect>
  </CCol>
</Crow>
</template>

<script setup>

const props = defineProps([
  "inputName",
  "inputTitle",
  "inputModel",
  "opName",
  "opTitle",
  "opModel",
  "opOptions",
  "inputPlaceholder",
])

const emit = defineEmits([
  "update:inputModel",
  "update:opModel",
])

const handleInput = (event) => {
  emit("update:inputModel", event.target.value)
}

const handleOp = (event) => {
  emit("update:opModel", event.target.value)
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