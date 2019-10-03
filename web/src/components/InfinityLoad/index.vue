<template>
  <div style="overflow:auto">
    <ul v-infinite-scroll="croll" infinite-scroll-disabled="disabled">
      <div v-for="(item, index) in list" :key="index">
        <slot v-bind:item="item">
          <div>No data</div>
        </slot>
      </div>
    </ul>
    <slot name="loading" v-if="loading">
      <p class="loading">Loading...</p>
    </slot>
    <slot name="nomore" v-if="noMore">
      <p class="nomore">¯\_(ツ)_/¯</p>
    </slot>
  </div>
</template>

<script>
export default {
  props: {
    'fetch-data': {
      type: Function
    },
    'data': {
      type: Array
    }
  },
  data () {
    return {
      loading: false,
      noMoreData: false,
      list: []
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
    croll () {
      this.loading = true
      this.fetchData().then(response => {
        this.loading = false
        if (response.data === null || response.data.length === 0) {
          this.noMoreData = true
          return
        }
        this.list = this.list.concat(response.data)
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
