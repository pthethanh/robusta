<template>
  <div class="article">
    <article-detail :article="article" v-if="ready"></article-detail>
  </div>
</template>

<script>
import {
  fetchArticle
} from '@/api/article'
import ArticleDetail from '@/components/ArticleDetail'
export default {
  components: {
    ArticleDetail
  },
  data() {
    return {
      article: {},
      ready: false
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.article.id = this.$route.params.id
      fetchArticle(this.$route.params.id).then(response => {
        this.article = response.data
        this.ready = true
      })
    }
  }
}
</script>

<style scoped>
.article {
  background-color: white;
  margin: 0px 20px;
}
</style>
