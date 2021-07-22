<template>
  <div>

    <q-layout view="hHh Lpr lFf">

      <q-header reveal elevated class="bg-white text-grey-8 q-py-xs">
        <q-toolbar>
          <q-btn dense flat round icon="menu" @click="leftSide = !leftSide"/>

          <!-- лого -->
          <q-btn flat no-caps no-wrap class="q-ml-xs" @click="$router.push('/')">
            <!--            <q-avatar size="26px">-->
            <!--              <img src="https://www.defly.ru/website/defly/template/images/logo.png">-->
            <!--            </q-avatar>-->



            <q-toolbar-title shrink class="text-weight-bold">
              NLA Framework
            </q-toolbar-title>

          </q-btn>

          <q-space/>

          <!-- аватарка и меню пользователя -->

        </q-toolbar>
      </q-header>

      <!-- боковое меню     -->
<!--      <side-menu :leftSide="leftSide"  @hide="leftSide=false"/>-->

      <q-page-container>
        <router-view :key="$route.fullPath"/>
      </q-page-container>

    </q-layout>
  </div>
</template>

<script>
  // import sideMenu from './app/components/sidemenu/index.vue'

  export default {
    // components: {sideMenu},
    data() {
      return {
        leftSide: true,
      }
    },
  }
</script>
