<template>
  <div style="overflow:auto">
    <div v-infinite-scroll="scroll" infinite-scroll-disabled="disabled" infinite-scroll-delay="delay">
      <slot v-bind:data="data">
        <div v-if="noData" class="nodata">{{ $t('gen.no_data') }}</div>
      </slot>
    </div>
    <slot name="loading" v-if="loading">
      <div class="loading">{{ loadingText }}</div>
    </slot>
    <slot name="nomore" v-if="noMore && noMoreText !== ''">
      <div class="nomore">{{ noMoreText }}</div>
    </slot>
  </div>
</template>

<script>
export default {
  props: {
    fetchData: {
      type: Function
    },
    noMoreText: {
      type: String,
      default: function () {
        return '¯\\_(ツ)_/¯'
      }
    },
    loadingText: {
      type: String,
      default: function () {
        return this.$i18n.t('gen.loading')
      }
    },
    delay: {
      type: Number,
      default: function () {
        return 200
      }
    },
    limit: {
      type: Number,
      default: function () {
        return 0
      }
    }
  },
  data () {
    return {
      loading: false,
      noMoreData: false,
      data: []
    }
  },
  computed: {
    noMore () {
      return this.noMoreData
    },
    disabled () {
      return this.loading || this.noMore
    },
    noData () {
      return this.noMore && this.data.length === 0
    }
  },
  methods: {
    scroll () {
      this.loading = true
      this.fetchData().then(response => {
        this.loading = false
        if (response.data === null || response.data.length === 0) {
          this.noMoreData = true
          return
        }
        this.data = this.data.concat(response.data)
        this.$emit('update:data', this.data)
        this.noMoreData = this.limit > response.data.length
      }).catch((err) => {
        this.$emit('error', err)
      }).finally(() => {
        this.loading = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.loading,
.nomore,
.nodata {
  text-align: center;
  font-weight: 550;
}
</style>
