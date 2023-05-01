<template>
<CIcon
  v-if="props.sort_val"
  class="mx-2"
  :icon="props.icon_name_true"
  size="sm" 
  @click="props.method"
/>
<CIcon
  v-else
  class="mx-2"
  :icon="props.icon_name_false"
  size="sm" 
  @click="props.method"
/>
</template>

<script setup>

const props = defineProps([
  "sort_val",
  "icon_name_true",
  "icon_name_false",
  "method",
])

// sort_val
// icon_name_true
// icon_name_false
// method
</script>