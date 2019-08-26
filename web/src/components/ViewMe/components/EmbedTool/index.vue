<template>
    <div class="cdx-block">
        <iframe style="width:100%;" :height="data.height" :src="data.embed" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
        <div v-if="data.caption!== ''" class="caption" v-html="data.caption"></div>
    </div>
</template>

<script>
export default {
  props: {
    'data': Object
  }
}
</script>
