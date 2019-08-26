<template>
  <ul v-bind:class="{
      'cdx-block cdx-list cdx-list--ordered': data.style === 'ordered',
      'cdx-block cdx-list cdx-list--unordered': data.style !== 'ordered'
    }">
    <li v-for="(item,index) in data.items" class="cdx-list__item" :key="index" v-html="item"></li>
  </ul>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>

<style lang="scss" scoped>
.cdx-list {
    margin: 0;
    padding-left: 20px;
    outline: none;

    &__item {
        padding: 5.5px 0 5.5px 3px;
        line-height: 1.75em;
    }

    &--unordered {
        list-style: disc;
    }

    &--ordered {
        list-style: decimal;
    }
}
</style>
