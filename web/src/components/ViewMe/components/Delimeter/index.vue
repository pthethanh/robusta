<template>
    <div class="ce-delimiter cdx-block"></div>
</template>

<script>
export default {
}
</script>

<style lang="scss" scoped>
.ce-delimiter {
    line-height: 1.6em;
    width: 100%;
    text-align: center;
}

.ce-delimiter:before {
    display: inline-block;
    content: "***";
    font-size: 30px;
    line-height: 65px;
    height: 30px;
    letter-spacing: 0.2em;
}
</style>
