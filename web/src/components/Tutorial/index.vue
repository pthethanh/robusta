<template>
  <div class="tutorial" @click="$emit('selected', data)">
    <el-card class="box-card">
      <el-row>
        <el-col :span="1.7" hidden>
          <div class="box box-left">
            <i class="el-icon-arrow-up"></i>
            <div>{{data.likes}}</div>
          </div>
        </el-col>
        <el-col>
          <div class="box">
            <div class="header">Posted by {{data.created_by_name}} {{data.created_at | date}}</div>
            <div class="title">{{data.steps[0].title}}</div>
            <div class="content">
              <view-me :data="data.steps[0].content"></view-me>
            </div>
            <ul class="footer">
              <li>{{data.views}} views |</li>
              <li>{{data.comments}} comments</li>
            </ul>
          </div>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
import ViewMe from '@/components/ViewMe'
export default {
  props: {
    'data': Object
  },
  components: {
    ViewMe
  },
  methods: {

  }
}
</script>

<style scoped>
.tutorial {
  margin: 5px 0;
}

.el-card:hover {
  border-color: black;
  cursor: pointer;
}

.el-card__body {
  padding-top: 0px;
  margin-top: 0px;
}

.el-card .box {
  float: left;
  max-width: 95%;
}

.box-left {
  width: 30px;
  font-weight: 700;
  font-family: 'Courier New', Courier, monospace;
  font-size: 1em;
  margin-right: 5px;
}

.header {
  margin-top: -10px;
  margin-bottom: 5px;
  font-family: 'Courier New', Courier, monospace;
  font-size: 0.8em;
  font-weight: lighter;
  color: grey;
}

.title {
  font-size: 1.1em;
  font-weight: 700;
  padding-bottom: 10px;
}

.content {
  max-height: 150px;
  overflow-y: hidden;
  margin-bottom: 10px;
}

ul.footer {
  padding-inline-start: 0px;
  display: inline;
  font-family: 'Courier New', Courier, monospace;
  font-size: 0.8em;
  font-weight: 700;
  color: grey;
}

.footer li {
  display: inline;
  margin-right: 10px;
}

.el-icon-arrow-up {
  font-size: 1.4em;
}

.el-icon-arrow-up:hover {
  color: red;
}
</style>
