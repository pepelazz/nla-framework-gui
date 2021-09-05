<template>
  <q-page padding>
    <div class="row q-col-gutter-md" v-if="project">
      <div class="col-3">
        <!-- кнопки       -->
        <comp-add-doc-btn @update="getProject"/>
        <comp-menu-dialog :project="project" @update="v => project.menu = v"/>

        <!-- список таблиц       -->
        <q-list bordered class="q-mt-sm">
          <q-item v-for="doc in project.docs" :key="doc.id">
            <q-item-section avatar>
              <q-avatar rounded @click="showDoc(doc)" class="cursor-pointer">
                <q-img src="/image/database.svg"/>
              </q-avatar>
            </q-item-section>
            <q-item-section>
              <q-item-label>{{doc.name}}</q-item-label>
              <q-item-label caption>{{doc.name_ru}}</q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
      </div>
      <div class="col-9" v-if="selectedDoc">
        <doc-flds :project="project" :selected-doc="selectedDoc" @update="v => selectedDoc = v"/>
      </div>
    </div>
    <q-page-sticky position="bottom-right" :offset="[18, 18]">
      <q-btn fab icon="save" color="accent" @click="saveProject" :loading="isSaving"/>
    </q-page-sticky>
  </q-page>
</template>

<script>
    import {Notify} from "quasar";
    import docFlds from 'src/app/components/docFlds'
    import compAddDocBtn from 'src/app/components/comps/addDocBtn'
    import compMenuDialog from 'src/app/components/comps/menuDialog/index'
    import hookProjectBeforeSave from 'src/app/components/hookProjectBeforeSave'
    import hookProjectBeforeLoad from 'src/app/components/hookProjectBeforeLoad'
    import _ from 'lodash'

    export default {
        components: {docFlds, compAddDocBtn, compMenuDialog},
        data() {
            return {
              project: null,
              selectedDoc: null,
              isSaving: false,
            }
        },
      methods: {
        getProject() {
          this.$utils.postApiRequest({url: "/getProject", params: {}}).subscribe(res => {
            if (res.ok) {
              // проставляем уникальные id
              res.result.docs = res.result.docs.map((v, i) => {
                v.id = i
                v.flds = v.flds.map((fld, j) => {
                  fld.id = j
                  return fld
                })
                return v
              })
              res.result.menu = res.result.menu.map(v => {
                v.id = _.random(100000)
                if (v.IsFolder && !v.LinkList) v.LinkList = []
                if (v.LinkList) {
                  v.LinkList = v.LinkList.map(v1 => {
                    v1.id = _.random(100000)
                    return v1
                  })
                }
                return v
              })
              this.project = hookProjectBeforeLoad(res.result)
            }
          })
        },
        showDoc(doc) {
          this.selectedDoc = doc
        },
        saveProject() {
          this.isSaving = true
          this.$utils.postApiRequest({url: "/saveProject", params: hookProjectBeforeSave(this.project)}).subscribe(res => {
            this.isSaving = false
            if (res.ok) {
              Notify.create({
                color: 'positive',
                position: 'bottom-right',
                message: "изменения сохранены"
              })
            }
          })
        },
      },
      mounted() {
          this.getProject()
      }
    }
</script>
