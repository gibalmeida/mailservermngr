<script setup lang="ts">
import TheMenu from '@/components/TheMenu.vue'
</script>

<template>
  <header class="bg-white">
    <TheMenu />
  </header>
  <main>
    <div class="max-w-7xl mx-auto ring-1 px-4 pt-2 min-h-screen bg-gray-200">
      <slot></slot>
    </div>
  </main>
</template>
