<template>
  <div class="turorial">
    <el-row type="flex" justify="center">
      <el-col :span="20">
        <el-row type="flex" justify="center">
          <el-col :xs="24" :sm="24" :md="{span: 18, offset: 4}" :lg="{span: 14, offset: 4}" :xl="{span: 14, offset: 4}" class="menu">
            <el-button type="primary" size="mini" class="el-icon-check menu-btn" @click="save" circle></el-button>
            <el-button type="primary" size="mini" class="el-icon-plus menu-btn" @click="addNewSection" circle></el-button>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center" class="full-height" v-if="ready">
      <el-col :span="4" v-show="sidebarActive" class="sidebar-wrapper hidden-sm-and-down">
        <el-row>
          <el-menu default-active="2" :collapse="false" class="sidebar">
            <el-steps :space="50" direction="vertical" :active="index" finish-status="success">
              <el-step v-for="(step, i) in data.steps" :key="i" :title="step.title" @click.native="goToStep(i, step)"></el-step>
            </el-steps>
          </el-menu>
        </el-row>
      </el-col>
      <el-col class="content" :span="20">
        <el-row type="flex" justify="center">
          <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
            <div v-for="(step, i) in data.steps" :key="i">
              <div v-if="i===index">
                <el-input class="input-title" placeholder="title" v-model="step.title"></el-input>
                <editor ref="editor" @save="saveSection" :initData="step.content" class="editorjs"></editor>
              </div>
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-button-group style="bottom: 20px; position: fixed; width: 100%">
            <el-button type="info" size="mini" icon="el-icon-arrow-left" style="positon: fixed; left: 10px;" @click="prevStep()"></el-button>
            <el-button type="info" size="mini" icon="el-icon-arrow-right" style="position: fixed; right: 10px" @click="nextStep()"></el-button>
          </el-button-group>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Editor from '@/components/Editor'
import {
  createTutorial,
  updateTutorial,
  fetchTutorial
} from '@/api/tutorial'
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
      data: {
        id: '',
        steps: [{
          title: 'Introduction',
          content: {}
        }]
      },
      index: 0,
      sidebarActive: true,
      alreadyCreated: false,
      ready: false
    }
  },
  created () {
    let self = this
    if (this.isEditMode()) {
      fetchTutorial(this.id).then(response => {
        this.data = response.data
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
    this.$title = this.data.title
  },
  methods: {
    goToStep (i, step) {
      this.step = step
      this.index = i
    },
    nextStep () {
      var index = this.index + 1
      if (index > this.data.steps.length - 1) {
        index = this.data.steps.length - 1
      }
      this.index = index
      this.step = this.data.steps[this.index]
    },
    prevStep () {
      var index = this.index - 1
      if (index < 0) {
        index = 0
      }
      this.index = index
      this.step = this.data.steps[this.index]
    },
    toggleSideBar () {
      this.sidebarActive = !this.sidebarActive
    },
    addNewSection () {
      this.data.steps.push({
        title: 'Section ' + this.data.steps.length,
        content: {}
      })
    },
    triggerEditorSave () {
      // trigger the editor save
      this.$refs.editor[0].save()
    },
    saveSection (raw) {
      // listen value from the editor and save to the current section
      this.data.steps[this.index].content = raw
      // once we got data from callback, save it to db
      if (this.isEditMode()) {
        let self = this
        updateTutorial(this.data.id, this.data).then(function (response) {
          self.$message({
            message: 'Updated successfully',
            type: 'success'
          })
        })
        return
      }
      var self = this
      createTutorial(this.data).then(function (response) {
        self.data.id = response.data.id
        self.$message({
          message: 'Created successfully',
          type: 'success'
        })
      })
      this.alreadyCreated = true
    },
    save () {
      this.triggerEditorSave()
    },
    isEditMode () {
      return this.mode !== 'new' || this.alreadyCreated
    }
  }
}
</script>

<style lang="scss" scoped>
.turorial {
  background-color: white;
  height: 100%;

  .menu {
    margin-top: 10px;
    margin-bottom: 10px;

    .menu-btn {
      float: right;
      margin-left: 10px;
    }
  }

  .sidebar-wrapper {
    overflow-y: scroll;
    max-height: 100vh;

    .sidebar {
      top: 10px;
      left: 10px;
      height: 100%;
      border-right: 1px solid lightgray;
      color: black;
      font-weight: bold;
      background-color: whitesmoke;
    }
  }

  .content {
    overflow-y: scroll;
    max-height: 100vh;
    display: block;
    top: 20px;

    .input-title {
      font-size: 2em;
      font-weight: bold;
    }
  }
}
</style>

<style lang="scss">
.el-step__title {
  &.is-success {
    color: green;
  }

  &:hover {
    cursor: pointer;
  }
}
</style>
