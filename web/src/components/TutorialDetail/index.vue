<template>
  <div class="turorial">
    <el-row type="flex" justify="center" v-if="user.info.user_id === data.created_by_id">
      <el-col :span="20">
        <el-row type="flex" justify="center">
          <el-col :xs="24" :sm="24" :md="{span: 18, offset: 4}" :lg="{span: 14, offset: 4}" :xl="{span: 14, offset: 4}" class="menu">
            <el-button type="success" size="mini" class="el-icon-edit menu-btn" @click="goToEditPage" circle></el-button>
            <el-button type="danger" size="mini" class="el-icon-delete menu-btn" @click="deleteTutorial" circle></el-button>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center" class="full-height">
      <el-col :span="4" v-show="sidebarActive" class="sidebar-wrapper hidden-sm-and-down">
        <el-menu default-active="2" :collapse="false" class="sidebar">
          <el-steps :space="50" direction="vertical" :active="index" finish-status="success">
            <el-step v-for="(step, i) in data.steps" :key="i" :title="step.title" @click.native="goToStep(i, step)"></el-step>
          </el-steps>
        </el-menu>
      </el-col>
      <el-col class="content" :span="20">
        <el-row type="flex" justify="center">
          <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
            <div v-for="(s, i) in data.steps" :key="i" :data="step.content">
              <div v-if="i===index">
                <h1 class="title">{{step.title}}</h1>
                <view-me :data="step.content" v-if="step.content !== undefined"></view-me>
                <p v-if="step.content === undefined">....</p>
              </div>
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-button-group style="bottom: 20px; position: fixed; width: 100%">
            <el-button type="info" size="mini" icon="el-icon-arrow-left" style="positon: fixed; left: 10px;" @click="prevStep()"></el-button>
            <el-button type="info" size="mini" icon="el-icon-arrow-right" style="position: fixed; right: 10px" @click="nextStep()"></el-button>
            <!-- Not sure why cannot assign the shortkey to the el-button, hence do with hidden button temporary -->
            <button v-shortkey="['arrowright']" @shortkey="nextStep()" hidden></button>
            <button v-shortkey="['arrowleft']" @shortkey="prevStep()" hidden></button>
          </el-button-group>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import ViewMe from '@/components/ViewMe'
import {
  deleteTutorial,
  viewTutorial
} from '@/api/tutorial'
import {
  mapGetters
} from 'vuex'

export default {
  components: {
    ViewMe
  },
  props: {
    'data': Object
  },
  data() {
    return {
      step: {},
      index: 0,
      sidebarActive: true
    }
  },
  computed: {
    ...mapGetters([
      'user'
    ])
  },
  mounted() {
    if (this.data.steps !== undefined && this.data.steps.length > 0) {
      this.step = this.data.steps[0]
      this.index = 0
    }
  },
  title: 'Go tutorials',
  created() {
    viewTutorial(this.data.id)
    this.$title = this.data.steps[0].title
  },
  methods: {
    goToStep(i, step) {
      this.step = step
      this.index = i
    },
    nextStep() {
      var index = this.index + 1
      if (index > this.data.steps.length - 1) {
        index = this.data.steps.length - 1
      }
      this.index = index
      this.step = this.data.steps[this.index]
    },
    prevStep() {
      var index = this.index - 1
      if (index < 0) {
        index = 0
      }
      this.index = index
      this.step = this.data.steps[this.index]
    },
    toggleSideBar() {
      this.sidebarActive = !this.sidebarActive
    },
    goToEditPage() {
      this.$router.push('/tutorials/edit/' + this.data.id)
    },
    deleteTutorial() {
      let self = this
      this.$confirm('Are you sure you want to delete this tutorial permanently. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        deleteTutorial(this.data.id).then(function (respose) {
          self.$message({
            type: 'success',
            message: 'Delete successfully'
          })
          self.$router.push('/tutorials')
        })
      })
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
    // overflow-y: scroll;
    // max-height: 100vh;

    .sidebar {
      top: 20px;
      left: 10px;
      height: 100%;
      border-right: 1px solid lightgray;
      background-color: whitesmoke;
    }
  }

  .content {
    // overflow-y: scroll;
    // max-height: 100vh;
    display: block;
    margin-top: -25px;
  }
}
</style>

<style lang="scss">
.el-step__title {
  &.is-success {
    color: black;
  }

  &:hover {
    cursor: pointer;
  }
}
</style>
