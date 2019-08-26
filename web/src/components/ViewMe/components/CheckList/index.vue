<template>
  <div class="cdx-checklist">
    <div v-for="(item, index) in data.items" :key="index" v-bind:class="{'cdx-checklist__item--checked': item.checked}" class="cdx-checklist__item">
      <span class="cdx-checklist__item-checkbox"></span>
      <div class="cdx-checklist__item-text" v-html="item.text"></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>

<style lang="scss" scoped>
.cdx-checklist {
  &__item {
    display: flex;
    padding: 0 10px;
    box-sizing: content-box;

    &-text {
      outline: none;
      flex-grow: 1;
      padding: 10px 0;
    }

    &-checkbox {
      display: inline-block;
      flex-shrink: 0;
      position: relative;
      width: 20px;
      height: 20px;
      margin: 10px 10px 10px 0;
      border-radius: 50%;
      border: 1px solid #d0d0d0;
      background: #fff;
      cursor: pointer;
      user-select: none;

      &:hover {
        border-color: #b5b5b5;
      }

      &::after {
        position: absolute;
        top: 5px;
        left: 5px;
        width: 8px;
        height: 5px;
        border: 2px solid #fcfff4;
        border-top: none;
        border-right: none;
        background: transparent;
        content: '';
        opacity: 0;
        transform: rotate(-45deg);
      }
    }

    &--checked {
      &-checkbox {
        background: #388ae5;
        border-color: #388ae5;

        &:hover {
          background: #307cd1;
        }

        &::after {
          opacity: 1;
        }
      }
    }
  }
}
</style>
