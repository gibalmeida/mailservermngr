<script setup lang="ts">
import { ref } from 'vue'
import { Authenticator } from '@/repositories/authenticator'
import router from '@/router'
import { useToast } from 'vue-toastification'
import LoginLayout from '@/layouts/LoginLayout.vue'

const username = ref('')
const password = ref('')
const showLoader = ref(false)
const props = defineProps(['redirect'])
const toast = useToast()

function submit() {
  showLoader.value = true // TODO: The spinner is not working yet
  Authenticator.login(username.value, password.value)
    .then(() => {
       if (props.redirect) {
        router.push(props.redirect)
      } else {
        router.push({ name: 'home' })
      }
    })
    .catch((error) => {
      console.log(error)
      toast.error(error, {timeout:2000})
    })
}
</script>

<template>
  <LoginLayout>
    <div class="flex mx-auto flex-col sm:max-w-lg ring-1 ring-gray-900/5 rounded-md shadow-xl bg-white p-8">
      <div class="text-3xl font-extrabold text-center">Welcome to the Mail Server Manager</div>
      <div class="mx-auto flex mt-6">
        <form @submit.prevent="submit">
          <h1 class="text-2xl">Please Sign In</h1>

          <div class="mt-4">
            <label for="inputUsername" class="font-semibold text-sm block">Username</label>
            <input
              type="text"
              id="inputUsername"
              v-model="username"
              placeholder="Username"
              required
              autofocus
              autocomplete="username"
              class="p-2 ring-1 rounded-sm"
            />
          </div>

          <div class="mt-4">
            <label for="inputPassword" class="font-semibold text-sm block">Password</label>
            <input
              type="password"
              id="inputPassword"
              v-model="password"
              placeholder="Password"
              required
              autocomplete="current-password"
              class="p-2 ring-1 rounded-sm"
            />
          </div>
          <div class="mt-6 flex justify-center">
            <button type="submit" class="px-10 py-2 text-white font-semibold bg-black rounded-3xl hover:bg-gray-800">
              <fa-icon icon="fa-solid fa-right-to-bracket"></fa-icon>
              &nbsp;Sign In
              <span
                v-if="showLoader"
                class="spinner-grow spinner-grow-sm"
                style="width: 1.3rem; height: 1.3rem"
                role="status"
                aria-hidden="true"
              ></span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </LoginLayout>
</template>
