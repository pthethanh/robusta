<template>
  <div class="article">
    <el-card
      class="box-card">
        <el-row>
          <el-col :span="1.7">
            <div class="box box-left">
              <i class="el-icon-arrow-up"></i>
              <div>{{article.likes}}</div>
            </div>
          </el-col>
          <el-col :span="22">
            <div class="box">
              <div class="header">Posted by {{article.created_by_name}} 20 hours ago</div>
              <div class="title">{{article.title}}</div>
              <div class="content" v-html="article.abstract"></div>
              <ul class="footer">
                <li>{{article.views}} views |</li>
                <li>{{article.comments}} comments</li>
              </ul>
            </div>
          </el-col>
        </el-row>
    </el-card>
  </div>
</template>

<script>
export default {
  props: {
    'article': Object
  }
}
</script>

<style scoped>
  .article {
    margin: 20px 0;
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
  .footer li{
    display: inline;
    margin-right: 10px;
  }
  .el-icon-arrow-up {
    font-size: 1.4em;
  }
  .el-icon-arrow-up:hover{
    color: red;
  }
</style>
