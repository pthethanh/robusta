<template>
  <div class="folders">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <infinity-load :fetch-data="fetchData" @error="offset -= limit" :limit="limit">
          <template v-slot:default="{data}">
            <el-card v-for="folder in data" :key="folder.id" class="folder" @click="goToFolder(folder)">
              <div class="name" @click="goToFolder(folder)">{{folder.name}} </div>
              <div class="description" @click="goToFolder(folder)">{{folder.description}}</div>
            </el-card>
          </template>
        </infinity-load>
      </el-col>
    </el-row>
  </div>
</template>

<script>
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
      offset: -15,
      limit: 15
    }
  },
  methods: {
    fetchData () {
      return listFolders(this.getQueryStr())
    },
    getQueryStr () {
      this.offset += this.limit
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
@import '@/styles/variables.scss';

.folders {
  min-height: 100vh;
  background-color: $backgroundColorPrimary;

  .folder {
    margin: 5px 0px;

    .name {
      font-size: 1.1em;
      font-weight: 700;
      padding-bottom: 10px;
      color: $fontColorHeading;
    }

    .description {
      max-height: 250px;
      overflow-y: hidden;
      margin-bottom: 5px;
      color: $fontColorContent;
    }
  }

  .loading {
    text-align: center;
    font-weight: 600;
  }
}
</style>
