<template>
  <div class="playground">
    <el-row type="flex" justify="center">
      <el-col :span="4" class="left">
        <el-menu default-active="1" class="el-menu-vertical-demo" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
          <el-menu-item v-for="challenge in  challenges" :key="challenge.id" @click="onClick(challenge)">{{challenge.title}}</el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="20" class="right">
        <div class="description" v-if="selected !== null">
          <div class="title">{{selected.title}}</div>
          <div class="content">
            {{selected.description}}
          </div>
        </div>
        <play-ground v-if="selected !== null" :code="selected.sample" class="editor"></play-ground>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import PlayGround from '@/components/PlayGround'
import {
  listChallenges
} from '@/api/challenge'

export default {
  components: {
    PlayGround
  },
  data () {
    return {
      offset: 0,
      limit: 15,
      selected: null,
      challenges: []
    }
  },
  mounted () {
    listChallenges(this.getQueryStr()).then(response => {
      this.challenges = response.data
    })
  },
  methods: {
    onClick (challenge) {
      this.selected = challenge
    },
    getQueryStr () {
      let query = 'offset=' + this.offset + '&limit=' + this.limit
      return query
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
