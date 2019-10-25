<template>
  <div class="home">
    <div class="filter">
      <el-tag v-for="tag in tags" :key="tag.label" effect="dark" :type="tag.type" closable class="tag" @close="removeFilter(tag)">{{ tag.label }}</el-tag>
    </div>
    <el-row type="flex" justify="center" class="infinite-list-wrapper">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" v-infinite-scroll="fetchData" infinite-scroll-disabled="disabled">
        <article-item v-for="(article,index) in articles" :key="index" :article="article" @selected="openDetail" @tagSelected="tagSelected"></article-item>
        <div v-if="loading" class="loading">{{ $t('gen.loading') }}</div>
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
  getTags,
  tagsEquals
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
      tags: [],
      init: true
    }
  },
  mounted () {
    window.onpopstate = this.handleBackButton
    this.tags = getTags(this.$route.query.tags)
  },
  watch: {
    tags: function (o, n) {
      if (this.init) {
        this.init = false
        return
      }
      this.reload()
    }
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
        if (response.data.length < this.limit) {
          this.noMoreArticles = true
          return
        }
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
      history.pushState({}, null, this.getURL())
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
      }
    },
    async reload () {
      this.articles = []
      this.offset = 0
      this.noMoreArticles = false
      history.pushState({}, null, this.getURL())
      this.fetchData()
    },
    removeFilter (tag) {
      for (var i = 0; i < this.tags.length; i++) {
        if (this.tags[i].label === tag.label) {
          this.tags.splice(i, 1)
        }
      }
    },
    handleBackButton (event) {
      if (this.isOpenDetail) {
        this.isOpenDetail = false
      }
      var tags = getTags(this.$route.query.tags)
      if (!tagsEquals(tags, this.tags)) {
        this.tags = tags
      }
    },
    getURL () {
      var url = '/articles'
      if (this.tags.length > 0) {
        url += '?tags=' + this.tags[0].label
      }
      for (var i = 1; i < this.tags.length; i++) {
        url += '&tags=' + this.tags[i].label
      }
      return url
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
