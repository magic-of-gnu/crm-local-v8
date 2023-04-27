<template>
<CToaster placement="top-end" visible>
  <CToast
    v-for="(toast, index) in props.toasts"
    :key="index"
    :color="toast.color"
    visible="true"
  >
    <div class="d-flex">
      <CToastBody>
        {{ toast.content }}
      </CToastBody>  
      <CToastClose class="me-2 m-auto"/>
    </div>
  </CToast>
</CToaster>

</template>

<script setup>
import { CToast, CToaster, CToastBody, CToastClose } from '@coreui/vue'
const props = defineProps(['toasts'])
</script>