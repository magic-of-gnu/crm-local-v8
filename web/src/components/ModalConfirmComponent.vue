<template>
<CModal 
  alignment="center" 
  backdrop="static"
  :visible="props.showModal"
>
  <CModalBody>
    {{ props.content }}
  </CModalBody>
  <CModalFooter>
    <CButton color="secondary" @click="props.confirmDeleteClick(false)">
      {{ props.closeBtnText }}
    </CButton>
    <CButton color="primary" @click="props.confirmDeleteClick(true)">
      {{ props.confirmBtnText }}
    </CButton>
  </CModalFooter>
</CModal>
</template>

<script setup>
const props = defineProps([
  "content",
  "closeBtnText",
  "confirmBtnText",
  "showModal",
  "confirmDeleteClick"
])
</script>