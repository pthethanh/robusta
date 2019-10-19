<template>
  <div class="policy">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="24" :lg="22" :xl="18">
        <div class="title">Add Folder Policies</div>
        <el-form class="add-policy-form" :model="policy" @submit.native.prevent="addPolicy" :rules="rules" ref="add-policy-form" inline>
          <el-form-item prop="object">
            <el-select v-model="policy.object" placeholder="Select folder" filterable no-data-text="No data" no-match-text="No matching option" @change="reloadFolderPolicies">
              <infinity-load :fetch-data="fetchFolders" v-bind:data.sync="folders" @error="folderOffset -= folderLimit" no-more-text="" :delay="0" :limit="folderLimit">
                <template v-slot:default="{data}">
                  <el-option v-for="item in data" :key="item.id" :label="item.name" :value="item.id">
                  </el-option>
                </template>
              </infinity-load>
            </el-select>
          </el-form-item>
          <el-form-item prop="subject">
            <el-select v-model="policy.subject" placeholder="Select user" filterable no-data-text="No data" no-match-text="No matching option">
              <el-option v-for="item in users" :key="item.user_id" :label="getDisplayName(item)" :value="item.user_id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item prop="action">
            <el-select v-model="policy.action" placeholder="Select action" filterable no-data-text="No data" no-match-text="No matching option">
              <el-option v-for="item in folderActions" :key="item" :label="item" :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item prop="effect">
            <el-select v-model="policy.effect" placeholder="Select effect" filterable no-data-text="No data" no-match-text="No matching option">
              <el-option v-for="item in effect" :key="item" :label="item" :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="addPolicy">Add Policy</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="folderPolicies" style="width: 100%" empty-text="No data" max-height="500">
          <el-table-column label="User" prop="user" fixed min-width="250" sortable>
          </el-table-column>
          <el-table-column label="Action" prop="action" min-width="200" sortable>
          </el-table-column>
          <el-table-column label="Effect" prop="effect" min-width="200" sortable>
          </el-table-column>
          <el-table-column align="right" min-width="200">
            <template  slot-scope="scope">
              <el-button size="mini" type="danger" @click="handleRemovePolicy(scope.$index, scope.row)">Remove</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  listUsers
} from '@/api/user'
import {
  addPolicy,
  listPolicyActions,
  listPolicy,
  removePolicy
} from '@/api/policy'
import {
  listFolders
} from '@/api/folder'
import InfinityLoad from '@/components/InfinityLoad'
export default {
  components: {
    InfinityLoad
  },
  data () {
    return {
      users: [],
      folders: [],
      policy: {
        subject: '',
        object: '',
        action: '',
        effect: 'allow'
      },
      folderPolicies: [],
      folderActions: [],
      loading: false,
      effect: ['allow', 'deny'],
      error: '',
      rules: {
        subject: [{
          required: true,
          message: 'Please select user',
          trigger: 'change'
        }],
        object: [{
          required: true,
          message: 'Object is required',
          trigger: 'change'
        }],
        action: [{
          required: true,
          message: 'Action is required',
          trigger: 'change'
        }],
        effect: [{
          required: true,
          message: 'Effect is required',
          trigger: 'change'
        }]
      },
      folderOffset: -15,
      folderLimit: 15
    }
  },
  mounted () {
    listUsers().then((response) => {
      this.users.push({
        name: 'Any User',
        user_id: '*',
      })
      this.users = response.data
    })
    listPolicyActions().then((response) => {
      for (var i = 0; i < response.data.length; i++) {
        if (response.data[i].startsWith('folder:')) {
          this.folderActions.push(response.data[i])
        }
      }
    })
  },
  methods: {
    addPolicy () {
      this.loading = true
      var isValid = false
      this.$refs['add-policy-form'].validate((valid) => {
        isValid = valid
      })
      if (!isValid) {
        return
      }
      addPolicy(JSON.stringify(this.policy)).then((response) => {
        this.policy.user = this.getUserName(this.policy.subject)
        this.folderPolicies.unshift(this.policy)
        this.$message({
          message: 'Added policy successfully',
          type: 'success'
        })
      }).catch((error) => {
        this.$message({
          message: 'Failed to add policy: ' + error,
          type: 'error'
        })
      }).finally(() => {
        this.loading = false
      })
    },
    fetchFolders () {
      this.folderOffset += this.folderLimit
      return listFolders('offset=' + this.folderOffset + '&limit=' + this.folderLimit)
    },
    getDisplayName (user) {
      if (user.name === '') {
        return user.email
      }
      return user.name + '- ' + user.email
    },
    reloadFolderPolicies (folder) {
      this.folderPolicies.splice(0, this.folderPolicies.length)
      listPolicy('actions=folder&objects=' + folder).then((response) => {
        for (var i = 0; i < response.data.length; i++) {
          var rs = response.data[i]
          rs.user = this.getUserName(rs.subject)
          this.folderPolicies.push(rs)
        }
      }).catch((error) => {
        this.$message({
          message: 'Failed to load policies: ' + error,
          type: 'error'
        })
      })
    },
    getUserName (id) {
      for (var i = 0; i < this.users.length; i++) {
        if (this.users[i].user_id === id) {
          if (this.users[i].name !== undefined && this.users[i].name.trim() !== '') {
            return this.users[i].name
          }
          return this.users[i].email
        }
      }
      return ''
    },
    handleRemovePolicy (index, row) {
      var i = index
      removePolicy(JSON.stringify(this.folderPolicies[i])).then(() => {
        this.folderPolicies.splice(i, 1)
        this.$message({
          message: 'Removed policy successfully',
          type: 'success'
        })
      }).catch((error) => {
        this.$message({
          message: 'Failed remove policy: ' + error,
          type: 'error'
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.policy {
  .title {
    font-size: 1.5em;
    font-weight: 700;
  }
}
</style>
