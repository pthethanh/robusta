<template>
  <div class="home">
    <el-row type="flex" justify="center">
      <el-col :span="18">
        <Article v-for="article in articles" :key="article.id" :article="article"></Article>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { fetchList } from '@/api/article'
import Article from '@/components/Article'
export default {
  components: {
    Article
  },
  data () {
    return {
      articles: [
      ]
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      fetchList().then(response => {
        console.log(response.data)
        this.articles = response.items
      })
    }
  }
}
</script>

<style>
  .home {
      margin: 0px 75px 0 75px;
  }
</style>
