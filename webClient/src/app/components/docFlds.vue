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
      <add-new-fld/>
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

</template>

<script>
  // import _ from 'lodash'
  import docGridEdit from 'src/app/components/comps/docGridEdit'
  import {toRefs, computed, ref, watch, watchEffect} from 'vue'
  import fldsGroupByRows from 'src/app/components/docFuncs/fldsGroupByRows'
  import getIconByFldType from 'src/app/components/docFuncs/getIconByFldType'
  import docGridUpdateFunc from 'src/app/components/docFuncs/docGridUpdate'
  import addNewFld from 'src/app/components/comps/addNewFld'

  export default {
    props: ['selectedDoc', 'project'],
    emits: ['update'],
    components: {docGridEdit, addNewFld},
    setup(props, {emit}) {
      const doc = toRefs(props).selectedDoc
      const fldsByRows = computed(() => fldsGroupByRows(doc.value.flds))
      const getIcon = getIconByFldType
      const isShowDocGridEdit = ref(false)

      const docGridUpdate = (flds) => docGridUpdateFunc(doc, flds)


      watch(doc, (v) => isShowDocGridEdit.value = false )
      // watch(doc, (v) => {
      //   emit('update', doc)
      // }, { deep: true } )

      return {
        doc,
        fldsByRows,
        getIcon,
        isShowDocGridEdit,
        docGridUpdate,
      }
    },
  }

</script>

<style lang="sass" scoped>
  .my-card
    width: 100%
  /*max-width: 350px*/
</style>
