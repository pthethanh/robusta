<template>
  <div class="image">
    <el-image :src="data.file.url" lazy v-bind:class="{
      'image__stretched': data.stretched,
      'image__withBackground': data.withBackground,
      'image__withBorder': data.withBorder,
    }">
      <div slot="placeholder" class="image-slot">
        {{ $t('gen.loading') }}<span class="dot">...</span>
      </div>
      <div slot="error" class="image-slot">
        <i class="el-icon-picture-outline"></i>
      </div>
    </el-image>
    <div class="caption" v-if="data.caption !== ''" v-html="data.caption"></div>
  </div>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>

<style lang="scss" scoped>
.image {
  &__stretched {
    width: 100%;
  }
  &__withBorder {
    border: 1px solid lightgrey;
  }
  &__withBackground {
    text-align: center;
    background-color: grey;
  }
}
</style>
