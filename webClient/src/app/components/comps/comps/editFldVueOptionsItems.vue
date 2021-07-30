<template>
    <q-list dense>
      <q-item-label header>Список вариантов</q-item-label>
      <q-item v-for="(item, index) in localOptions" :key="item.id" style="padding-left: 0; padding-right: 0">
        <q-item-section >
          <q-input label="название" v-model="item.label" outlined dense type="text"/>
        </q-item-section>
        <q-item-section>
          <q-input label="значение" v-model="item.value" outlined dense type="text"/>
        </q-item-section>
        <q-item-section>
          <q-input
            filled
            dense
            v-model="item.color"
            style="max-width: 250px"
          >
            <template v-slot:append>
              <q-icon name="colorize" class="cursor-pointer">
                <q-popup-proxy transition-show="scale" transition-hide="scale">
                  <q-color v-model="item.color" />
                </q-popup-proxy>
              </q-icon>
            </template>
          </q-input>
        </q-item-section>
        <q-item-section side>
          <q-btn icon="clear" size="sm" flat round @click="remove(index)"/>
        </q-item-section>
      </q-item>
      <q-item style="padding-left: 0; padding-right: 0">
        <q-item-section side>
          <q-btn icon="add" flat size="sm" @click="add" round/>
        </q-item-section>
      </q-item>
    </q-list>
</template>

<script>
    import {toRefs} from "vue";

    export default {
        props: ['options'],
        setup(props) {
          const localOptions = toRefs(props).options
          const add = () => {
            localOptions.value.push({id: `${_.random(1,1000)}`, label: null, value: null, color: '#027be3'})
          }
          const remove = (i) => {
            localOptions.value.splice(i, 1)
          }
          return {
            localOptions,
            add,
            remove,
          }
        },

    }
</script>

