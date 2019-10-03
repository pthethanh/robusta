<template>
  <div>
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <infinity-load :fetch-data="fetchData">
          <template v-slot:default="{item}">
            <div>
              {{item.title}}
            </div>
          </template>
          <template v-slot:loading>
            <div>
              Please wait...
            </div>
          </template>
        </infinity-load>
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
      offset: -1,
      limit: 1
    }
  },
  methods: {
    fetchData () {
      this.offset += this.limit
      return new Promise((resolve, reject) => {
        fetchList('offset=' + this.offset + '&limit=' + this.limit)
          .then(response => {
            this.data = response.data
            resolve(response)
          })
          .catch(response => {
            reject(response)
          })
      })
    }
  }
}
</script>
