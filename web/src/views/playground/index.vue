<template>
  <div>
    <el-row type="flex" justify="center">
      <el-col :span="14">
        <play-ground :code="code"></play-ground>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import PlayGround from '@/components/PlayGround'

export default {
  components: {
    PlayGround
  },
  data () {
    return {
      code: `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
`
    }
  }
}
</script>
