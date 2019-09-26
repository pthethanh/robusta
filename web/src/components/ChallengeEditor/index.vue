<template>
  <div class="challenge-editor">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" v-if="ready">
        <div class="menu">
          <el-button type="primary" size="mini" class="menu-btn" @click="saveEditor">Save</el-button>
          <el-button type="primary" size="mini" class="menu-btn" @click="openReviewEditor">Preview</el-button>
        </div>
        <el-form :model="challenge" :rules="rules" ref="editorForm" class="form">
          <el-form-item prop="title">
            <el-input class="input-title" placeholder="title" v-model="challenge.title"></el-input>
          </el-form-item>
          <editor ref="editor" @save="save" class="editorjs" :initData="challenge.description"></editor>
          <el-form-item prop="sample">
            <el-input type="textarea" :rows="10" placeholder="Code sample" v-model="challenge.sample">
            </el-input>
          </el-form-item>
          <el-form-item prop="test">
            <el-input type="textarea" :rows="15" placeholder="Unit test" v-model="challenge.test">
            </el-input>
          </el-form-item>
          <el-form-item prop="tips">
            <el-input type="textarea" :rows="15" placeholder="Tips" v-model="challenge.tips">
            </el-input>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Editor from '@/components/Editor'
import {
  createChallenge,
  getChallenge,
  updateChallenge
} from '@/api/challenge'

export default {
  components: {
    Editor
  },
  props: {
    'id': String,
    'mode': String
  },
  data () {
    return {
      ready: false, // only render component if everything is ready
      alreadyCreated: false,
      challenge: {},
      rules: {
        title: [{
          required: true,
          message: 'Please enter title',
          trigger: 'blur'
        }, {
          max: 256,
          message: 'Length should be less than 256 characters',
          trigger: 'blur'
        }],
        sample: [{
          required: true,
          message: 'Please enter sample',
          trigger: 'blur'
        }],
        test: [{
          required: true,
          message: 'Please enter test',
          trigger: 'blur'
        }]
      }
    }
  },
  title: '',
  created () {
    let self = this
    if (this.isEditMode()) {
      getChallenge(this.id).then(response => {
        this.challenge = response.data
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
    this.$title = this.challenge.title
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
      this.challenge.description = raw
      this.challenge.content_type = 'editor_js'
      createChallenge(this.challenge).then((response) => {
        self.challenge = response.data
        this.$message({
          message: 'Created successfully',
          type: 'success'
        })
      })
    },
    update (raw) {
      this.challenge.description = raw
      this.challenge.content_type = 'editor_js'
      let id = this.id
      if (this.challenge !== undefined) {
        id = this.challenge.id
      }
      updateChallenge(id, this.challenge).then((response) => {
        this.$message({
          message: 'Update successfully',
          type: 'success'
        })
      })
    },
    isEditMode () {
      return this.mode !== 'new' || this.alreadyCreated
    },
    saveEditor () {
      this.$refs.editor.save()
    },
    openReviewEditor () {
      this.$refs.editor.openReview()
    },
  }
}
</script>

<style lang="scss" scoped>
.challenge-editor {
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
