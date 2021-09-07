<template>
  <div class="fixed fixed-center bg-grey-1">
    <form v-on:submit.prevent="onSubmit()">
      <q-card square class="flex-center text-center" style="width: 400px; padding:50px">
        <h6>Fink Admin</h6>
        <q-card-separator/>
        <div class="text-left">
          <q-input id="username" type="text" v-model.trim="username"

                   required autofocus/>
          <q-input id="password" type="password" v-model="password" required/>
          <br>
        </div>
        <q-card-actions>
          <q-btn type="submit" class="fit" color="primary">Login</q-btn>
        </q-card-actions>
      </q-card>
    </form>
  </div>
</template>

<style>
</style>

<script>
import {ref} from "vue";

export default {
  name: 'Login',
  setup() {

    const username = ref("")
    const password = ref("")


    function onSubmit() {
api.post("/api")

    }

    return {
      username, password,
      onSubmit
    }
  },

}
</script>
