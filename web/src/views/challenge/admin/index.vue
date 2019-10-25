<template>
  <div class="challenges-admin">
    <div class="loading" v-loading="loading"></div>
    <el-row type="flex" justify="center">
      <el-table :data="challenges.filter(data => !search || data.title.toLowerCase().includes(search.toLowerCase()))" style="width: 100%" empty-text="No data" max-height="500">
        <el-table-column :label="$t('challenge.title')" prop="title" fixed min-width="250">
        </el-table-column>
        <el-table-column :label="$t('challenge.created_by')" prop="created_by_name" min-width="200">
        </el-table-column>
        <el-table-column :label="$t('challenge.created_at')" prop="created_at" min-width="200">
        </el-table-column>
        <el-table-column :label="$t('challenge.updated_at')" prop="updated_at" min-width="200">
        </el-table-column>
        <el-table-column align="right" min-width="200">
          <template slot="header">
            <el-input v-model="search" size="mini" :placeholder="$t('challenge.type_to_search')" />
          </template>
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">{{ $t('gen.edit') }}</el-button>
            <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">{{ $t('gen.delete') }}</el-button>
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
      this.$confirm(this.$i18n.t('challenge.delete_confirm'), this.$i18n.t('gen.warning'), {
        confirmButtonText: this.$i18n.t('gen.ok'),
        cancelButtonText: this.$i18n.t('gen.cancel'),
        type: 'warning'
      }).then(() => {
        deleteChallenge(this.challenges[index].id).then((respose) => {
          this.$message({
            type: 'success',
            message: this.$i18n.t('gen.delete_success')
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
