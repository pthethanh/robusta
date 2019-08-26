<template>
  <div>
    <el-image :src="data.url" lazy>
      <div slot="placeholder" class="image-slot">
        Loading<span class="dot">...</span>
      </div>
    </el-image>
    <div class="cdx-input" v-if="data.caption !== ''" v-html="data.caption"></div>
  </div>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>
