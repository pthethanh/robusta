<template>
  <div class="playground">
    <el-row type="flex" justify="center">
      <el-col :span="4" class="left">
        <el-menu default-active="1" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
          <el-submenu v-for="(topic, index) in course.topics" :index="index + ''" :key="topic.id">
            <template slot="title">
              <span>{{topic.title}}</span>
            </template>
            <el-menu-item v-for="(ex, exIdx) in topic.exercises" :index="index + '-' + exIdx" :key="ex.id" @click="onClick(ex)">{{ex.title}}</el-menu-item>
          </el-submenu>
        </el-menu>
      </el-col>
      <el-col :span="20" class="right">
        <div v-if="selected === null" class="course-info">
          <div class="title"> {{course.title}}</div>
          <div class="content">{{course.introduction}}</div>
        </div>
        <div class="description" v-if="selected !== null">
          <div class="title">{{selected.title}}</div>
          <div class="content">
            {{selected.content}}
          </div>
        </div>
        <play-ground v-if="selected !== null" :code="selected.code" class="editor"></play-ground>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import PlayGround from '@/components/PlayGround'

export default {
  components: {
    PlayGround
  },
  data () {
    return {
      selected: null,
      course: {
        title: 'Go basic',
        introduction: 'Hello',
        topics: [{
          id: 1,
          title: 'Hello world',
          exercises: [{
            id: 1,
            title: 'Exercise 1',
            content: 'Print "Hello world" to standard output'
          }]
        },
        {
          id: 2,
          title: 'Variables',
          exercises: [{
            id: 1,
            title: 'Exercise 1',
            content: 'Declare x as float64 and init its value to 3.14'
          },
          {
            id: 2,
            title: 'Exercise 2',
            content: 'Print "Hello world" to standard output'
          }]
        }]
      }
    }
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    },
    onClick (ex) {
      this.selected = ex
    }
  }
}
</script>

<style lang="scss" scoped>
.playground {
  height: 100%;
  line-height: 1.5em;

  .left {
    height: 100%;
    border: 1px solid lightgrey;
  }

  .right {
    padding-top: 5px;
    height: 100%;

    .course-info {
      .title {
        font-size: 2.5em;
        font-family: 'Open Sans', sans-serif;
        word-break: normal;
        font-weight: 700;
      }

      .introduction {
        font-family: "Merriweather", serif;
        letter-spacing: 0.01rem;
        font-size: 0.95rem;
        line-height: 1.75em;
        color: rgba(0, 0, 0, 0.84);
      }
    }

    .description {
      padding: 10px 25px 20px 25px;
      font-family: 'Open Sans', sans-serif;
      word-break: normal;

      .title {
        font-size: 1.4em;
        font-weight: 700;
      }

      .content {
        font-family: 'Open Sans', sans-serif;
        line-height: 1.5em;
      }
    }
  }
}
</style>
