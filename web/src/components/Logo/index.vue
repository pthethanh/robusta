<template>
  <span v-if="raw" @click="onClick" class="logo-raw"><b :style="{color: color1}">Go</b><b :style="{color: _color2}">way</b></span>
  <div v-else v-bind:class="{'logo-large': isLarge, 'logo': !isLarge}" @click="onClick">
    <svg @click="onClick" :height="height" :width="width">
      <text y="30" font-weight="bold" font-family="'Open Sans', sans-serif">
        <tspan :fill="color1">Go</tspan>
        <tspan :fill="_color2">way</tspan>
      </text>
    </svg>
  </div>
</template>

<script>
import variables from '@/styles/variables.scss'

export default {
  props: {
    size: String,
    color1: {
      type: String,
      default: function () {
        return variables.logoColor1
      }
    },
    color2: {
      type: String,
      default: function () {
        return variables.logoColor2
      }
    },
    contrast: {
      type: Boolean,
      default: function () {
        return false
      }
    },
    raw: Boolean
  },
  computed: {
    height () {
      if (this.size === 'small' || this.size === undefined) {
        return 48
      }
      if (this.size === 'mini') {
        return 48
      }
      return 48
    },
    width () {
      if (this.size === 'small' || this.size === undefined) {
        return 72
      }
      if (this.size === 'mini') {
        return 70
      }
      return 105
    },
    isLarge () {
      return this.size === 'large'
    },
    _color2 () {
      if (this.contrast) {
        return variables.logoColor3
      }
      return this.color2
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

<style lang="scss" scoped>
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

.logo-raw {
  cursor: pointer;
}
</style>
