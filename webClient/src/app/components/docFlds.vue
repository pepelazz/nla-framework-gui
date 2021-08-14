<template>

  <!-- название  -->
  <div class="row q-col-gutter-md q-mb-md">
    <div class="col">
      <div class="text-h6">{{doc.name}} / {{doc.name_ru}}</div>
      <q-separator/>
    </div>
  </div>

  <!-- поля -->
  <div class="row q-col-gutter-md q-mb-sm">
    <div class="col-auto">
      <div>
        <q-btn flat round icon="view_column" color="secondary" size="sm" @click="isShowDocGridEdit = true" :disable="isShowDocGridEdit">
          <q-tooltip>настройка сетки</q-tooltip>
        </q-btn>
      </div>
      <comp-add-new-fld :selectedDoc="selectedDoc" :project="projectLocal" @update="addNewFld"/>
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
                <div class="text-weight-medium">{{fld.name || 'title'}}
                  <!-- редактирование title. Возможность переключение между title и titleComputed    -->
                  <comp-title-fld-edit v-if="['GetFldTitle', 'GetFldTitleComputed'].includes(fld.func_name)" :fld="fld"/>

                </div>
              </q-card-section>
            </q-card-section>

            <q-separator />
            <q-card-actions style="padding-top: 2px; padding-bottom: 2px;">
              <comp-edit-fld v-if="fld.name" :fld="fld" class="q-ml-lg"/>
              <!-- иконки для GetFldRef -->
              <q-icon name="link" v-if="isShowLink(fld)" color="grey" class="q-mr-sm"/>
              <q-icon name="add" v-if="isAddNew(fld)" color="grey" class="q-mr-sm"/>
              <q-icon name="delete" v-if="isClearable(fld)" color="grey" class="q-mr-sm"/>

              <q-space/>
              <comp-delete-fld v-if="fld.name" :name="fld.name_ru" @remove="removeFld(fld)"/>
            </q-card-actions>
          </q-card>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
  // import _ from 'lodash'
  import docGridEdit from 'src/app/components/comps/docGridEdit'
  import {toRefs, computed, ref, watch} from 'vue'
  import fldsGroupByRows from 'src/app/components/docFuncs/fldsGroupByRows'
  import getIconByFldType from 'src/app/components/docFuncs/getIconByFldType'
  import docGridUpdateFunc from 'src/app/components/docFuncs/docGridUpdate'
  import compAddNewFld from 'src/app/components/comps/addNewFld'
  import compDeleteFld from 'src/app/components/comps/deleteFld'
  import compEditFld from 'src/app/components/comps/editFld'
  import compTitleFldEdit from 'src/app/components/comps/titleFldEdit'

  export default {
    props: ['selectedDoc', 'project'],
    emits: ['update'],
    components: {docGridEdit, compAddNewFld, compDeleteFld, compEditFld, compTitleFldEdit},
    setup(props, {emit}) {
      const doc = toRefs(props).selectedDoc
      const projectLocal = toRefs(props).project
      const fldsByRows = computed(() => fldsGroupByRows(doc.value.flds))
      const getIcon = getIconByFldType
      const isShowDocGridEdit = ref(false)

      const docGridUpdate = (flds) => docGridUpdateFunc(doc, flds)

      const isShowLink = computed(() => (fld) => fld.func_name === 'GetFldRef' && fld.params?.includes('isShowLink'))
      const isAddNew = computed(() => (fld) => fld.func_name === 'GetFldRef' && fld.params?.includes('isAddNew'))
      const isClearable = computed(() => (fld) => fld.func_name === 'GetFldRef' && fld.params?.includes('isClearable'))


      watch(doc, (v) => isShowDocGridEdit.value = false )
      // watch(doc, (v) => {
      //   emit('update', doc)
      // }, { deep: true } )

      const addNewFld = (v) => {
        doc.value.flds.push(v)
      }

      const removeFld = (fld) => {
        // TODO: добавить удаление колонки в postgres
        const i = doc.value.flds.findIndex(v => v.name === fld.name)
        doc.value.flds.splice(i, 1)
      }

      return {
        projectLocal,
        doc,
        fldsByRows,
        getIcon,
        isShowDocGridEdit,
        docGridUpdate,
        addNewFld,
        removeFld,
        isShowLink, isAddNew, isClearable,
      }
    },
  }

</script>

<style lang="sass" scoped>
  .my-card
    width: 100%
  /*max-width: 350px*/
</style>
