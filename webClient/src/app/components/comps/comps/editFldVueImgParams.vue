<template>
    <div>
      <q-input label="ограничение расширений файлов" v-model="localParams.Accept" outlined dense class="q-mt-sm">
        <template v-slot:append>
          <q-icon name="help" color="grey">
            <q-tooltip>перечень через запятую. Например "png, jpg, gif" </q-tooltip>
          </q-icon>
        </template>
      </q-input>
      <q-input label="максимальный размер файла (байтах)" v-model="localParams.MaxFileSize" type="number" outlined dense  class="q-mt-sm">
        <template v-slot:append>
          <q-icon name="help" color="grey">
            <q-tooltip>если 0, то ограничений нет</q-tooltip>
          </q-icon>
        </template>
      </q-input>
      <q-input label="максмиальная ширина (px)" v-model="localParams.Width" type="number" outlined dense  class="q-mt-sm">
        <template v-slot:append>
          <q-icon name="help" color="grey">
            <q-tooltip>если 0, то ограничений нет</q-tooltip>
          </q-icon>
        </template>
      </q-input>
      <q-input label="кадрирование" v-model="localParams.Crop" outlined dense class="q-mt-sm">
        <template v-slot:append>
          <q-icon name="help" color="grey">
            <q-tooltip>[ширины]х[высота] - например, 600x400</q-tooltip>
          </q-icon>
        </template>
      </q-input>
      <q-checkbox label="можно ли добавлять внешние ссылки" v-model="localParams.CanAddUrls" outlined dense class="q-mt-sm"/>
    </div>
</template>

<script>
    import {toRefs} from "vue";

    export default {
        props: ['params'],
        setup(props) {
          const localParams = toRefs(props).params

          return {
            localParams,
          }
        },

    }
</script>

