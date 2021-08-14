<template>
    <div class="q-gutter-sm">
      <q-checkbox label="ссылка" v-model="localFld.params" val="isShowLink"><q-tooltip>показывать ссылку</q-tooltip></q-checkbox>
      <q-checkbox label="добаление" v-model="localFld.params" val="isAddNew"><q-tooltip>возможность добавлять новую запись</q-tooltip></q-checkbox>
      <q-checkbox label="удаление" v-model="localFld.params" val="isClearable"><q-tooltip>иконка с возможностью удалить ссылку</q-tooltip></q-checkbox>
    </div>
</template>

<script>
    import {toRefs} from "vue";

    export default {
        props: ['fld'],
        setup(props) {
          const localFld = toRefs(props).fld
          if (!localFld.value.params) localFld.value.params = []

          return {
            localFld,
          }
        },

    }
</script>

