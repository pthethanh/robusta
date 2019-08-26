<template>
  <!-- we are using this kind of table instead of el-table because it can contains some HTML inside the row values -->
  <div class="tc-editor cdx-block">
    <div class="tc-table__wrap">
        <table class="tc-table">
            <tbody>
              <tr v-for="(col, index) in data.content" :key="index">
                <td class="tc-table__cell" v-for="(row, rowIndex) in col" :key="rowIndex">
                  <div class="tc-table__area">
                    <div class="tc-table__inp" v-html="row"></div>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
      </div>
  </div>
</template>

<script>
export default {
  props: {
    'data': Object
  },
  methods: {
    numColumns () {
      return this.data.content[0].length
    },
    tableData () {
      return this.data.content
    }
  }
}
</script>

<style lang="scss" scoped>
.tc-table {
  width: 100%;
  height: 100%;
  border-collapse: collapse;
  table-layout: fixed;

  &__wrap {
    border: 1px solid #DBDBE2;
    border-radius: 3px;
    position: relative;
    height: 100%;
    width: 100%;
    box-sizing: border-box;
  }

  &__cell {
    border: 1px solid #DBDBE2;
    padding: 0;
    vertical-align: top;
  }

  &__area {
    padding: 10px;
    height: 100%;
  }

  &__inp {
    outline: none;
    flex-grow: 100;
    min-height: 1.5em;
    height: 100%;
    overflow: hidden;
  }

  tbody tr:first-child td {
    border-top: none;
  }

  tbody tr:last-child td {
    border-bottom: none;
  }

  tbody tr td:last-child {
    border-right: none;
  }

  tbody tr td:first-child {
    border-left: none;
  }
}
</style>
