<template>
    <blockquote class="quote">
        <p class="quotation" v-html="data.text"></p>
        <footer v-if="data.caption !== ''">{{data.caption}}</footer>
    </blockquote>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>
