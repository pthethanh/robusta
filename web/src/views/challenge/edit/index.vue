<template>
  <div class="challenge-edit-view">
    <challenge-editor :mode="mode" :id="id"></challenge-editor>
  </div>
</template>

<script>
import ChallengeEditor from '@/components/ChallengeEditor'

export default {
  components: {
    ChallengeEditor
  },
  data () {
    return {
      mode: 'edit'
    }
  },
  computed: {
    id: function () {
      return this.$route.params.id
    }
  },
  methods: {}
}
</script>

<style scoped>
.challenge-edit-view {
  padding-bottom: 100px;
}
</style>
