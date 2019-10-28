<template>
  <div class="playground">
    <el-row>
      <div v-if="ready && challenges.length == 0" class="error">
        {{ $t('challenge.folder_no_challenges') }}
      </div>
    </el-row>
    <split-pane :min-percent='5' :default-percent='20' split="vertical" v-loading="_loading" :element-loading-text="$t('gen.loading')" element-loading-background="rgba(0, 0, 0, 0.7)">
      <template slot="paneL">
        <el-menu default-active="0" class="left" v-if="ready">
          <el-menu-item v-for="(challenge,index) in  challenges" :key="challenge.id" @click="onClick(challenge)" :index="index + ''">
            <i class="el-icon-check challenge-completed" v-if="challenge.completed"></i>
            <el-badge value="new" class="badge" v-if="challenge.is_new">
              <span>{{index + 1 + '. ' + challenge.title}}</span>
            </el-badge>
            <span v-if="!challenge.is_new">{{index + 1 + '. ' + challenge.title}}</span>
          </el-menu-item>
        </el-menu>
      </template>
      <template slot="paneR" class="right">
        <el-tabs type="border-card" v-model="activeTab" @tab-click="handleTabClick" v-if="ready">
          <el-tab-pane :label="$t('challenge.detail')" name="detail" style="height: 100vh; overflow: scroll;">
            <div class="description" v-if="selected !== null">
              <div class="title">{{selected.title}}</div>
              <view-me :data="selected.description" class="description"></view-me>
            </div>
            <challenge-player v-if="selected !== null" :code="selected.sample" :challenge-id="selected.id" :folder-id="folder.id" class="editor" @run-completed="handlePlayerRunCompleted"></challenge-player>
          </el-tab-pane>
          <el-tab-pane :label="$t('challenge.submissions')" name="submissions" style="height: 100vh; overflow: scroll;">
            <el-table :data="submissions" style="width: 100%" empty-text="No sucess submission found" v-loading="loadingSolution">
              <el-table-column prop="created_by_name" :label="$t('challenge.name')">
              </el-table-column>
              <el-table-column prop="created_at_date" :label="$t('challenge.date')">
              </el-table-column>
              <el-table-column prop="status" :label="$t('challenge.status')">
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane :label="$t('challenge.discussions')" name="discussions" style="height: 100vh; overflow: scroll;">
            <comments :targetID="selected.id" targetType='challenge' class="comments"></comments>
          </el-tab-pane>
        </el-tabs>
      </template>
    </split-pane>
  </div>
</template>

<script>
import variables from '@/styles/variables.scss'
import ChallengePlayer from '@/components/ChallengePlayer'
import ViewMe from '@/components/ViewMe'
import SplitPane from 'vue-splitpane'
import Comments from '@/components/Comments'
import {
  listChallenges
} from '@/api/challenge'
import {
  getFolder
} from '@/api/folder'
import {
  completionReport
} from '@/api/solution'
import {
  mapGetters
} from 'vuex'
export default {
  components: {
    ChallengePlayer,
    ViewMe,
    SplitPane,
    Comments
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
      submissions: [],
      activeTab: 'detail',
      loadingSolution: true
    }
  },
  mounted () {
    this.fetchData()
  },
  computed: {
    _loading () {
      return !this.ready
    },
    _submissions () {
      return this.submissions
    },
    ...mapGetters([
      'user'
    ]),
    variables () {
      return variables
    }
  },
  methods: {
    onClick (challenge) {
      this.selected = challenge
      if (this.activeTab === 'submissions') {
        this.handleOpenSubmissionsTab()
      }
    },
    async fetchData () {
      this.ready = false
      this.id = this.$route.params.id
      getFolder(this.id).then((response) => {
        this.folder = response.data
        listChallenges(this.getQueryStr()).then(response => {
          var data = response.data
          for (var i = 0; i < data.length; i++) {
            data[i].completed = false // populate property for Vuejs watch
            this.challenges.push(data[i])
          }
          if (this.challenges.length <= 0) {
            return
          }
          this.selected = this.challenges[0]
          // call async to load completion
          this.updateChallengeCompletion()
        }).finally(() => {
          this.ready = true
        })
      }).catch(() => {
        this.ready = true
      })
    },
    async fetchSubmissions () {
      this.loadingSolution = true
      completionReport('status=success&challenge_ids=' + this.selected.id).then((response) => {
        this.submissions = response.data
        for (var i = 0; i < this.submissions.length; i++) {
          this.submissions[i].created_at_date = this.submissions[i].created_at.substring(0, 10)
          if (this.submissions[i].created_by_name === undefined) {
            this.submissions[i].created_by_name = 'Gopher'
          }
        }
      }).finally(() => {
        this.loadingSolution = false
      })
    },
    getQueryStr () {
      var ids = ''
      var children = this.folder.children
      for (var i = 0; i < children.length; i++) {
        ids += '&ids=' + children[i]
      }
      let query = 'folder_id=' + this.folder.id + '&offset=' + this.offset + '&limit=' + this.limit + ids + '&sort_by=-created_at'
      return query
    },
    handleTabClick (tab, event) {
      if (this.activeTab === 'submissions') {
        this.handleOpenSubmissionsTab()
      }
    },
    handleOpenSubmissionsTab () {
      this.fetchSubmissions()
    },
    async updateChallengeCompletion () {
      if (!this.user.authenticated) {
        return
      }
      var q = 'include_detail=true&created_by_id=' + this.user.info.user_id
      var children = this.folder.children
      for (var i = 0; i < children.length; i++) {
        q += '&challenge_ids=' + children[i]
      }
      completionReport(q).then((response) => {
        var completed = response.data
        for (var i = 0; i < this.challenges.length; i++) {
          for (var j = 0; j < completed.length; j++) {
            if (this.challenges[i].id === completed[j].challenge_id) {
              if (completed[j].status === 'success') {
                this.challenges[i].completed = true
                this.challenges[i].sample = completed[j].content
              } else {
                this.challenges[i].completed = false
                this.challenges[i].sample = completed[j].content
              }
            }
          }
        }
      })
    },
    handlePlayerRunCompleted (passed, code) {
      this.selected.completed = passed
      this.selected.sample = code
    }
  }
}
</script>

<style lang="scss" scoped>
.el-tab-pane {
  padding-bottom: 200px;
}

.playground {
  line-height: 1.5em;
  height: 100vh;
  width: 100vw;
  position: fixed;

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
    height: 100vh;
    overflow-y: scroll;
    padding-bottom: 50px;

    .badge {
      display: inline;
    }

    .challenge-completed {
      color: green;
      font-weight: 700;
    }
  }

  .right {
    right: 20px;
    padding-left: 20px; // same with right above

    .description {
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
