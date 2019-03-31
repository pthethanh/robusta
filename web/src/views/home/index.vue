<template>
  <div class="home">
    <el-row type="flex" justify="center">
      <el-col :span="14">
        <Article v-for="(article,index) in articles" :key="index" :article="article" @selected="openDetail"></Article>
      </el-col>
      <el-col :span="5" style="margin: 19px 10px;" class="hidden-xs-only">
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
    <el-row>
      <el-dialog
        :visible.sync="isOpenDetail"
        :center=true
        :modal=true
        :append-to-body=true
        :fullscreen=true
        @close="detailClosed"
        >
        <ArticleDetail :article="selectedArticle"></ArticleDetail>
      </el-dialog>
    </el-row>
  </div>
</template>

<script>
import { fetchList } from '@/api/article'
import Article from '@/components/Article'
import ArticleDetail from '@/components/ArticleDetail'
export default {
  components: {
    Article,
    ArticleDetail
  },
  data () {
    return {
      articles: [
      ],
      isOpenDetail: false,
      selectedArticle: null
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
    },
    openDetail (article) {
      this.selectedArticle = article
      this.isOpenDetail = true
      history.pushState({}, null, '/articles/detail/' + this.selectedArticle.id)
    },
    detailClosed () {
      history.pushState({}, null, '/')
    }
  }
}
</script>

<style>
  .home {
      margin: 0px 75px 0 75px;
  }
  .el-dialog{
    height: 100%;
  }
  .el-dialog__close.el-icon.el-icon-close {
    color: grey;
    font-weight: 700;
    font-size: 1.5em;
    border: 0.5px solid grey;
    position: fixed;
    top: 10px;
    right: 20px;
  }
</style>
