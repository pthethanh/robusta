<template>
  <div>
    <article-editor :mode="mode"></article-editor>
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
      mode: 'new'
    }
  }
}
</script>
