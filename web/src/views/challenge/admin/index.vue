<template>
  <div class="challenges-admin">
    <div class="loading" v-loading="loading"></div>
    <el-row type="flex" justify="center">
      <el-table :data="challenges.filter(data => !search || data.title.toLowerCase().includes(search.toLowerCase()))" style="width: 100%" empty-text="No data" max-height="500">
        <div v-infinite-scroll="fetchData" infinite-scroll-disabled="disabled">
          <el-table-column label="Title" prop="title" fixed min-width="250">
          </el-table-column>
          <el-table-column label="Created By" prop="created_by_name" min-width="200">
          </el-table-column>
          <el-table-column label="Created At" prop="created_at" min-width="200">
          </el-table-column>
          <el-table-column align="right" min-width="200">
            <template slot="header">
              <el-input v-model="search" size="mini" placeholder="Type to search" />
            </template>
            <template slot-scope="scope">
              <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
              <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
            </template>
          </el-table-column>
        </div>
      </el-table>
    </el-row>
  </div>
</template>

<script>
import {
  listChallenges,
  deleteChallenge
} from '@/api/challenge'
export default {
  data () {
    return {
      challenges: [],
      search: '',
      searchx: '',
      loading: false,
      noMoreData: false,
      offset: 0,
      limit: 15
    }
  },
  computed: {
    disabled () {
      return this.loading || this.noMore
    },
    noMore () {
      return this.noMoreData === true
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.loading = true
      listChallenges(this.getQueryStr()).then(response => {
        this.loading = false
        if (response.data === null || response.data.length === 0) {
          this.noMoreData = true
          return
        }
        this.challenges = this.challenges.concat(response.data)
        this.offset += response.data.length
      }).finally(() => {
        this.loading = false
      })
    },
    getQueryStr () {
      let query = 'offset=' + this.offset + '&limit=' + this.limit
      return query
    },
    handleEdit (index, row) {
      this.$router.push('/challenges/edit/' + this.challenges[index].id)
    },
    handleDelete (index, row) {
      this.$confirm('This will delete this challenge permanently. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        deleteChallenge(this.challenges[index].id).then((respose) => {
          this.$message({
            type: 'success',
            message: 'Delete successfully'
          })
          this.challenges.splice(index, 1)
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.challenges-admin {
  margin-top: 20px;
}
</style>
