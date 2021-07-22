<template>

  <div class="row q-col-gutter-md q-mb-md">
    <div class="col">
        <div class="text-h6">{{doc.name}} / {{doc.name_ru}}</div>
        <q-separator/>
    </div>
  </div>
  <div class="row q-col-gutter-md q-mb-sm">
    <div class="col-auto">
      <div>
      <q-btn flat round icon="add" color="secondary" size="sm" @click="isShowDialogAddFld = true">
        <q-tooltip>добавить поле</q-tooltip>
      </q-btn>
      </div>
      <div>
      <q-btn flat round icon="view_column" color="secondary" size="sm" @click="isShowDocGridEdit = true" :disable="isShowDocGridEdit">
        <q-tooltip>настройка сетки</q-tooltip>
      </q-btn>
      </div>
    </div>
    <div class="col">
      <doc-grid-edit v-if="isShowDocGridEdit" :fldsByRows="fldsByRows" class="q-mb-md" @update="docGridUpdate" @close="isShowDocGridEdit = false"/>
      <div class="row q-col-gutter-md q-mb-sm" v-for="row in fldsByRows">
        <div :class="fld.col_class" v-for="fld in row" :key="fld.id">
          <q-card class="my-card" bordered>
            <q-card-section horizontal>
              <q-card-section class="q-pa-none">
                <q-avatar>
                  <q-icon :name="getIcon(fld)" color="primary"/>
                </q-avatar>
              </q-card-section>
              <q-card-section class="q-pa-none">
                <div class="text-weight-regular">{{fld.name_ru || 'название'}}</div>
                <div class="text-weight-medium">{{fld.name || 'title'}}</div>
              </q-card-section>
            </q-card-section>

            <q-separator />
          </q-card>
        </div>
      </div>
    </div>
  </div>
  <q-dialog v-model="isShowDialogAddFld" position="top">
    <q-card style="width: 450px">
      <q-card-section class="q-pb-none">
        <div class="text-h6">Добавить поле</div>
      </q-card-section>
      <q-card-section>
        <q-select v-model="newFld" :options="optionsFldType" label="тип поля" />
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="OK" color="primary" @click="addNewFld"/>
        <q-btn flat label="Отмена" color="secondary"  v-close-popup/>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
  import _ from 'lodash'
  import docGridEdit from 'src/app/components/docGridEdit'

  export default {
    props: ['selectedDoc', 'project'],
    emits: ['update'],
    components: {docGridEdit},
    computed: {
      fldsByRows() {
        const res = _.chain(this.doc.flds)
          // группируем по строкам
          .groupBy(v => v.row_col ? v.row_col[0][0] : 1)
          // внутри строки сортируем по колонкам
          .forEach(row => row.sort((a, b) => {
            return (a.row_col ? a.row_col[0][1] : 1) - (b.row_col ? b.row_col[0][1] : 1)
          }))
          .value()
        return res
      }
    },
    data() {
      return {
        doc: {},
        isShowDialogAddFld: false,
        isShowDocGridEdit: false,
        newFld: null,
        optionsFldType: [
          {label: 'строка', value: 'GetFldString'},
          {label: 'дата', value: 'GetFldDate'},
          {label: 'число целое', value: 'GetFldInt64'},
          {label: 'число дробное', value: 'GetFldInt64'},
        ]
      }
    },
    watch: {
      selectedDoc(v) {
        this.doc = this.selectedDoc
        this.isShowDocGridEdit = false
      },
      doc: {
        handler() {
          this.$emit('update', this.doc)
        },
        deep: true
      }
    },
    methods: {
      getIcon(fld){
        switch (fld.func_name) {
          case 'GetFldTitle':
            return 'title'
          case 'GetFldRef':
            return 'link'
          case 'GetFldFiles':
            return 'file_present'
          case 'GetFldTag':
            return 'sell'
          case 'GetFldDate':
            return 'event'
          case 'GetFldDateTime':
            return 'schedule'
          case 'GetFldCheckbox':
            return 'check_circle_outline'
          case 'GetFldEmail':
            return 'email'
          case 'GetFldPhone':
            return 'phone'
          default:
            return 'text_fields'
        }
      },

      // обновляем сетку - col_class и row_col
      docGridUpdate(flds) {
        flds.map(v => {
          if (v.name === 'title') {
            this.doc.flds[0].col_class = v.col_class
          } else {
            let fld = this.doc.flds.find(v1 => v1.name === v.name)
            if (fld) {
              fld.col_class = v.col_class
              fld.row_col = v.row_col
            }
          }
        })
      },
      // добавление нового поля
      addNewFld() {
        console.log('addNewFld:', this.newFld)
      }
    },
    mounted() {
      this.doc = this.selectedDoc
    }
  }

</script>

<style lang="sass" scoped>
  .my-card
    width: 100%
    /*max-width: 350px*/
</style>
