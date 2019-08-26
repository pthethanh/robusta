<template>
    <div>
        <el-row type="flex" justify="center">
            <el-col :span="14">
                <view-me :data="data"></view-me>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import ViewMe from '@/components/ViewMe'
import * as configs from './play.js'

export default {
  components: {
    ViewMe
  },
  data () {
    return {
      data: configs.playground
    }
  }
}
</script>
