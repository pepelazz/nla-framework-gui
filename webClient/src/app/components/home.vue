<template>
  <q-page padding>
    <div class="row q-col-gutter-md" v-if="project">
      <div class="col-3">
        <!-- кнопки       -->
        <comp-add-doc-btn @update="getProject"/>

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
      <q-btn fab icon="save" color="accent" @click="saveProject"/>
    </q-page-sticky>
  </q-page>
</template>

<script>
    import {Notify} from "quasar";
    import docFlds from 'src/app/components/docFlds'
    import compAddDocBtn from 'src/app/components/comps/addDocBtn'
    import hookProjectBeforeSave from 'src/app/components/hookProjectBeforeSave'
    import hookProjectBeforeLoad from 'src/app/components/hookProjectBeforeLoad'


    export default {
        components: {docFlds, compAddDocBtn},
        data() {
            return {
              project: null,
              selectedDoc: null,
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
              this.project = hookProjectBeforeLoad(res.result)
            }
          })
        },
        showDoc(doc) {
          this.selectedDoc = doc
        },
        saveProject() {
          this.$utils.postApiRequest({url: "/saveProject", params: hookProjectBeforeSave(this.project)}).subscribe(res => {
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
