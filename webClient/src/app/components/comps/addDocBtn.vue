<template>
  <q-btn icon="add" @click="isShowDialog=true" size="sm" flat round color="primary"><q-tooltip>добавить таблицу</q-tooltip></q-btn>
  <q-dialog v-model="isShowDialog" persistent position="top">
    <q-card style="min-width: 350px">
      <q-card-section>
        <div class="text-h6">Добавить таблицу</div>
      </q-card-section>
      <q-card-section>
        <q-input v-model="name" label="name" dense/>
      </q-card-section>
      <q-card-section>
        <q-input v-model="name_ru" label="название (рус)" dense/>
      </q-card-section>
      <q-card-section>
        <q-input v-model="name_ru_plural" label="название (рус. множественное)" dense/>
      </q-card-section>
      <q-card-actions align="right" class="text-primary">
        <q-btn flat label="Отмена" v-close-popup/>
        <q-btn flat label="Ок" @click="add"/>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
  import {ref} from 'vue'
  import {Notify} from 'quasar'
  import $utils from '../../plugins/utils'
  export default {
    setup() {
      const isShowDialog = ref(false)
      const name = ref(null)
      const name_ru = ref(null)
      const name_ru_plural = ref(null)

      const add = () => {
        if (!name.value || name.value.length < 4) {
          Notify.create({type: 'negative', message: 'поле name не должно быть меньше 4х символов'})
          return
        }
        if (!name_ru.value) {
          Notify.create({type: 'negative', message: 'поле "название" не заполнено'})
          return
        }
        if (!name_ru_plural.value) {
          Notify.create({type: 'negative', message: 'поле "название множественное" не заполнено'})
          return
        }
        $utils.postApiRequest({
          url: "/addDoc",
          params: {name: name.value, name_ru: name_ru.value, name_ru_plural: name_ru_plural.value}
        }).subscribe(res => {
          console.log('/addDoc res', res)
        })
      }

      return {
        isShowDialog,
        name, name_ru, name_ru_plural,
        add,
      }
    }
  }
</script>
