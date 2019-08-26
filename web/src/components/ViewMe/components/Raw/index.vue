<template>
     <div v-html="data.html"></div>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>
