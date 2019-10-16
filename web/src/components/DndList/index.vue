<template>
  <div class="dnd-list">
    <div :style="{width:width1}" class="dnd-list-list">
      <h3>{{ list1Title }}</h3>
      <draggable :set-data="setData" :list="list1" group="item" class="drag-area">
        <div v-for="element in list1" :key="element.id" class="list-complete-item">
          <div class="list-complete-item-handle">
            <slot name="list1" :data="element">
              {{ element.id }}
            </slot>
          </div>
          <div style="position:absolute;right:0px;">
            <span style="float: right ;margin-top: -20px;margin-right:5px;" @click="deleteEle(element)">
              <i style="color:#ff4949" class="el-icon-delete" />
            </span>
          </div>
        </div>
      </draggable>
    </div>
    <div :style="{width:width2}" class="dnd-list-list">
      <h3>{{ list2Title }}</h3>
      <el-input v-if="search !== undefined" placeholder="Type to search" prefix-icon="el-icon-search" v-model="keyword" v-on:keyup.enter="searchKeyword(keyword)">
        <el-button :loading="loading" slot="append" icon="el-icon-search" @click="searchKeyword(keyword)"></el-button>
      </el-input>
      <draggable :list="list2" group="item" class="drag-area2">
        <div v-for="element in list2" :key="element.id" class="list-complete-item">
          <div class="list-complete-item-handle2" @click="pushEle(element)">
            <slot name="list2" :data="element">
              {{ element.id }}
            </slot>
          </div>
        </div>
      </draggable>
    </div>
  </div>
</template>

<script>
import draggable from 'vuedraggable'
export default {
  name: 'DndList',
  components: {
    draggable
  },
  props: {
    list1: {
      type: Array,
      default () {
        return []
      }
    },
    list2: {
      type: Array,
      default () {
        return []
      }
    },
    list1Title: {
      type: String,
      default: 'Selected'
    },
    list2Title: {
      type: String,
      default: 'Available'
    },
    width1: {
      type: String,
      default: '48%'
    },
    width2: {
      type: String,
      default: '48%'
    },
    search: {
      type: Function
    }
  },
  data () {
    return {
      keyword: '',
      loading: false
    }
  },
  methods: {
    isNotInList1 (v) {
      return this.list1.every(k => v.id !== k.id)
    },
    isNotInList2 (v) {
      return this.list2.every(k => v.id !== k.id)
    },
    deleteEle (ele) {
      for (const item of this.list1) {
        if (item.id === ele.id) {
          const index = this.list1.indexOf(item)
          this.list1.splice(index, 1)
          break
        }
      }
      if (this.isNotInList2(ele)) {
        this.list2.unshift(ele)
      }
      this.$emit('update:list1', this.list1)
      this.$emit('update:list2', this.list2)
    },
    pushEle (ele) {
      for (const item of this.list2) {
        if (item.id === ele.id) {
          const index = this.list2.indexOf(item)
          this.list2.splice(index, 1)
          break
        }
      }
      if (this.isNotInList1(ele)) {
        this.list1.push(ele)
      }
      this.$emit('update:list1', this.list1)
      this.$emit('update:list2', this.list2)
    },
    setData (dataTransfer) {
      // to avoid Firefox bug
      // Detail see : https://github.com/RubaXa/Sortable/issues/1012
      dataTransfer.setData('Text', '')
    },
    searchKeyword (keyword) {
      this.loading = true
      this.search(keyword).then(response => {
        this.list2.splice(0, this.list2.length)
        for (var i = 0; i < response.data.length; i++) {
          if (this.isNotInList1(response.data[i])) {
            this.list2.push(response.data[i])
          }
        }
      }).finally(() => {
        this.loading = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dnd-list {
  background: #fff;
  padding-bottom: 40px;

  &:after {
    content: "";
    display: table;
    clear: both;
  }

  .dnd-list-list {
    float: left;
    padding-bottom: 30px;

    &:first-of-type {
      margin-right: 2%;
    }

    .drag-area {
      margin-top: 15px;
      min-height: 50px;
      padding-bottom: 30px;
      max-height: 450px;
      overflow-y: auto;
    }

    .drag-area2 {
      margin-top: 15px;
      min-height: 50px;
      padding-bottom: 30px;
      max-height: 394px;
      overflow-y: auto;
    }
  }
}

.list-complete-item {
  cursor: pointer;
  position: relative;
  font-size: 14px;
  padding: 5px 12px;
  margin-top: 4px;
  border: 1px solid #bfcbd9;
  transition: all 1s;
}

.list-complete-item-handle {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-right: 50px;
}

.list-complete-item-handle2 {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-right: 20px;
}

.list-complete-item.sortable-chosen {
  background: #4AB7BD;
}

.list-complete-item.sortable-ghost {
  background: #30B08F;
}

.list-complete-enter,
.list-complete-leave-active {
  opacity: 0;
}
</style>
