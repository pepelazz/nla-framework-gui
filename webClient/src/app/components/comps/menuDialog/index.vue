<template>
  <q-btn size="sm" flat round icon="menu" color="primary" @click="isShowDialog=true"><q-tooltip>редактировать меню</q-tooltip></q-btn>
  <q-dialog v-model="isShowDialog" position="left">
    <q-card style="width: 450px; max-width: 80vw;">
      <q-card-section>
        <comp-list :project="project" :menu="project.menu" @update="v => $emit('update', v)"/>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
  import {ref} from 'vue'
  import compList from './list'
  export default {
    components: {compList},
    props: ['project'],
    emits: ['update'],
    setup(props, {emit}) {
      const isShowDialog = ref(false)

      return {
        isShowDialog,
      }
    }
  }
</script>
