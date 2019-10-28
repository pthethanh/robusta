<template>
  <div class="challenge-editor">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" v-if="ready">
        <div class="menu">
          <el-button type="primary" size="mini" class="menu-btn" @click="saveEditor">{{ $t('gen.save') }}</el-button>
          <el-button type="primary" size="mini" class="menu-btn" @click="openReviewEditor">{{ $t('gen.preview') }}</el-button>
        </div>
        <el-form :model="challenge" :rules="rules" ref="editorForm" class="form">
          <el-form-item prop="title">
            <el-input class="input-title" :placeholder="$t('challenge.title')" v-model="challenge.title"></el-input>
          </el-form-item>
          <editor ref="editor" @save="save" class="editorjs" :initData="challenge.description"></editor>
          <el-form-item prop="sample">
            <el-input type="textarea" :rows="10" :placeholder="$t('challenge.code_sample')" v-model="challenge.sample">
            </el-input>
          </el-form-item>
          <el-form-item prop="test">
            <el-input type="textarea" :rows="15" :placeholder="$t('challenge.test')" v-model="challenge.test">
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
          message: this.$i18n.t('validation.title_required'),
          trigger: 'blur'
        }, {
          max: 256,
          message: this.$i18n.t('validation.title_max_256'),
          trigger: 'blur'
        }],
        sample: [{
          required: true,
          message: this.$i18n.t('challenge.validation.sample_required'),
          trigger: 'blur'
        }],
        test: [{
          required: true,
          message: this.$i18n.t('challenge.validation.test_required'),
          trigger: 'blur'
        }]
      }
    }
  },
  title: '',
  created () {
    if (this.isEditMode()) {
      getChallenge(this.id).then(response => {
        this.challenge = response.data
        this.ready = true
      }).catch((err) => {
        this.$message({
          message: this.$i18n.t('gen.load_data_failed') + ': ' + err,
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
          message: this.$i18n.t('gen.create_success'),
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
          message: this.$i18n.t('gen.update_success'),
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
    }
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
