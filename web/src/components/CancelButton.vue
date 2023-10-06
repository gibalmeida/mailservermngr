<script setup lang="ts">
const props =defineProps(['click'])
</script>
<template>
    <button class="bg-gray-500 text-white px-5 py-2 rounded-full hover:bg-gray-700" @click="props.click">
        <slot></slot>
    </button>
</template>