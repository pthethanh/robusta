<template>
  <div class="viewme">
    <div v-for="block in blocks" :key="block.id">
      <code-ex v-if="block.type==='codeExample'" :code="block.data.code" :sample="block.data.sample">
      </code-ex>
      <Header v-if="block.type==='header'" :data="block.data">
      </Header>
      <check-list v-if="block.type==='checklist'" :data="block.data">
      </check-list>
      <Delimeter v-if="block.type==='delimiter'">
      </Delimeter>
      <embed-tool v-if="block.type==='embed'" :data="block.data">
      </embed-tool>
      <simple-image v-if="block.type==='simpleImage'" :data="block.data">
      </simple-image>
      <link-tool v-if="block.type==='linkTool'" :data="block.data" :mode="mode">
      </link-tool>
      <Raw v-if="block.type==='raw'" :data="block.data">
      </Raw>
      <paragraph v-if="block.type==='paragraph'" :data="block.data">
      </paragraph>
      <quote-tool v-if="block.type==='quote'" :data="block.data">
      </quote-tool>
      <table-tool v-if="block.type==='table'" :data="block.data">
      </table-tool>
      <Warning v-if="block.type==='warning'" :data="block.data">
      </Warning>
      <List v-if="block.type==='list'" :data="block.data">
      </List>
      <code-tool v-if="block.type==='code'" :data="block.data">
      </code-tool>
      <image-tool v-if="block.type==='image'" :data="block.data">
      </image-tool>
    </div>
  </div>
</template>

<script>
import CodeEx from './components/CodeEx'
import Header from './components/Header'
import CheckList from './components/CheckList'
import Delimeter from './components/Delimeter'
import EmbedTool from './components/EmbedTool'
import SimpleImage from './components/SimpleImage'
import LinkTool from './components/LinkTool'
import Paragraph from './components/Paragraph'
import QuoteTool from './components/QuoteTool'
import Raw from './components/Raw'
import TableTool from './components/TableTool'
import Warning from './components/Warning'
import List from './components/List'
import CodeTool from './components/CodeTool'
import ImageTool from './components/ImageTool'

export default {
  components: {
    CodeEx,
    Header,
    CheckList,
    Delimeter,
    EmbedTool,
    SimpleImage,
    LinkTool,
    Paragraph,
    QuoteTool,
    Raw,
    TableTool,
    Warning,
    List,
    CodeTool,
    ImageTool
  },
  props: {
    'data': {
      type: Object,
      default: () => {},
      required: true
    },
    'mode': {
      type: String,
      default: () => {},
      required: false
    }
  },
  data () {
    return {

    }
  },
  computed: {
    blocks: function () {
      if (this.mode !== 'preview') {
        return this.data.blocks
      }
      if (this.data.blocks.length === 0) {
        return this.data.blocks
      }
      // display minimum as much as possible to improve performance
      var firstType = this.data.blocks[0].type
      var oneLineFirstTypes = ['image', 'simpleImage', 'linkTool', 'warning', 'embed', 'quote', 'table']
      for (var i = 0; i < oneLineFirstTypes.length; i++) {
        if (oneLineFirstTypes[i] === firstType) {
          return this.data.blocks.slice(0, 1)
        }
      }
      return this.data.blocks.slice(0, 2)
    }
  }
}
</script>

<style lang="scss" scoped>
.viewme {
  word-break: keep-all;

  ::-webkit-scrollbar {
    display: none;
  }
}
</style>
