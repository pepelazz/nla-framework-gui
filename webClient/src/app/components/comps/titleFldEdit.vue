<template>

  <span v-if="localFld.func_name === 'GetFldTitleComputed'" class="text-caption q-ml-sm">computed</span>
  <q-icon name="settings" color="grey" @click="showDialog" class="q-ml-sm cursor-pointer"><q-tooltip>редактировать</q-tooltip></q-icon>

  <q-dialog v-model="isShowDialog" persistent>
    <q-card style="min-width: 350px">
      <q-card-section>
        <div class="text-h6">Title</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <div class="q-gutter-sm">
          <q-radio v-model="localFld.func_name" val="GetFldTitle" label="title" />
          <q-radio v-model="localFld.func_name" val="GetFldTitleComputed" label="titleComputed" />
        </div>
      </q-card-section>

      <q-card-section class="q-pt-none" v-if="localFld.func_name === 'GetFldTitleComputed'">
        <!-- редактирование параметров для GetFldTitleComputed     -->
        <q-input  autogrow dense v-model="localFld.params[0]" autofocus outlined label="sql для формирования значения" @keyup.enter="isShowDialog=false"/>

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
  import editFldVueFilesParams from 'src/app/components/comps/comps/editFldVueFilesParams'
  import editFldVueImgParams from 'src/app/components/comps/comps/editFldVueImgParams'
  import {ref, toRefs, onMounted, onUpdated} from 'vue'
    export default {
      props: ['fld'],
      setup(props) {
        let initFld
        const isShowDialog = ref(false)
        const localFld = toRefs(props).fld

        const cancel = () => {
          localFld.value.name_ru = initFld.name_ru
          localFld.value.func_name = initFld.func_name
        }

        const showDialog = () => {
          isShowDialog.value = true
          initFld =Object.assign({}, props.fld)
        }

        onUpdated(() => {
          if (!localFld.value.params) localFld.value.params = []
        })

        return {
          isShowDialog,
          localFld,
          cancel,
          showDialog,
        }
      }
    }
</script>
