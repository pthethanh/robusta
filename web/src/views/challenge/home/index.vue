<template>
  <div class="folders">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <el-card v-for="folder in folders" :key="folder.id" class="folder" @click="goToFolder(folder)">
          <div class="name" @click="goToFolder(folder)">{{folder.name}} </div>
          <div class="description">{{folder.description}}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  listFolders
} from '@/api/folder'
export default {
  data () {
    return {
      offset: 0,
      limit: 50,
      folders: []
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.folders = []
      listFolders(this.getQueryStr()).then((response) => {
        this.folders = response.data
      })
    },
    getQueryStr () {
      let query = 'offset=' + this.offset + '&limit=' + this.limit + '&type=challenge'
      return query
    },
    goToFolder (folder) {
      this.$router.push('/challenges/groups/' + folder.id)
    }
  }
}
</script>

<style lang="scss" scoped>
.folders {
  margin-top: 20px;

  .folder {
    margin: 5px 0px;

    .name {
      font-size: 1.1em;
      font-weight: 700;
      padding-bottom: 10px;
    }

    .description {
      max-height: 250px;
      overflow-y: hidden;
      margin-bottom: 5px;
    }
  }
}
</style>
