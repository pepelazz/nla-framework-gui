<template>
  <q-btn flat round icon="add" color="secondary" size="sm" @click="isShowDialogAddFld = true">
    <q-tooltip>добавить поле</q-tooltip>
  </q-btn>
  <q-dialog v-model="isShowDialogAddFld" position="top">
    <q-stepper
      v-model="step"
      vertical
      color="primary"
      animated
      style="width: 650px"
    >
      <q-step
        :name="1"
        title="Выберите тип контрола"
        icon="settings"
        :done="step > 1"
      >
        <q-select v-model="newFld" :options="optionsFldType" label="тип поля" />

        <q-select v-if="newFld.value === 'GetFldRef'" v-model="refTableName" :options="docListOptions" label="ссылка на таблицу" />

        <q-stepper-navigation>
          <q-btn @click="toStep2" color="primary" label="Далее" />
        </q-stepper-navigation>
      </q-step>

      <q-step
        :name="2"
        title="Название и параметры контрола"
        icon="create_new_folder"
        :done="step > 2"
      >
<!--        Название переменной на английском (без заглавных букв и пробелов snake_case) и название на русском, которое будет отображаться на интерфейсе-->
        <q-input label="название для базы данных" v-model="name" dense outlined class="q-mb-sm">
          <template v-slot:append>
            <q-icon name="help" color="grey">
              <q-tooltip>Название переменной на английском (без заглавных букв и пробелов snake_case)</q-tooltip>
            </q-icon>
          </template>
        </q-input>

        <q-input label="название на интерфейсе" v-model="name_ru" dense outlined class="q-mb-sm">
          <template v-slot:append>
            <q-icon name="help" color="grey">
              <q-tooltip>название на русском, которое будет отображаться на интерфейсе</q-tooltip>
            </q-icon>
          </template>
        </q-input>

        <q-input v-if="newFld.value === 'GetFldString'" label="максимальное количество знаков" v-model="size" type="number" dense outlined class="q-mb-sm">
          <template v-slot:append>
            <q-icon name="help" color="grey">
              <q-tooltip>если 0 то нет ограничений по длине</q-tooltip>
            </q-icon>
          </template>
        </q-input>

        <edit-fld-vue-options-items v-if="['GetFldSelectString', 'GetFldSelectMultiple', 'GetFldRadioString'].includes(newFld.value)" :options="fldVueOptionsItems"/>
        <edit-fld-vue-files-params v-if="newFld.value === 'GetFldFiles'" :params="fldVueFilesParams"/>

        <q-checkbox label="обязательное к заполнению" v-model="is_required"/>

        <q-stepper-navigation>
          <q-btn @click="addNewFld" color="primary" label="Добавить" />
          <q-btn flat @click="step = 1" color="primary" label="Назад" class="q-ml-sm" />
        </q-stepper-navigation>
      </q-step>

    </q-stepper>

  </q-dialog>
</template>

