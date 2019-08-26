<template>
  <div class="cdx-block">
    <div class="link-tool">
      <a class="link-tool__content link-tool__content--rendered" target="_blank" rel="nofollow noindex noreferrer" :href="data.link"
        v-bind:class="{
          'disabled-clickable': mode==='preview',
        }"
      >
        <div class="link-tool__image" v-bind:style="styleValue"></div>
        <div class="link-tool__title">{{data.meta.title}}</div>
        <p class="link-tool__description">{{data.meta.description}}</p>
        <span class="link-tool__anchor">{{data.link}}</span>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    'data': Object,
    'mode': String
  },
  data () {
    return {
      styleValue: 'background-image: url(' + this.data.meta.image.url + ')'
    }
  }
}
</script>

<style lang="scss" scoped>
.link-tool {
  position: relative;
  font-family: 'Open Sans', sans-serif;
  overflow: hidden;

  .disabled-clickable {
    pointer-events: none;
    cursor: default;
  }

  &__content {
    display: block;
    padding: 25px;
    border-radius: 2px;
    -webkit-box-shadow: 0 0 0 2px #fff;
    box-shadow: 0 0 0 2px #fff;
    color: initial !important;
    text-decoration: none !important;

    &::after {
      content: "";
      clear: both;
      display: table;
    }

    &--rendered {
      background: #fff;
      border: 1px solid rgba(201, 201, 204, 0.48);
      -webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, .1);
      box-shadow: 0 1px 3px rgba(0, 0, 0, .1);
      border-radius: 6px;
      will-change: filter;
      -webkit-animation: link-in 450ms 1 cubic-bezier(0.215, 0.61, 0.355, 1);
      animation: link-in 450ms 1 cubic-bezier(0.215, 0.61, 0.355, 1);

      &:hover {
        -webkit-box-shadow: 0 0 3px rgba(0, 0, 0, .16);
        box-shadow: 0 0 3px rgba(0, 0, 0, .16);
      }
    }
  }

  &__image {
    background-position: center center;
    background-repeat: no-repeat;
    background-size: cover;
    margin: 0 0 0 30px;
    width: 65px;
    height: 65px;
    border-radius: 3px;
    float: right;
  }

  &__title {
    font-size: 17px;
    font-weight: 600;
    line-height: 1.5em;
    margin: 0 0 10px 0;

    +&__anchor {
      margin-top: 25px;
    }
  }

  &__description {
    margin: 0 0 20px 0;
    font-size: 15px;
    line-height: 1.55em;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  &__anchor {
    display: block;
    overflow: hidden;
    font-size: 15px;
    line-height: 1em;
    color: #888 !important;
    border: 0 !important;
    padding: 0 !important;
  }
}

.codex-editor--narrow .link-tool__image {
  display: none;
}
</style>
