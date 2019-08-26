<template>
  <div class="article-edit-view">
    <article-editor :mode="mode" :id="id"></article-editor>
  </div>
</template>

<script>
import ArticleEditor from '@/components/ArticleEditor'

export default {
  components: {
    ArticleEditor
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
  },
  methods: {}
}
</script>

<style scoped>
.article-edit-view {
  padding-bottom: 100px;
}
</style>
