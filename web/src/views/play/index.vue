<template>
  <el-row type="flex" justify="center">
    <el-col :span="14">
      <dnd-list :list1="list1" :list2="list2" list1-title="Selected" list2-title="Available">
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
  fetchList
} from '@/api/article'
export default {
  name: 'DndListDemo',
  components: {
    DndList
  },
  data () {
    return {
      list1: [],
      list2: []
    }
  },
  created () {
    this.getData()
  },
  methods: {
    getData () {
      this.listLoading = true
      fetchList('').then(response => {
        this.list2 = response.data.splice(0, 5)
        this.list1 = response.data
      })
    }
  }
}
</script>
