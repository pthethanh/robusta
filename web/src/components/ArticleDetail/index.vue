<template>
  <div>
    <el-row type="flex" justify="center">
      <el-col :span="14">
        <div class="title">{{article.title}}</div>
        <div class="info">
          Posted by {{article.created_by_name}} {{article.created_at|moment('from')}}
          <br>
          Source: {{article.source}}
        </div>
        <div class="abstract" v-html="article.abstract"></div>
        <div class="content" v-html="article.content"></div>
      </el-col>
    </el-row>  
  </div>
</template>

<script>
import { viewArticle } from '@/api/article'
import router from '@/router'
export default {
  props: {
    'article': Object
  },
  created () {
    this.increaseArticleViews()
  },
  methods: {
    increaseArticleViews() {
      console.log('article.id: ' + this.article.id)
      viewArticle(this.article.id)
    }
  }
}
</script>

<style scoped>
.title {
  font-size: 2.5em;
  font-weight: 900;
  padding-bottom: 30px;
}
.info {
  font-weight: 500;
  font-size: 1.1em;
  font-family: 'Courier New', Courier, monospace;
  font-style: initial;
  padding-left: 20px;
  border-left: 2px solid grey;
}
.abstract, .content {
  padding-top: 20px;
  font-size: 17px;
  line-height: 1.58;
  letter-spacing: -.004em;
}
</style>
