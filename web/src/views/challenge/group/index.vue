<template>
  <div class="playground">
    <el-row type="flex" justify="center" v-if="ready && challenges.length > 0">
      <el-col :span="4" class="left">
        <el-menu default-active="0" v-if="ready">
          <el-menu-item v-for="(challenge,index) in  challenges" :key="challenge.id" @click="onClick(challenge)" :index="index + ''">{{index + 1 + '. ' + challenge.title}}</el-menu-item>
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
    <el-row type="flex" justify="center" v-if="ready && challenges.length == 0">
      <div>
        No Challenges Found
      </div>
    </el-row>
  </div>
</template>

<script>
import ChallengePlayer from '@/components/ChallengePlayer'
import {
  listChallenges
} from '@/api/challenge'
import {
  getFolder
} from '@/api/folder'

export default {
  components: {
    ChallengePlayer
  },
  data () {
    return {
      id: '',
      ready: false,
      offset: 0,
      limit: 50,
      selected: null,
      folder: null,
      challenges: [],
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    onClick (challenge) {
      this.selected = challenge
    },
    fetchData () {
      this.ready = false
      this.id = this.$route.params.id
      getFolder(this.id).then((response) => {
        this.folder = response.data
        listChallenges(this.getQueryStr()).then(response => {
          this.challenges = response.data
          if (this.challenges.length > 0) {
            this.selected = this.challenges[0]
          }
        })
      }).finally(() => {
        this.ready = true
      })
    },
    getQueryStr () {
      var ids = ''
      var children = this.folder.children
      for (var i = 0; i < children.length; i++) {
        ids += '&ids=' + children[i]
      }
      let query = 'offset=' + this.offset + '&limit=' + this.limit + ids
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
    top: 50px;
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
