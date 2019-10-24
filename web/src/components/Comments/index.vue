<template>
  <div class="comment-wrapper">
    <div class="comment-info">
      <span v-loading="_loading" class="info-txt">{{comments.length}} {{ $t('comment.comments') }}</span>
    </div>
    <div class="comment-box">
      <el-form :model="comment" :rules="rules" ref="commentForm">
        <el-form-item prop="content">
          <el-input type="textarea" class="text" :rows="3" :placeholder="$t('comment.placeholder')" v-model="comment.content"></el-input>
        </el-form-item>
        <el-button class="comment-btn" v-if="user.authenticated" @click="addNewComment" type="primary" size="mini">{{ $t('comment.add') }}</el-button>
        <el-button class="comment-btn" type="danger" size="mini" v-if="!user.authenticated" @click="goToLogin">{{ $t('comment.login') }}</el-button>
      </el-form>
    </div>
    <div v-if="comments !== null && comments.length > 0">
      <div v-for="(comment, index) in comments" :key="comment.id">
        <div class="comment" v-bind:class="{
                'comment-level-0': comment.level === 0,
                'comment-level-1': comment.level === 1,
                'comment-level-2': comment.level === 2,
                'comment-level-3': comment.level === 3,
                'comment-level-4': comment.level >= 4,
              }">
          <avatar class="avatar" size="small" :src="comment.created_by_avatar" :names="[comment.created_by_name]"></avatar>
          <div class="content">
            <a class="author">{{comment.created_by_name}}</a>
            <div class="metadata">
              <span class="date">{{comment.created_at | date }}</span>
            </div>
            <div class="menu" v-if="user.info.user_id === comment.created_by_id && comment.mode!=='edit'">
              <el-popover placement="bottom" trigger="click">
                <div>
                  <el-button @click="handleEditComment(comment)" icon="el-icon-edit" circle size="mini" type="primary"></el-button>
                  <el-button @click="handleDeleteComment(index, comment)" icon="el-icon-delete" circle size="mini" type="danger"></el-button>
                </div>
                <i slot="reference" class="el-icon-more"></i>
              </el-popover>
            </div>
            <div class="text" v-if="comment.mode!=='edit'">
              {{comment.content}}
            </div>

            <el-form :model="comment" :rules="rules" v-if="comment.mode==='edit'" ref="editForm" class="edit">
              <el-form-item prop="content">
                <el-input v-model="comment.content" type="textarea" class="text" :rows="3" :placeholder="$t('comment.placeholder')"></el-input>
              </el-form-item>
              <el-button class="comment-btn" type="primary" size="mini" @click="addReplyComment(comment)" v-if="user.authenticated && comment.id === ''">{{ $t('comment.reply') }}</el-button>
              <el-button class="comment-btn" type="primary" size="mini" @click="updateComment(comment)" v-if="user.authenticated && comment.id !== ''">{{ $t('comment.update') }}</el-button>
              <el-button class="comment-btn" type="danger" size="mini" v-if="!user.authenticated" @click="goToLogin">{{ $t('comment.login') }}</el-button>
              <el-button class="comment-btn" type="primary" size="mini" @click="cancel(index, comment)">{{ $t('gen.cancel') }}</el-button>
            </el-form>
            <div class="actions" v-if="comment.mode!=='edit'">
              <li class="footer-item el-icon-arrow-up" @click="upvote(comment)">
                <span class="footer-item-value">{{comment.reaction_upvote}}</span>
              </li>
              <li class="footer-item el-icon-arrow-down" @click="downvote(comment)">
                <span class="footer-item-value">{{comment.reaction_downvote}}</span>
              </li>
              <li @click="replyTo(index, comment)">{{ $t('comment.reply') }}</li>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  findComments,
  createComment,
  reactToComment,
  deleteComment,
  updateComment
} from '@/api/comment'
import {
  mapGetters
} from 'vuex'
import Avatar from '@/components/Avatar'
export default {
  props: {
    'targetID': String,
    'targetType': String
  },
  components: {
    Avatar
  },
  computed: {
    ...mapGetters([
      'user'
    ]),
    _loading () {
      return this.loading
    }
  },
  watch: {
    // whenever new targetID is passed, we need to reload the comments.
    // we have to do this because VueJS will create the component only once.
    targetID: function (oldID, newID) {
      this.loadComments()
    }
  },
  data () {
    return {
      comments: [],
      comment: {
        target: this.targetID,
        target_type: this.targetType,
        content: ''
      },
      rules: {
        content: [{
          required: true,
          message: this.$i18n.t('comment.validation.required'),
          trigger: 'blur'
        }]
      },
      loading: true
    }
  },
  created () {
    // only once when the component is created for the first time.
    this.loadComments()
  },
  methods: {
    async loadComments () {
      this.loading = true
      this.comments = []
      var query = 'target=' + this.targetID + '&sort_by=level&sort_by=-created_at'
      findComments(query).then((response) => {
        if (response.data !== null && response.data.length > 0) {
          this.buildCommentTree(response.data)
        }
      }).finally(() => {
        this.loading = false
      })
    },
    buildCommentTree (loadedComments) {
      for (var i = 0; i < loadedComments.length; i++) {
        var c = loadedComments[i]
        if (c.level === 0) {
          this.addChildCommentsToTree(c, loadedComments, this.comments)
        }
      }
    },
    addChildCommentsToTree (root, list, result) {
      root.mode = 'view' // set default to view
      result.push(root)
      for (var i = 0; i < list.length; i++) {
        var child = list[i]
        if (child.reply_to_id === root.id) {
          this.addChildCommentsToTree(child, list, result)
        }
      }
    },
    addNewComment () {
      this.$refs['commentForm'].validate((valid) => {
        if (!valid) {
          return false
        }
        createComment(this.comment).then((response) => {
          this.comments.unshift(response.data)
          this.comment.content = ''
        })
      })
    },
    addReplyComment (comment) {
      this.$refs['editForm'][0].validate((valid) => {
        if (!valid) {
          return false
        }
        createComment(comment).then((response) => {
          comment.id = response.data.id
          comment.mode = 'view'
        })
      })
    },
    updateComment (comment) {
      this.$refs['editForm'][0].validate((valid) => {
        if (!valid) {
          return false
        }
        updateComment(comment.id, comment).then((response) => {
          comment.mode = 'view'
        })
      })
    },
    upvote (comment) {
      comment.reaction_upvote++
      this.react('upvote', comment)
    },
    downvote (comment) {
      comment.reaction_downvote++
      this.react('downvote', comment)
    },
    react (reactionType, comment) {
      reactToComment(comment.id, reactionType).then((response) => {
        comment.reaction_downvote = response.data.downvote
        comment.reaction_upvote = response.data.upvote
      })
    },
    replyTo (index, comment) {
      var replyComment = {
        id: '',
        content: '',
        target: comment.target,
        target_type: comment.target_type,
        reply_to_id: comment.id,
        level: comment.level + 1,
        created_by_id: this.user.info.user_id,
        created_by_name: this.user.info.name,
        created_by_avatar: this.user.info.avatar_url,
        created_at: new Date(),
        reaction_upvote: 0,
        reaction_downvote: 0,
        thread_id: comment.thread_id,
        mode: 'edit' // enable edit mode
      }
      this.comments.splice(index + 1, 0, replyComment)
    },
    cancel (index, comment) {
      if (comment.id === '') { // new reply
        this.comments.splice(index, 1)
        return
      }
      comment.content = comment.orgContent
      comment.mode = 'view'
    },
    goToLogin () {
      this.$store.dispatch('ToggleLogin', true)
    },
    handleDeleteComment (index, comment) {
      this.$confirm(this.$i18n.t('comment.delete_confirm'), this.$i18n.t('gen.warning'), {
        confirmButtonText: this.$i18n.t('gen.ok'),
        cancelButtonText: this.$i18n.t('gen.cancel'),
        type: 'warning'
      }).then(() => {
        deleteComment(comment.id).then((respose) => {
          this.$message({
            type: 'success',
            message: this.$i18n.t('gen.delete_success')
          })
          this.comments.splice(index, 1)
        })
      })
    },
    handleEditComment (comment) {
      comment.mode = 'edit'
      comment.orgContent = comment.content
    }
  }
}
</script>

