<template>
  <div class="article-detail">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="16" :lg="12" :xl="12">
        <h1 class="article-heading title">{{article.title}}</h1>
        <div class="metadata">
          <div class="info">
            <avatar class="info__avatar" :src="article.created_by_avatar" :names="[article.created_by_name]"></avatar>
            <div class="info__detail">
              <span class="info__detail__author">{{article.created_by_name}}</span>
              <span>{{article.created_at | date }}</span>
            </div>
          </div>
          <el-dropdown class="ctx-menu" v-show="user.info.user_id === article.created_by_id" @command="handleSettings" trigger="click">
            <i class="el-icon-more ctx-menu-btn"></i>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item icon="el-icon-edit" command="edit">{{ $t('gen.edit') }}</el-dropdown-item>
              <el-dropdown-item icon="el-icon-delete" command="delete">{{ $t('gen.delete') }}</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
        <div class="abstract" v-html="article.abstract"></div>
        <view-me :data="article.content" class="content"></view-me>
        <div class="tags">
          <el-tag v-for="tag in getAllTags()" :key="tag.label" effect="dark" size="mini" :type="tag.type" class="tag">
            {{ tag.label }}
          </el-tag>
        </div>
        <comments :targetID="article.id" targetType='article' class="comments"></comments>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  updateArticleView,
  deleteArticle
} from '@/api/article'
import ViewMe from '@/components/ViewMe'
import {
  mapGetters
} from 'vuex'
import {
  getTags
} from '@/utils/tag'
import Comments from '@/components/Comments'
import Avatar from '@/components/Avatar'
export default {
  props: {
    'article': Object
  },
  components: {
    ViewMe,
    Comments,
    Avatar
  },
  title: '',
  computed: {
    ...mapGetters([
      'user'
    ])
  },
  data () {
    return {}
  },
  created () {
    this.increaseArticleViews()
    if (this.article.created_by_name === '') {
      this.article.created_by_name = 'goway'
    }
    if (this.article.source === '') {
      this.article.source = 'goway'
    }
    this.$title = this.article.title
  },
  methods: {
    async increaseArticleViews () {
      updateArticleView(this.article.id)
    },
    handleSettings (command) {
      if (command === 'edit') {
        this.$router.push('/articles/edit/' + this.article.id)
        return
      }
      if (command === 'delete') {
        this.$confirm(this.$i18n.t('article.delete_confirm'), this.$i18n.t('gen.warning'), {
          confirmButtonText: this.$i18n.t('gen.ok'),
          cancelButtonText: this.$i18n.t('gen.cancel'),
          type: 'warning'
        }).then(() => {
          deleteArticle(this.article.id).then((respose) => {
            this.$message({
              type: 'success',
              message: this.$i18n.t('gen.delete_success')
            })
            this.$router.push('/')
            this.$emit('deleted')
          })
        })
      }
    },
    getAllTags () {
      return getTags(this.article.tags)
    }
  }
}
</script>

<style lang="scss" scoped>
.article-detail {
  .article-heading {
    margin-top: 10px;
  }

  .metadata {
    margin-bottom: 10px;

    .info {
      display: inline-block;
      margin-bottom: 10px;
      width: 90%;
      word-wrap: break-word;

      &__avatar {
        vertical-align: text-bottom;
        display: inline-block
      }

      &__detail {
        display: inline-block;
        font-weight: 500;
        font-size: 0.95em;
        letter-spacing: -0.01rem;
        line-height: 1.5em;
        color: grey;
        font-family: 'Courier New', Courier, monospace;
        font-style: normal;
        margin-left: 10px;

        &__author {
          font-weight: 700;
        }

        span {
          display: block;
        }
      }
    }

    .ctx-menu {
      float: right;
      vertical-align: top;

      .ctx-menu-btn {
        font-size: 1.2em;
        color: black;

        &:hover {
          cursor: pointer;
        }
      }
    }
  }

  .tags {
    margin-top: 10px;
    margin-bottom: 10px;

    .tag {
      margin-left: 5px;
    }
  }

  .comments {
    margin-top: 50px;
  }
}
</style>
