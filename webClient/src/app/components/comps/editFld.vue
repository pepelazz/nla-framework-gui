<template>
  <q-btn flat round icon="edit" color="grey" size="sm" @click="showDialog"><q-tooltip>редактировать</q-tooltip></q-btn>

  <q-dialog v-model="isShowDialog" persistent>
    <q-card style="min-width: 350px">
      <q-card-section>
        <div class="text-h6">{{localFld.name || 'TitleComputed'}}</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-input v-if="localFld.func_name !== 'GetFldTitleComputed'" dense v-model="localFld.name_ru" autofocus outlined label="название" @keyup.enter="isShowDialog=false"/>


        <edit-fld-vue-options-items v-if="['GetFldSelectString', 'GetFldSelectMultiple', 'GetFldRadioString'].includes(localFld.func_name)" :options="localFld.fld_vue_options_item"/>
        <edit-fld-vue-files-params v-if="localFld.func_name === 'GetFldFiles'" :params="localFld.fld_vue_files_params"/>
        <edit-fld-vue-img-params v-if="['GetFldImg', 'GetFldImgList'].includes(localFld.func_name)" :params="localFld.fld_vue_img_params"/>

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
  import {ref, toRefs, onMounted} from 'vue'
    export default {
      props: ['fld'],
      components: {editFldVueOptionsItems, editFldVueFilesParams, editFldVueImgParams},
      setup(props) {
        let initFld
        const isShowDialog = ref(false)
        const localFld = toRefs(props).fld

        // добавляем дефолтный объект fld_vue_files_params
        if (localFld.value.func_name === 'GetFldFiles' && !localFld.value.fld_vue_files_params) {
          localFld.value.fld_vue_files_params = {Accept: null, MaxFileSize: null}
        }
        // добавляем дефолтный объект fld_vue_img_params
        if (['GetFldImg', 'GetFldImgList'].includes(localFld.func_name)  && !localFld.value.fld_vue_img_params) {
          localFld.value.fld_vue_img_params = {Accept: null, MaxFileSize: null, Crop: null, Width: null, CanAddUrls: false}
        }
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
          cancel,
          showDialog,
        }
      }
    }
</script>
