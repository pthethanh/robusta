<template>
  <div class="challenges-admin">
    <div class="loading" v-loading="loading"></div>
    <el-row type="flex" justify="center">
      <el-table :data="challenges.filter(data => !search || data.title.toLowerCase().includes(search.toLowerCase()))" style="width: 100%" empty-text="No data" max-height="500">
        <el-table-column label="Title" prop="title" fixed min-width="250">
        </el-table-column>
        <el-table-column label="Created By" prop="created_by_name" min-width="200">
        </el-table-column>
        <el-table-column label="Created At" prop="created_at" min-width="200">
        </el-table-column>
        <el-table-column label="Updated At" prop="updated_at" min-width="200">
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
        <infinity-load :fetch-data="fetchData" @error="offset -= limit" :data.sync="challenges" :limit="limit">
        </infinity-load>
      </el-table>
    </el-row>
  </div>
</template>

<script>
import {
  listChallenges,
  deleteChallenge
} from '@/api/challenge'
import InfinityLoad from '@/components/InfinityLoad'
export default {
  components: {
    InfinityLoad
  },
  data () {
    return {
      challenges: [],
      search: '',
      searchx: '',
      offset: -15,
      limit: 15
    }
  },
  methods: {
    fetchData () {
      return listChallenges(this.getQueryStr())
    },
    getQueryStr () {
      this.offset += this.limit
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
