<template>
  <div class="playground">
    <el-row type="flex" justify="center">
      <el-col :span="4" class="left">
        <el-menu default-active="1">
          <el-menu-item v-for="(challenge,index) in  challenges" :key="challenge.id" @click="onClick(challenge)">{{index + 1 + '. ' + challenge.title}}</el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="20">
        <div class="right">
          <div class="description" v-if="selected !== null">
            <div class="title">{{selected.title}}</div>
            <div class="content">
              {{selected.description}}
            </div>
          </div>
          <challenge-player v-if="selected !== null" :code="selected.sample" :challenge_id="selected.id" class="editor"></challenge-player>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import ChallengePlayer from '@/components/ChallengePlayer'
import {
  listChallenges
} from '@/api/challenge'

export default {
  components: {
    ChallengePlayer
  },
  data () {
    return {
      offset: 0,
      limit: 50,
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
  line-height: 1.5em;

  .left {
    border: 1px solid lightgrey;
    min-width: 200px;
    overflow-x: scroll;
  }

  .right {
    position: fixed;
    top: 10px;
    height: 100%;
    min-width: 70%;

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
