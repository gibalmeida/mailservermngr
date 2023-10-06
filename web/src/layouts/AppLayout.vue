<script setup lang="ts">
import TheMenu from '@/components/TheMenu.vue'
</script>

<template>
  <header class="bg-white dark:bg-black dark:text-white">
    <TheMenu />
  </header>
  <main>
    <div class="max-w-6xl mx-auto px-4 py-2 sm:p-4 bg-gray-400 dark:bg-black dark:text-white flex sm:shadow-lg sm:rounded-lg dark:border-white dark:border-2 sm:mt-4">
      <slot></slot>
    </div>
  </main>
</template>
