<template>
  <div>
    <TutorialEditor mode="edit" :id="id"></TutorialEditor>
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
      mode: 'edit'
    }
  },
  computed: {
    id: function () {
      return this.$route.params.id
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
