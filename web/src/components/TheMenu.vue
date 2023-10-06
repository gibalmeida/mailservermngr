<script setup lang="ts">
import { Authenticator } from '@/repositories/authenticator'
import router from '@/router'
import { ref } from 'vue'

const mobileMenuOpened = ref(false)

function logout() {
  Authenticator.logout()
  router.push({ name: 'login' })
}
</script>

<template>
  <nav class="flex justify-between items-center w-[92%] max-w-6xl mx-auto">
    <div class="h-14 flex items-center gap-6 xl:text-xl">
      <!-- <fa-icon
        v-if="!mobileMenuOpened"
        icon="fa-solid fa-bars"
        class="text-3xl cursor-pointer md:hidden"
        @click="mobileMenuOpened = !mobileMenuOpened"
      />
      <fa-icon
        v-if="mobileMenuOpened"
        icon="fa-solid fa-close"
        class="text-3xl cursor-pointer md:hidden"
        @click="mobileMenuOpened = !mobileMenuOpened"
      /> -->
      <span class="font-bold">Mail Server Manager</span>
    </div>
    <!-- <div
      class="nav-links duration-500 md:static absolute bg-white dark:bg-black md:min-h-fit min-h-[40vh] left-0 md:w-auto w-full flex items-center px-5"
      :class="{ 'top-[10%] shadow-lg': mobileMenuOpened, 'top-[-100%]': !mobileMenuOpened }"
    >
      <ul class="flex md:flex-row flex-col md:items-center md:gap-[4vw] gap-4">
        <li>
          <RouterLink class="hover:border-b-4 hover:text-gray-600" to="/">Home</RouterLink>
        </li>
        <li>
          <RouterLink class="hover:border-b-4 hover:text-gray-600" to="/accounts">Accounts</RouterLink>
        </li>
        <li>
          <RouterLink class="hover:border-b-4 hover:text-gray-600" to="/domains">Domains</RouterLink>
        </li>
        <li>
          <RouterLink class="hover:border-b-4 hover:text-gray-600" to="/domains-aliases">Domains Aliases</RouterLink>
        </li>
        <li>
          <RouterLink class="hover:border-b-4 hover:text-gray-600" to="/addresses-aliases"
            >Addresses Aliases</RouterLink
          >
        </li>
      </ul>
    </div> -->
    <div class="flex items-center gap-6 text-xs xl:text-sm">
      <button class="bg-[#2f79f8] text-white px-5 py-2 rounded-full hover:bg-[#6885b8] dark:bg-slate-600" @click="logout">
        <fa-icon icon="fa-icon-solid fa-right-from-bracket" /> Sign Out
      </button>
    </div>
  </nav>
</template>
