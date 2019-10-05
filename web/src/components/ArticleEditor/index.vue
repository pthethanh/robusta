<template>
  <div class="article-editor">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" v-if="ready">
        <div class="menu">
          <el-button type="primary" size="mini" class="menu-btn" @click="saveEditor">Save</el-button>
          <el-button type="primary" size="mini" class="menu-btn" @click="openReviewEditor">Preview</el-button>
        </div>
        <el-form :model="article" :rules="rules" ref="editorForm" class="form">
          <el-form-item prop="title">
            <el-input class="input-title" placeholder="title" v-model="article.title"></el-input>
          </el-form-item>
          <tags ref="tags" :initData="article.tags"></tags>
          <editor ref="editor" @save="save" class="editorjs" :initData="article.content"></editor>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Editor from '@/components/Editor'
import {
  createArticle,
  fetchArticle,
  updateArticle
} from '@/api/article'
import Tags from '@/components/Tags'

export default {
  components: {
    Editor,
    Tags
  },
  props: {
    'id': String,
    'mode': String
  },
  data () {
    return {
      ready: false, // only render component if everything is ready
      alreadyCreated: false,
      article: {
        tags: []
      },
      rules: {
        title: [{
          required: true,
          message: 'Please enter title',
          trigger: 'blur'
        }, {
          max: 256,
          message: 'Length should be less than 256 characters',
          trigger: 'blur'
        }]
      }
    }
  },
  title: '',
  created () {
    let self = this
    if (this.isEditMode()) {
      fetchArticle(this.id).then(response => {
        this.article = response.data
        this.ready = true
      }).catch(function (err) {
        self.$message({
          message: 'Failed to load data: ' + err,
          type: 'error'
        })
      })
      return
    }
    this.ready = true
    this.$title = this.article.title
  },
  methods: {
    save (raw) {
      var isValid = true
      this.$refs['editorForm'].validate((valid) => {
        if (!valid) {
          isValid = false
          return false
        }
      })
      if (!isValid) {
        return
      }
      // handle save call back
      if (this.isEditMode() || this.alreadyCreated) {
        this.update(raw)
        return
      }
      this.create(raw)
      this.alreadyCreated = true
    },
    create (raw) {
      this.article.content = raw
      this.article.content_type = 'editor_js'
      this.article.tags = this.getTags()
      let self = this
      createArticle(this.article).then(function (response) {
        self.article = response.data
        self.$message({
          message: 'Created successfully',
          type: 'success'
        })
      })
    },
    update (raw) {
      this.article.content = raw
      this.article.content_type = 'editor_js'
      this.article.tags = this.getTags()
      let id = this.id
      if (this.article !== undefined) {
        id = this.article.id
      }
      let self = this
      updateArticle(id, this.article).then(function (response) {
        self.$message({
          message: 'Update successfully',
          type: 'success'
        })
      })
    },
    isEditMode () {
      return this.mode !== 'new' || this.alreadyCreated
    },
    saveEditor () {
      // trigger editor save
      this.$refs.editor.save()
    },
    openReviewEditor () {
      // trigger editor preview
      this.$refs.editor.openReview()
    },
    getTags () {
      return this.$refs.tags.getTags()
    }
  }
}
</script>

<style lang="scss" scoped>
.article-editor {
  margin-top: 10px;

  .menu {
    margin-bottom: 30px;

    .menu-btn {
      float: right;
      margin-left: 10px;
    }
  }

  .form {
    .input-title {
      font-size: 2em;
      font-weight: bold;
    }

    .editorjs {
      border-radius: 4px;
      border: 1px solid #dcdfe6;
    }
  }
}
</style>
