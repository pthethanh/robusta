<template>
  <div class="home">
    <div class="filter">
      <el-tag v-for="tag in tags" :key="tag.label" effect="dark" :type="tag.type" closable class="tag" @close="removeFilter(tag)">{{ tag.label }}</el-tag>
    </div>
    <el-row type="flex" justify="center" class="infinite-list-wrapper">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" v-infinite-scroll="fetchData" infinite-scroll-disabled="disabled">
        <article-item v-for="(article,index) in articles" :key="index" :article="article" @selected="openDetail" @tagSelected="tagSelected"></article-item>
        <div v-if="loading" class="loading">Loading...</div>
        <div v-if="!loading&&noMore" class="loading">¯\_(ツ)_/¯</div>
      </el-col>
    </el-row>
    <el-row>
      <el-dialog :visible.sync="isOpenDetail" :center="true" :modal="true" :append-to-body="true" :fullscreen="true" @close="detailClosed">
        <article-detail :article="selectedArticle" @deleted="isOpenDetail=false"></article-detail>
      </el-dialog>
    </el-row>
  </div>
</template>

<script>
import {
  fetchList
} from '@/api/article'
import ArticleItem from '@/components/ArticleItem'
import ArticleDetail from '@/components/ArticleDetail'
import {
  getTags
} from '@/utils/tag'
export default {
  components: {
    ArticleItem,
    ArticleDetail
  },
  title: 'Goway - Learn to Go',
  data () {
    return {
      articles: [],
      isOpenDetail: false,
      selectedArticle: null,
      offset: 0,
      limit: 15,
      loading: false,
      noMoreArticles: false,
      tags: []
    }
  },
  created () {
    this.tags = getTags(this.$route.query.tags)
    this.reload()
  },
  computed: {
    disabled () {
      return this.loading || this.noMore
    },
    noMore () {
      return this.noMoreArticles === true
    }
  },
  methods: {
    fetchData () {
      this.loading = true
      fetchList(this.getQueryStr()).then(response => {
        this.loading = false
        if (response.data === null || response.data.length === 0) {
          this.noMoreArticles = true
          return
        }
        this.articles = this.articles.concat(response.data)
        this.offset += response.data.length
      })
    },
    getQueryStr () {
      let query = 'offset=' + this.offset + '&limit=' + this.limit
      for (var i = 0; i < this.tags.length; i++) {
        query += '&tags=' + this.tags[i].label
      }
      return query
    },
    async openDetail (article) {
      this.selectedArticle = article
      this.isOpenDetail = true
      let id = this.selectedArticle.article_id
      if (id === undefined) {
        id = this.selectedArticle.id
      }
      history.pushState({}, null, '/articles/detail/' + id)
    },
    detailClosed () {
      history.back()
    },
    tagSelected (tag) {
      if (tag !== this.$route.query.tags) {
        let found = false
        for (var i = 0; i < this.tags.length; i++) {
          if (this.tags[i].label === tag.label) {
            found = true
          }
        }
        if (found) {
          return
        }
        this.tags.push(tag)
        this.reload()
      }
    },
    async reload () {
      this.articles = []
      this.offset = 0
      history.pushState({}, null, '/articles?' + this.getQueryStr())
      this.fetchData()
    },
    removeFilter (tag) {
      for (var i = 0; i < this.tags.length; i++) {
        if (this.tags[i].label === tag.label) {
          this.tags.splice(i, 1)
        }
      }
      this.reload()
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.home {
  background-color: $backgroundColorPrimary;
  min-height: 100vh;

  .filter {
    text-align: center;

    .tag {
      margin-left: 5px;
    }
  }
}

.loading {
  background-color: $backgroundColorPrimary;
  text-align: center;
  font-weight: 600;
}
</style>
