<template>
  <div v-bind:class="{'logo-large': _isLarge, 'logo': !_isLarge}" @click="onClick">
    <svg @click="onClick" :height="_height" :width="_width"><text y="30" font-weight="bold" font-family="'Open Sans', sans-serif">
        <tspan :fill="color1">Go</tspan>
        <tspan :fill="color2">way</tspan>
      </text></svg>
  </div>
</template>

<script>
export default {
  props: {
    'size': String,
    'color1': {
      type: String,
      default: function () {
        return '#51AEE7'
      }
    },
    'color2': {
      type: String,
      default: function () {
        return 'white'
      }
    }
  },
  computed: {
    _height () {
      if (this.size === 'small' || this.size === undefined) {
        return 48
      }
      return 48
    },
    _width () {
      if (this.size === 'small' || this.size === undefined) {
        return 72
      }
      return 110
    },
    _isLarge () {
      return this.size === 'large'
    }
  },
  methods: {
    onClick () {
      this.$router.push('/')
      this.$emit('click')
    }
  }
}
</script>

<style scoped>
.logo {
  font-size: 1.2em;
  vertical-align: middle;
  fill: currentColor;
  overflow: hidden;
  align-items: center;
}

.logo-large {
  font-size: 1.8em;
  vertical-align: middle;
  fill: currentColor;
  overflow: hidden;
  align-items: center;
}
</style>
