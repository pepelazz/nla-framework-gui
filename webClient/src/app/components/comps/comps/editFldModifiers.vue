<template>
    <div class="q-gutter-sm">
      <q-checkbox label="поиск" v-model="isSearch"><q-tooltip>поле участвует в поиске</q-tooltip></q-checkbox>
      <q-checkbox label="обязательно к заполнению" v-model="isRequired"><q-tooltip>поле не должно быть пустым</q-tooltip></q-checkbox>
      <q-checkbox label="уникальное" v-model="isUniq"><q-tooltip>ограничение, что больше записи с таким значением быть не может</q-tooltip></q-checkbox>
    </div>
    <div class="row q-gutter-md">
      <div class="col-4">
        <q-input label="v-if" outlined dense v-model="vIf"><q-tooltip>vue условие для показа/скрытия поля</q-tooltip></q-input>
      </div>
      <div class="col-4">
        <q-input label="readonly" outlined dense v-model="readonly"><q-tooltip>поле только на чтение</q-tooltip></q-input>
      </div>

    </div>
</template>

<script>
    import {toRefs, ref, watch} from "vue"
    import _ from 'lodash'

    export default {
        props: ['fld'],
        setup(props) {
          const localFld = toRefs(props).fld
          if (!localFld.value.modifier_list) localFld.value.modifier_list = []
          const isSearch = ref(false)
          const isRequired = ref(false)
          const isUniq = ref(false)

          const vIf = ref(null)
          const readonly = ref(null)

          const setModifier = (v, name) => {
            if (localFld.value.modifier_list.findIndex(v => v.Name===name) > -1) v.value = true
            watch(v, (n)=> n ? localFld.value.modifier_list.push({Name: name}) : _.remove(localFld.value.modifier_list, (f) => f.Name===name))
          }

          // модификатор с аргументом
          const setModifierWithArg = (v, name) => {
            const i = localFld.value.modifier_list.findIndex(v => v.Name===name)
            if ( i > -1) v.value = localFld.value.modifier_list[i].Args[0]
            watch(v, (newV, oldV)=> {
              // вариант добавления
              if (!oldV && newV) {
                localFld.value.modifier_list.push({Name: name, Args: [newV]})
              }
              // удаление модификатора
              if (!newV) _.remove(localFld.value.modifier_list, (f) => f.Name===name)
              // редактирование аргумента
              if (oldV && newV) {
                const i = localFld.value.modifier_list.findIndex(v => v.Name===name)
                if ( i > -1) localFld.value.modifier_list[i].Args[0] = newV
              }
            })
          }

          setModifier(isSearch, 'SetIsSearch')
          setModifier(isRequired, 'SetIsRequired')
          setModifier(isUniq, 'SetIsUniq')
          setModifierWithArg(vIf, 'SetVif')
          setModifierWithArg(readonly, 'SetReadonly')

          return {
            localFld,
            isSearch, isRequired, isUniq, vIf, readonly
          }
        },

    }
</script>

