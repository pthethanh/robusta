<template>
  <el-row type="flex" justify="center">
    <el-col :span="14">
      <dnd-list :list1.sync="list1" :list2.sync="list2" list1-title="Selected" list2-title="Available" :search="search">
        <template v-slot:list1="{data}">
          <div class="list-complete-item-handle">
            {{ data.title }}
          </div>
        </template>
        <template v-slot:list2="{data}">
          <div class="list-complete-item-handle2">
            {{ data.title }}
          </div>
        </template>
      </dnd-list>
    </el-col>
  </el-row>
</template>

<script>
import DndList from '@/components/DndList'
import {
  listChallenges
} from '@/api/challenge'
export default {
  name: 'DndListDemo',
  components: {
    DndList
  },
  data () {
    return {
      list1: [],
      list2: [],
      loading: false
    }
  },
  created () {
    this.getData()
  },
  methods: {
    getData () {
      this.listLoading = true
      listChallenges('').then(response => {
        this.list2 = response.data.splice(0, 5)
        this.list1 = response.data
      })
    },
    isNotInList1 (v) {
      return this.list1.every(k => v.id !== k.id)
    },
    search (keyword) {
      this.list2 = []
      listChallenges('title=' + keyword).then(response => {
        for (var i = 0; i < response.data.length; i++) {
          if (this.isNotInList1(response.data[i])) {
            this.list2.push(response.data[i])
          }
        }
      })
    }
  }
}
</script>
