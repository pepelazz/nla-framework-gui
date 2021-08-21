<template>
    <div class="q-gutter-sm">
      <q-checkbox label="поиск" v-model="isSearch"><q-tooltip>поле участвует в поиске</q-tooltip></q-checkbox>
      <q-checkbox label="обязательно к заполнению" v-model="isRequired"><q-tooltip>поле не должно быть пустым</q-tooltip></q-checkbox>
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

          const setModifier = (v, name) => {
            if (localFld.value.modifier_list.findIndex(v => v.Name===name) > -1) v.value = true
            watch(v, (n)=> n ? localFld.value.modifier_list.push({Name: name}) : _.remove(localFld.value.modifier_list, (f) => f.Name===name))
          }

          setModifier(isSearch, 'SetIsSearch')
          setModifier(isRequired, 'SetIsRequired')

          return {
            localFld,
            isSearch, isRequired,
          }
        },

    }
</script>

