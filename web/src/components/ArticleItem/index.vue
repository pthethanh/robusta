<template>
  <div class="article">
    <el-card class="card">
      <div class="header" @click="select">{{ $t('gen.posted_by') }} {{article.created_by_name}} {{article.created_at | date }}</div>
      <div class="title" @click="select">{{article.title}}</div>
      <div class="content" @click="select" v-html="article.abstract" v-if="isAbstractPresent()"></div>
      <div class="content" @click="select" v-if="!isAbstractPresent()">
        <view-me :data="article.content" mode="preview"></view-me>
      </div>
      <ul class="footer">
        <li class="footer-item el-icon-arrow-up" @click="upvote">
          <span class="footer-item-value">{{article.reaction_upvote}}</span>
        </li>
        <li class="footer-item el-icon-arrow-down" @click="downvote">
          <span class="footer-item-value">{{article.reaction_downvote}}</span>
        </li>
        <li class="footer-item el-icon-view" @click="select">
          <span class="footer-item-value">{{article.views}}</span>
        </li>
        <li class="footer-item el-icon-chat-line-square" @click="select">
          <span class="footer-item-value">{{article.comments}}</span>
        </li>
        <li class="footer-item tags hidden-xs-only">
          <el-tag v-for="tag in importantTags()" :key="tag.label" effect="light" size="mini" :type="tag.type" class="tag" @click="tagSelected(tag)">
            {{ tag.label }}
          </el-tag>
        </li>
      </ul>
    </el-card>
  </div>
</template>

<script>
import ViewMe from '@/components/ViewMe'
import {
  reactToArticle
} from '@/api/article'
import {
  getTagType
} from '@/utils/tag'
import {
  mapGetters
} from 'vuex'
export default {
  props: {
    'article': Object
  },
  components: {
    ViewMe
  },
  computed: {
    ...mapGetters([
      'authenticated'
    ])
  },
  methods: {
    isAbstractPresent () {
      return this.article.abstract !== ''
    },
    upvote () {
      this.react('upvote')
    },
    downvote () {
      this.react('downvote')
    },
    react (reactionType) {
      if (!this.authenticated) {
        this.$store.dispatch('ToggleLogin', true)
        return
      }
      if (reactionType === 'downvote') {
        this.article.reaction_downvote++
      } else {
        this.article.reaction_upvote++
      }
      reactToArticle(this.article.id, reactionType).then((response) => {
        this.article.reaction_downvote = response.data.downvote
        this.article.reaction_upvote = response.data.upvote
      })
    },
    select () {
      this.$emit('selected', this.article)
    },
    importantTags () {
      let limit = 5
      let tags = this.article.tags
      if (tags === null) {
        return []
      }
      tags = tags.slice(0, limit)
      let typedTags = []
      for (var i = 0; i < tags.length; i++) {
        typedTags.push({
          type: getTagType(tags[i]),
          label: tags[i]
        })
      }
      return typedTags
    },
    tagSelected (tag) {
      this.$emit('tagSelected', tag)
    }
  }
}
</script>

<style lang="scss">
@import '@/styles/variables.scss';

.article {
  margin: 5px 0px;

  .header {
    margin-bottom: 5px;
    font-family: 'Courier New', Courier, monospace;
    font-size: 0.8em;
    font-weight: lighter;
    color: $fontColorInfo;
  }

  .title {
    font-size: 1.1em;
    font-weight: 700;
    padding-bottom: 10px;
    color: $fontColorHeading;
  }

  .content {
    max-height: 250px;
    overflow-y: hidden;
    margin-bottom: 5px;
    color: $fontColorContent;
  }

  .footer {
    padding-inline-start: 0px;
    display: inline-block;
    color: $fontColorInfo;
    margin-bottom: 0px;
    width: 100%;

    .footer-item {
      display: inline;
      margin-right: 20px;
      font-weight: 700;
      font-size: 0.85em;

      .footer-item-value {
        padding-left: 5px;
        font-family: 'Open Sans', sans-serif;
      }

      &:hover {
        color: $fontColorInfoHover;
      }
    }

    .tags {
      overflow: hidden;

      .tag {
        margin-left: 5px;
      }
    }
  }
}
</style>
