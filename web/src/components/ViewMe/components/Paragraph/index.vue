<template>
  <p class="ce-paragraph cdx-block" v-html="data.text"></p>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>

<style lang="scss" scoped>
.ce-paragraph {
    line-height: 1.75em;
    outline: none;
}

.ce-paragraph p:first-of-type{
    margin-top: 0;
}

.ce-paragraph p:last-of-type{
    margin-bottom: 0;
}
</style>
