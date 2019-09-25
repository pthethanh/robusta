<template>
  <div class="policy">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <div class="form">
          <div class="title">Add Folder Policies</div>
          <el-form class="add-policy-form" :model="policy" @submit.native.prevent="addPolicy">
            <el-select v-model="policy.subject" placeholder="Select">
              <el-option v-for="item in users" :key="item.value" :label="item.email" :value="item.user_id">
              </el-option>
            </el-select>
            <el-select v-model="policy.object" placeholder="Select">
              <el-option v-for="item in folders" :key="item.value" :label="item.name" :value="item.id">
              </el-option>
            </el-select>
            <el-select v-model="policy.action" placeholder="Select">
              <el-option v-for="item in folderActions" :key="item" :label="item" :value="item">
              </el-option>
            </el-select>
            <el-select v-model="policy.effect" placeholder="Select">
              <el-option v-for="item in effect" :key="item" :label="item" :value="item">
              </el-option>
            </el-select>
            <el-form-item>
              <el-button :loading="loading" type="primary" @click="addPolicy">Add Policy</el-button>
            </el-form-item>
          </el-form>
        </div>
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
  listPolicyActions
} from '@/api/policy'
import {
  listFolders
} from '@/api/folder'
export default {
  data () {
    return {
      users: [],
      folders: [],
      policy: {},
      folderActions: [],
      loading: false,
      effect: ['allow', 'deny']
    }
  },
  mounted () {
    listUsers().then((response) => {
      this.users = response.data
    })
    listPolicyActions().then((response) => {
      for (var i = 0; i < response.data.length; i++) {
        if (response.data[i].startsWith('folder:')) {
          this.folderActions.push(response.data[i])
        }
      }
    })
    listFolders('offset=0&limit=100&type=challenge').then((response) => {
      this.folders = response.data
    })
  },
  methods: {
    addPolicy () {
      this.loading = true
      console.log(JSON.stringify(this.policy))
      addPolicy(JSON.stringify(this.policy)).then((response) => {
        this.$message({
          message: 'Created successfully',
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
    }
  }
}
</script>
