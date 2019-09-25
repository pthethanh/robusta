<template>
  <div class="challenge-new-view">
    <challenge-editor :mode="mode"></challenge-editor>
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
      mode: 'new'
    }
  }
}
</script>

<style scoped>
.challenge-new-view {
  padding-bottom: 100px;
}
</style>
