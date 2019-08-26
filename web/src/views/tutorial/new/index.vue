<template>
  <div>
    <TutorialEditor mode="new"></TutorialEditor>
  </div>
</template>

<script>
import TutorialEditor from '@/components/TutorialEditor'

export default {
  components: {
    TutorialEditor
  },
  data() {
    return {
      tutorial: {
        steps: []
      }
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
