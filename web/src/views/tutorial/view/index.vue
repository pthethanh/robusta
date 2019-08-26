<template>
  <div v-if="ready">
    <tutorial-detail :data="tutorial"></tutorial-detail>
  </div>
</template>

<script>
import TutorialDetail from '@/components/TutorialDetail'
import {
  fetchTutorial
} from '@/api/tutorial'

export default {
  components: {
    TutorialDetail
  },
  title: 'Goway - Latest tutorials',
  data() {
    return {
      ready: false,
      tutorial: {}
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.tutorial.id = this.$route.params.id
      fetchTutorial(this.$route.params.id).then(response => {
        this.tutorial = response.data
        this.ready = true
      })
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
