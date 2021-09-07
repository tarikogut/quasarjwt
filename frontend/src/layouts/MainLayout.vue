<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          Soft STP
        </q-toolbar-title>

        <div>v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      class="bg-grey-1"
    >
      <div class="q-pa-md q-gutter-sm">
        <q-select filled bottom-slots v-model="model" :options="options" label="Select Server" counter maxlength="3"
                  :dense="dense" :options-dense="denseOpts">


          <template v-slot:hint>
            ESTP Servers
          </template>

          <template v-slot:prepend>
            <q-btn round dense flat icon="add" @click="showaddestpDialog"/>
          </template>
        </q-select>
        <q-list dense bordered padding class="rounded-borders">
          <q-item clickable v-ripple @click="alert('ddddddd')">
            <q-item-section>
              Statistics
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              MTP3 Instance
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              SCCP Instance
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              Connections
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              Routing Tables
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              Partner Management
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              User Management
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              General Settings
            </q-item-section>
          </q-item>
          <q-item clickable v-ripple>
            <q-item-section>
              List Definitions
            </q-item-section>
          </q-item>

        </q-list>
      </div>
    </q-drawer>
    <q-dialog ref="dialogRef" @hide="onDialogHide">
      <q-card class="q-dialog-plugin">

        <q-card-actions align="right">
          <q-btn color="primary" label="OK" @click="onOKClick"/>
          <q-btn color="primary" label="Cancel" @click="onCancelClick"/>
        </q-card-actions>
      </q-card>
    </q-dialog>
    <q-page-container>
      <router-view/>
    </q-page-container>
  </q-layout>
</template>

<script>


import {defineComponent, ref} from 'vue'
import {useDialogPluginComponent} from "quasar";

export default defineComponent({
  name: 'MainLayout',


  setup() {
    const leftDrawerOpen = ref(false)
    const {dialogRef, onDialogHide, onDialogOK, onDialogCancel} = useDialogPluginComponent()
    return {
      alert(data) {
        alert(data)
      },
      options: [
        'estp1', 'estp2', 'estp3'
      ],
      dense: ref(false),
      denseOpts: ref(false),
      model: ref(null),
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value
      },
      showaddestpDialog() {
        onDialogHide.value = false
      },
      dialogRef,
      onDialogHide,
      onOKClick() {
        // on OK, it is REQUIRED to
        // call onDialogOK (with optional payload)
        onDialogOK()
        // or with payload: onDialogOK({ ... })
        // ...and it will also hide the dialog automatically
      },
    }
  }
})
</script>