<style lang="scss" scoped>
.comment-wrapper {
  padding-bottom: 100px;

  .comment-info {
    .info-txt {
      font-size: 1.2em;
      font-weight: 700;
    }
  }

  .comment-box {
    margin-bottom: 40px;

    .text {
      margin: 0px;
    }

    .comment-btn {
      margin-top: -18px;
      float: right;
    }
  }

  .comment-level-0 {
    margin-left: 0em;
  }

  .comment-level-1 {
    margin-left: 2.5em;
  }

  .comment-level-2 {
    margin-left: 5em;
  }

  .comment-level-3 {
    margin-left: 7.5em;
  }

  .comment-level-4 {
    margin-left: 10em;
  }

  .comment {
    font-family: 'Open Sans', sans-serif;
    margin-top: 10px;
    padding-left: 5px;
    border-left: 0.8px solid lightgrey;

    .avatar {
      display: inline-block;
      vertical-align: text-bottom
    }

    .content {
      display: inline;
      vertical-align: text-bottom;
      margin-left: 10px;

      .author {
        font-size: 1em;
        color: rgba(0, 0, 0, .87);
        font-weight: 700;
      }

      .metadata {
        display: inline-block;
        margin-left: 1em;
        color: rgba(0, 0, 0, .4);
        font-size: .875em;
      }

      .menu {
        float: right;

        &:hover {
          cursor: pointer;
        }
      }

      .text {
        margin: .25em 0 .5em 2.5em;
        font-size: 1em;
        word-break: keep-all;
        overflow-x: auto;
        color: rgba(0, 0, 0, .87);
        line-height: 1.75em;
      }

      .edit {
        .text {
          margin: 0px;
        }

        .comment-btn {
          margin-top: -18px;
          margin-left: 5px;
          float: right;
        }
      }

      .actions {
        font-size: 0.95em;
        cursor: pointer;
        display: inline-block;
        margin: .25em 0 .5em 2.7em;
        color: rgba(0, 0, 0, .4);

        li {
          font-weight: 700;
          margin-left: 5px;
          display: inline;
          font-family: 'Open Sans', sans-serif;
        }
      }
    }
  }
}
</style>
