<template>
  <div class="playground">
    <el-row type="flex" justify="left" v-if="ready && challenges.length > 0">
      <el-col :span="5" class="left">
        <el-menu default-active="0" v-if="ready">
          <el-menu-item v-for="(challenge,index) in  challenges" :key="challenge.id" @click="onClick(challenge)" :index="index + ''">{{index + 1 + '. ' + challenge.title}}</el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="19" class="right">
        <el-tabs type="border-card">
          <el-tab-pane label="Detail">
            <div>
              <div class="description" v-if="selected !== null">
                <div class="title">{{selected.title}}</div>
                <div class="content">
                  {{selected.description}}
                </div>
              </div>
              <challenge-player v-if="selected !== null" :code="selected.sample" :challenge_id="selected.id" class="editor"></challenge-player>
            </div>
          </el-tab-pane>
          <el-tab-pane label="Submissions">
            List of submissions. This feature will be coming very soon.
          </el-tab-pane>
          <el-tab-pane label="Tips">
            Articles, tips. This feature will be coming very soon.
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <div v-loading="_loading" v-if="!ready" class="loading">Loading...</div>
      <div v-if="ready && challenges.length == 0" class="error">
        There is no challenges in this folder.
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
      challenges: []
    }
  },
  mounted () {
    this.fetchData()
  },
  computed: {
    _loading () {
      return !this.ready
    }
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
        }).finally(() => {
          this.ready = true
        })
      }).catch((error) => {
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

  .loading {
    text-align: center;
    margin-top: 10px;
  }

  .error {
    margin-top: 10px;
    font-size: 1.2em;
    font-weight: 700;
  }

  .left {
    border: 1px solid lightgrey;
    overflow: scroll;
  }

  .right {
    position: fixed;
    height: 100%;
    right: 20px;
    padding-left: 20px; // same with right above

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
