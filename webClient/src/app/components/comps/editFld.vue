<template>
  <q-btn flat round icon="edit" color="grey" size="sm" @click="showDialog"><q-tooltip>редактировать</q-tooltip></q-btn>

  <q-dialog v-model="isShowDialog" persistent>
    <q-card style="min-width: 350px">
      <q-card-section>
        <div class="text-h6">{{localFld.name}}</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-input dense v-model="localFld.name_ru" autofocus outlined label="название" @keyup.enter="isShowDialog=false"/>

        <edit-fld-vue-options-items v-show="localFld.func_name === 'GetFldSelectString'" :options="localFld.fld_vue_options_item"/>

      </q-card-section>

      <q-card-actions align="right" class="text-primary">
        <q-btn flat label="Отмена" v-close-popup @click="cancel"/>
        <q-btn flat label="Сохранить" v-close-popup/>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
  import editFldVueOptionsItems from 'src/app/components/comps/comps/editFldVueOptionsItems'
  import {ref, toRefs, onMounted} from 'vue'
    export default {
      props: ['fld'],
      components: {editFldVueOptionsItems},
      setup(props) {
        let initFld
        const isShowDialog = ref(false)
        const localFld = toRefs(props).fld
        const ref_EditFldVueOptionsItems = ref(null)
        const cancel = () => {
          localFld.value.name_ru = initFld.name_ru
        }

        const showDialog = () => {
          isShowDialog.value = true
          initFld =Object.assign({}, props.fld)
        }

        return {
          isShowDialog,
          localFld,
          ref_EditFldVueOptionsItems,
          cancel,
          showDialog,
        }
      }
    }
</script>
