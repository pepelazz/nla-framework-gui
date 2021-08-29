<template>
  <div class="row">
    <div class="col-10">
      <q-bar class="bg-teal text-white">
        <q-space />
        <q-btn dense flat icon="save"  @click="save"/>
        <q-btn dense flat icon="close" @click="close"/>
      </q-bar>
      <q-card class="shadow-1">
        <q-card-section>
          <section class="grid-stack grid-stack-4" style="margin-left: -10px"></section>
        </q-card-section>
      </q-card>
    </div>
  </div>
</template>

<style>
  .grid-stack-item-content {
    color: #2c3e50;
    text-align: center;
    background-color: #80cbc4;
    padding-top: 10px;
  }
</style>

<script>
  import {  ref, onMounted } from 'vue'
  import {GridStack} from 'gridstack'
  import 'gridstack/dist/gridstack.css'
  import 'gridstack/dist/gridstack-extra.css'
  import 'gridstack/dist/h5/gridstack-dd-native'
  import _ from 'lodash'
  export default {
    props: ['fldsByRows'],
    emits: ['close', 'update'],
    setup(props, {emit}) {
      let count = ref(0);
      let info = ref("");
      let grid = null;
      let items = []
      _.forEach(props.fldsByRows, (row, i) => {
        row.map((fld, j) => {
          let x = 0
          // суммируем ширину предыдущих элементов
          if (j > 0) {
            row.map((v, j1) => {
              if (j1 < j) x += row[j1].col_class.replace("col-", "")/2
            })
          }
          let y = i - 1
          let w = fld.col_class.replace("col-", "")/2
          // titleComputed не добавляем в сетку
          if (fld.func_name !== 'GetFldTitleComputed') {
            // noMove для поля title
            items.push({x, y, h: 1, w, content: fld.name || 'title', noMove:  !fld.name, locked:  !fld.name})
          }
        })
      })

      onMounted(() => {
        grid = GridStack.init({
          float: true,
          cellHeight: "70px",
          minRow: 1,
          column: 4,
          resizable: {handles: 'e, w'},
        });

        grid.load(items)

        // Use an arrow function so that `this` is bound to the Vue instance. Alternatively, use a custom Vue directive on the `.grid-stack` container element: https://vuejs.org/v2/guide/custom-directive.html
        // grid.on("dragstop", (event, element) => {
        //   const node = element.gridstackNode;
        //   // `this` will only access your Vue instance if you used an arrow function, otherwise `this` binds to window scope. see https://hacks.mozilla.org/2015/06/es6-in-depth-arrow-functions/
        //   console.log(`you just dragged node #${node.id} to ${node.x},${node.y} – good job!`);
        // });
      });

      function addNewWidget() {
        const node = items[count.value] || {
          x: Math.round(12 * Math.random()),
          y: Math.round(5 * Math.random()),
          w: Math.round(1 + 3 * Math.random()),
          h: Math.round(1 + 3 * Math.random()),
        };
        node.id = node.content = String(count.value++);
        grid.addWidget(node);
      }

      function save() {
          // модифицируем сетку перед сохранением
          // - прижимаем все блоки вправо и устраняя пробелы между ними
          // - удаляем пустые строки
          let modifyGrid = (items) => {
            let row = []
            let currentY = -1
            items.map(v => {
              // новая строка
              if (row.length === 0 || row[0].y !== v.y) {
                row = []
                currentY++
              }
              // проставляем номер строки по возрастанию, чтобы избежать пустых строк
              v.y = currentY
              // первый элемент
              if (row.length === 0) {
                // элемент не сначала строки, то сдвигаем в начало
                if (v.x !== 0) v.x = 0
              } else {
                // последующие элмменты прижимаем к предыдущим, если вдруг есть зазор
                v.x = row[row.length-1].x + row[row.length-1].w
              }
              // проставляем номер колонки
              v.col_pos = row.length + 1
              row.push(v)
            })
            return items
          }
        const items = modifyGrid(grid.save())
        // обновляем grid
        grid.removeAll()
        grid.load(items)

        // приводим данные к формату [][]int{{}} и col-...
        const flds = items.map(v => {
          return {name: v.content, row_col: [[v.y + 1, v.col_pos]], col_class: `col-${v.w * 2}`}
        })
        emit('update', flds)
      }

     function close() {
      emit('close')
     }

      return {
        info,
        addNewWidget,
        close,
        save,
      };
    },

    watch: {
      /**
       * Clear the info text after a two second timeout. Clears previous timeout first.
       */
      info: function (newVal) {
        if (newVal.length === 0) return;

        window.clearTimeout(this.timerId);
        this.timerId = window.setTimeout(() => {
          this.info = "";
        }, 2000);
      },
    },
  }
</script>