<script>
    import {ref, computed, reactive} from 'vue'
    import {Notify} from 'quasar'
    import editFldVueOptionsItems from 'src/app/components/comps/comps/editFldVueOptionsItems'
    import editFldVueFilesParams from 'src/app/components/comps/comps/editFldVueFilesParams'
    export default {
      components: {editFldVueOptionsItems, editFldVueFilesParams},
      props: ['selectedDoc', 'project'],
        emits: ['update'],
        setup(props, {emit}) {
          const isShowDialogAddFld = ref(false)
          const newFld = ref({label: 'строка', value: 'GetFldString'})
          const step = ref(1)
          const name = ref(null)
          const name_ru = ref(null)
          const size = ref(0)
          const is_required = ref(false)
          const refTableName = ref(null)
          const fldVueOptionsItems = ref([])
          const fldVueFilesParams = reactive({Accept: null, MaxFileSize: null})

          let docListOptions = computed(() => {
            return props.project.docs.filter(d => d.name !== props.selectedDoc.name)
            .map(d => {
              return {label: d.name_ru, value: d.name}
            })
          })

          const optionsFldType = [
            {label: 'строка', value: 'GetFldString'},
            {label: 'дата', value: 'GetFldDate'},
            {label: 'дата и время', value: 'GetFldDateTime'},
            {label: 'число целое', value: 'GetFldInt64'},
            {label: 'число дробное', value: 'GetFldInt64'},
            {label: 'check box', value: 'GetFldCheckbox'},
            {label: 'телефон', value: 'GetFldPhone'},
            {label: 'email', value: 'GetFldEmail'},
            {label: 'тэги', value: 'GetFldTag'},
            {label: 'ссылка на другую таблицу', value: 'GetFldRef'},
            {label: 'выбор из списка', value: 'GetFldSelectString'},
            {label: 'выбор нескольких вариантов из списка', value: 'GetFldSelectMultiple'},
            {label: 'выбор из списка (radio button)', value: 'GetFldRadioString'},
            {label: 'файлы', value: 'GetFldFiles'},
          ]

          const toStep2 = () => {
            if (newFld.value.value === 'GetFldRef') {
              if (refTableName.value == null) {
                Notify.create({type: 'negative', message: `надо указать название таблицы, на которую будет ссылка`})
                return
              }
              // автоматом проставляем имена в случае GetFldRef
              name.value = `${refTableName.value.value}_id`
              name_ru.value = `${refTableName.value.label}`
            }
            step.value = 2
          }

          const addNewFld = () => {
            // проверка заполнения полей
            if (!name.value) {
              Notify.create({type: 'negative', message: `поле 'name' не заполнено`})
              return
            }
            // убираем все лишиние символы и пробелы. Приводим к маленьким буквам
            name.value = name.value.replace(/[^\w\s]/gi, '').replace(/ /g, '').toLowerCase()
            if (name.value?.length < 4) {
              Notify.create({type: 'negative', message: `поле 'name' должно быть не меньше 4 букв`})
              return
            }
            // проверяем что нет дублирования уже существующих названий полей
            if (name.value === 'title' || props.selectedDoc.flds.findIndex(v => v.name === name.value) > -1) {
              Notify.create({type: 'negative', message: `${name.value} уже есть такое название. Придумайте другое`})
              return
            }
            // в случае GetFldRef проверяем, чтобы название поля заканчивалось на _id
            if (newFld.value.value === 'GetFldRef' && !name.value.endsWith('_id')) {
              Notify.create({type: 'negative', message: `в случае ссылки на другую таблицу поле 'name' должно оканчиваться на _id`})
              return
            }

            if (!name_ru.value || name_ru.value?.length < 4) {
              Notify.create({type: 'negative', message: `поле 'name_ru' должно быть не меньше 4 букв`})
              return;
            }

            // проверка для FldSelectString
            const {fld_vue_options_item, message} = checkGetFldSelectString(newFld, fldVueOptionsItems)
            if (message) {
                  Notify.create({type: 'negative', message})
                  return;
            }

            // определяем row_col
            const lastFld = props.selectedDoc.flds[props.selectedDoc.flds.length-1]
            const rc = lastFld.row_col ? lastFld.row_col[0] : [1, 1]
            let row_col = rc[1] > 1 || !['col-2', 'col-3', 'col-4'].includes(lastFld.col_class)  ? [[rc[0] + 1, 1]] : [[rc[0], 2]]

            emit('update', {
              func_name: newFld.value.value, name: name.value, name_ru: name_ru.value, size: +size.value, row_col, ref_table: refTableName.value?.value,
              fld_vue_options_item,
              fld_vue_files_params: fldVueFilesParams,
              col_class: 'col-4'})
            isShowDialogAddFld.value = false
            name.value = null
            name_ru.value = null
            size.value = 0
            is_required.value = false
            fldVueOptionsItems.value = []
            // fldVueFilesParams.value.MaxFileSize = null
            // fldVueFilesParams.value.Accent = null
          }

          return {
            isShowDialogAddFld,
            optionsFldType,
            newFld,
            docListOptions,
            addNewFld,
            step,
            name,
            name_ru,
            size,
            is_required,
            refTableName,
            fldVueOptionsItems,
            fldVueFilesParams,
            toStep2,
          }
        }
    }

    const checkGetFldSelectString = (newFld, fldVueOptionsItems) => {
      if (['GetFldSelectString', 'GetFldSelectMultiple', 'GetFldRadioString'].includes(newFld.value.value)) {
        let isError = false
        fldVueOptionsItems.value.map(v => {
          if ((!v.Label || v.Label.length === 0) || (!v.Value || v.Value.length === 0))  isError = true
        })
        if (isError) return {message: `не заполнены значения списка`}

        // проверяем уникальность label и value
        const labelArr = _(fldVueOptionsItems.value.map(v=>v.Label)).groupBy().pickBy(x => x.length > 1).keys().value()
        if (labelArr.length > 0) return {message: `дублирование ${labelArr[0]}`}

        const valueArr = _(fldVueOptionsItems.value.map(v=>v.Value)).groupBy().pickBy(x => x.length > 1).keys().value()
        if (valueArr.length > 0) return {message: `дублирование ${valueArr[0]}`}


        // перед сохранением модифицируем
        return {fld_vue_options_item: fldVueOptionsItems.value.map(v => {
            return {Label: v.Label, Value: v.Value, Color: v.Color}
          })}
      }
      return {fld_vue_options_item: []}
    }

</script>

