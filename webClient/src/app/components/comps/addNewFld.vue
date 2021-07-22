<template>
  <q-btn flat round icon="add" color="secondary" size="sm" @click="isShowDialogAddFld = true">
    <q-tooltip>добавить поле</q-tooltip>
  </q-btn>
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
    import {ref} from 'vue'
    export default {
        setup() {
          const isShowDialogAddFld = ref(false)
          const newFld = ref({label: 'строка', value: 'GetFldString'})

          const optionsFldType = [
            {label: 'строка', value: 'GetFldString'},
            {label: 'дата', value: 'GetFldDate'},
            {label: 'число целое', value: 'GetFldInt64'},
            {label: 'число дробное', value: 'GetFldInt64'},
          ]

          const addNewFld = () => {
            console.log('addNewFld', newFld)
          }

          return {
            isShowDialogAddFld,
            optionsFldType,
            newFld,
            addNewFld,
          }
        }
    }
</script>

