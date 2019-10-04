<template>
  <div style="overflow:auto">
    <div v-infinite-scroll="scroll" infinite-scroll-disabled="disabled">
      <div v-for="(item, index) in data" :key="index">
        <slot v-bind:item="item">
          <div>No data</div>
        </slot>
      </div>
    </div>
    <slot name="loading" v-if="loading">
      <p class="loading">{{ loadingText }}</p>
    </slot>
    <slot name="nomore" v-if="noMore">
      <p class="nomore">{{ noMoreText }}</p>
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
        return 'Loading...'
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
.nomore {
  text-align: center;
  font-weight: 550;
}
</style>
