<template>
  <div class="home">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <Tutorial v-for="(tutorial,index) in tutorials" :key="index" :data="tutorial" @selected="openDetail">></Tutorial>
      </el-col>
    </el-row>
    <el-row v-if="isOpenDetail">
      <el-dialog :visible.sync="isOpenDetail" :center=true :modal=true :append-to-body=true :fullscreen=true @close="detailClosed">
        <TutorialDetail :data="selectedTutorial" @deleted="isOpenDetail=false"></TutorialDetail>
      </el-dialog>
    </el-row>
  </div>
</template>

<script>
import {
  fetchTutorials
} from '@/api/tutorial'
import Tutorial from '@/components/Tutorial'
import TutorialDetail from '@/components/TutorialDetail'
export default {
  components: {
    Tutorial,
    TutorialDetail
  },
  title: 'Goway - Learn to Go',
  data() {
    return {
      tutorials: [],
      selectedTutorial: null,
      isOpenDetail: false
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      fetchTutorials().then(response => {
        this.tutorials = response.data
      })
    },
    openDetail(data) {
      this.selectedTutorial = data
      this.isOpenDetail = true
      history.pushState({}, null, '/tutorials/detail/' + this.selectedTutorial.id)
    },
    detailClosed() {
      history.back()
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.home {
  margin: 0px;
  background-color: $mainBackgroundColor;
}
</style>

<style>
.el-dialog__body {
    padding: 0px;
}
.el-dialog--center .el-dialog__body {
    padding: 0px;
}
</style>
