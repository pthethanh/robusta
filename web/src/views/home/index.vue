<template>
  <div class="home">
    <el-row type="flex" justify="center">
      <el-col>
        <Article v-for="article in articles" :key="article.id" :article="article"></Article>
      </el-col>
      <el-col :span="8" style="margin: 19px 10px;" class="hidden-xs-only">
        <el-card>
          <div>
            <h1>Hot news</h1>
          </div>
          <div>
            <ul>
              <li>Go 12.1 released</li>
              <li>New program for developer to interview with GO</li>
            </ul>
          </div>
        </el-card>
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
