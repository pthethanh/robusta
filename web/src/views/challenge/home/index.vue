<template>
  <div class="folders">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" v-infinite-scroll="fetchData" infinite-scroll-disabled="disabled">
        <el-card v-for="folder in folders" :key="folder.id" class="folder" @click="goToFolder(folder)">
          <div class="name" @click="goToFolder(folder)">{{folder.name}} </div>
          <div class="description" @click="goToFolder(folder)">{{folder.description}}</div>
        </el-card>
        <div v-if="loading" class="loading">Loading...</div>
        <div v-if="!loading&&noMore" class="loading">¯\_(ツ)_/¯</div>
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
      loading: false,
      noMoreData: false,
      offset: 0,
      limit: 50,
      folders: []
    }
  },
  created () {
    this.fetchData()
  },
  computed: {
    disabled () {
      return this.loading || this.noMore
    },
    noMore () {
      return this.noMoreData === true
    }
  },
  methods: {
    fetchData () {
      this.loading = true
      listFolders(this.getQueryStr()).then((response) => {
        if (response.data === null || response.data.length === 0) {
          this.noMoreData = true
          return
        }
        this.folders = this.folders.concat(response.data)
        this.offset += response.data.length
        if (response.data.length < this.limit) {
          this.noMoreData = true
        }
      }).finally(() => {
        this.loading = false
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
  height: 100%;

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

  .loading {
    text-align: center;
    font-weight: 600;
  }
}
</style>
