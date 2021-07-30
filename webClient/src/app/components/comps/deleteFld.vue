<template>
  <q-btn flat round icon="delete" color="grey" size="sm" @click="isShowDialog=true"><q-tooltip>удалить</q-tooltip></q-btn>
  <q-dialog v-model="isShowDialog">
    <q-card style="min-width: 350px">
      <q-card-section class="row items-center">
        <q-avatar icon="warning" color="warning" text-color="white" />
        <span class="q-ml-sm">Удалить поле: <b>"{{name}}"</b> ?</span>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat label="Отмена" color="secondary" v-close-popup />
        <q-btn flat label="Удалить" color="red" @click="removeFld" v-close-popup />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
    import {ref} from 'vue'
    export default {
      props: ['name'],
      emits: ['remove'],
      setup(props, {emit}) {
        const isShowDialog = ref(false)
        const removeFld = () => emit('remove')
        return {
          isShowDialog,
          removeFld,
        }
      }
    }
</script>

