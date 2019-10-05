<template>
  <div>
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <infinity-load :fetch-data="fetchData" v-bind:data.sync="data" @error="offset = offset - limit">
          <template v-slot:default="{data}">
            <div v-for="(item, index) in filterData" :key="index">
              <div>
                {{item.title}} -- {{item.created_by_name}}
              </div>
            </div>
          </template>
        </infinity-load>
        <el-button @click="logData">Log Data</el-button>
      </el-col>
    </el-row>
  </div>

</template>

<script>
import InfinityLoad from '@/components/InfinityLoad'
import {
  fetchList
} from '@/api/article'
export default {
  components: {
    InfinityLoad
  },
  data () {
    return {
      data: [],
      offset: -10,
      limit: 10,
      tags: []
    }
  },
  computed: {
    filterData() {
      var rs = []
      for (var i = 0; i < this.data.length; i++) {
        if (this.tags.length == 0) {
          rs.push(this.data[i])
          continue
        }
        for (var j = 0; j < this.tags.length; j++) {
          if (this.data[i].tags.includes(this.tags[j])) {
            rs.push(this.data[i])
            break
          }
        }
      }
      return rs
    }
  },
  methods: {
    fetchData () {
      this.offset += this.limit
      return fetchList('offset=' + this.offset + '&limit=' + this.limit)
    },
    logData () {
      console.log(this.data)
      if (this.tags.length == 0) {
        this.tags.push('interview preparation')
        return
      }
      this.tags =[]
    }
  }
}
</script>
