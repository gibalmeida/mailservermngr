<script setup lang="ts">
import TheMenu from '@/components/TheMenu.vue'
</script>

<template>
  <header class="bg-white dark:bg-black dark:text-white">
    <TheMenu />
  </header>
  <main>
    <div class="max-w-3xl mx-auto px-4 py-4 bg-gray-200 dark:bg-black dark:text-white flex shadow-lg rounded-lg dark:border-white dark:border-2 mt-8">
      <slot></slot>
    </div>
  </main>
</template>
