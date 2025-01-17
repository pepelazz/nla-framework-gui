<template>
  <q-list bordered separator dense>
    <draggable v-model="localMenu" item-key="id" :group="{ name: 'g1' }" style="min-height: 30px;">
      <template #item="{element}">
        <q-item>
          <q-item-section avatar @click="edit(element)" class="cursor-pointer">
            <q-avatar rounded>
              <img :src="element.Icon || docNameIcon(element.DocName)">
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label>{{element.DocName ? docNameTitle(element.DocName) : element.Text}}</q-item-label>
            <!-- вложенное меню  -->
            <nested-draggable v-if="element.IsFolder" :menu="element.LinkList" :project="project" @update="v => element.LinkList = v"/>
          </q-item-section>
          <q-item-section side>
            <q-btn icon="delete" color="grey" flat round size="sm" @click="remove(element.id)"/>
          </q-item-section>
        </q-item>
      </template>
    </draggable>
    <q-item>
      <q-item-section side>
        <q-btn icon="add" flat round size="sm" @click="isShowAddDialog=true"><q-tooltip>добавить пункт меню</q-tooltip></q-btn>
      </q-item-section>
    </q-item>
  </q-list>

  <!-- Диалог добавления -->
  <q-dialog v-model="isShowAddDialog" position="left">
    <q-card style="width: 450px; max-width: 80vw;">
      <q-card-section>
        <div class="q-gutter-md" style="max-width: 300px">
          <q-select label="таблица" :options="options" v-model="newItem.DocName"/>
          <q-input label="url" v-model="newItem.Url"/>
          <q-input label="текст" v-model="newItem.Text"/>
          <q-input label="иконка" v-model="newItem.Icon"/>
          <q-checkbox label="вложенное меню" v-model="newItem.IsFolder"/>
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="OK" color="primary" @click="add" />
      </q-card-actions>
    </q-card>
  </q-dialog>

  <!-- Диалог редактирования -->
  <q-dialog v-model="isShowEditDialog" position="left">
    <q-card style="width: 450px; max-width: 80vw;">
      <q-card-section>
        <div class="q-gutter-md" style="max-width: 300px">
          <q-select label="таблица" :options="options" v-model="editItem.DocName"/>
          <q-input label="url" v-model="editItem.Url"/>
          <q-input label="текст" v-model="editItem.Text"/>
          <q-input label="иконка" v-model="editItem.Icon"/>
          <q-checkbox label="вложенное меню" v-model="editItem.IsFolder"/>
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="OK" color="primary" @click="editOK"/>
      </q-card-actions>
    </q-card>
  </q-dialog>

</template>

<script>
  import {ref, toRefs, computed, watch, reactive} from 'vue'
  import _ from 'lodash'
  import draggable from 'vuedraggable'
  export default {
    name: "nested-draggable",
    components: {draggable},
    props: ['project', 'menu'],
    emits: ['update'],
    setup(props, {emit}) {
      const localMenu = ref(props.menu)

      const isShowAddDialog = ref(false)
      const isShowEditDialog = ref(false)
      const newItem = reactive({DocName: null, Url: null, Text: null, Icon: null, IsFolder: null})
      const editItem = ref(null)

      // список документов для меню выбора
      const options = props.project.docs.map(v => {
        return {label: v.name_ru, value: v.name}
      })

      watch(localMenu, ()=> {
        emit('update', localMenu)
      })

      const docNameTitle = computed(()=> {
        return (docName)=> {
          const doc = props.project.docs.find(v => v.name === docName)
          return doc ? _.upperFirst(doc.name_ru) : `doc: ${docName} не найден`
        }
      })

      const docNameIcon = computed(()=> {
        return (docName)=> {
          return '/image/technical-support.svg'
          // const doc = props.project.docs.find(v => v.name === docName)
          // return doc ? doc.menu_icon : `doc: ${docName} не найден`
        }
      })

      const remove = (id) => {
        const i = localMenu.value.findIndex(v => v.id === id)
        localMenu.value.splice(i, 1)
      }

      const add = () => {
        let v = _.cloneDeep(newItem)
        if (v.DocName) v.DocName = v.DocName.value
        if (v.IsFolder) v.LinkList = []
        localMenu.value.push(Object.assign({id: _.random(100000)}, v))
        isShowAddDialog.value = false
      }

      const edit = (item) => {
        editItem.value = _.cloneDeep(item)
        const d = props.project.docs.find(v => v.name === item.DocName)
        if (item.DocName) editItem.value.DocName = {label: d.name_ru, value: d.name}
        isShowEditDialog.value = true
      }

      const editOK = () => {
        const item = localMenu.value.find(v => v.id === editItem.value.id)
        if (editItem.value.DocName) editItem.value.DocName = editItem.value.DocName.value
        Object.assign(item, editItem.value)
        isShowEditDialog.value = false
      }

      return {
        localMenu, docNameTitle, docNameIcon,
        isShowAddDialog, options, newItem,
        isShowEditDialog, editItem,
        remove, add, edit, editOK,
      }
    }
  }
</script>
