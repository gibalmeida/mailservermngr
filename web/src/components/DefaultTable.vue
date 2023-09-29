<script setup lang="ts">
const props = defineProps<{
  title: string
  col_names: string[]
  rows: string[]
  url_prefix: string
  url_colname: string
}>()
</script>
<template>
  <h1 class="text-3xl px-4 py-4">{{ props.title }}</h1>
  <table class="min-w-full text-left text-xs md:text-xl font-light">
    <thead class="border-b bg-white font-medium dark:border-neutral-500 dark:bg-neutral-700">
      <th class="px-6 py-4" scope="col">#</th>
      <th class="px-6 py-4 capitalize" scope="col" v-for="col_name in col_names" :key="col_name">
        {{ col_name }}
      </th>
      <th class="px-6 py-4" scope="col">Actions</th>
    </thead>
    <tbody>
      <tr
        v-for="(row, idx) in props.rows"
        :key="idx"
        class="border-b transition duration-300 ease-in-out hover:bg-neutral-100 dark:border-neutral-500 dark:hover:bg-neutral-600"
        @click="$router.push(url_prefix + row[url_colname])"
      >
        <td class="whitespace-nowrap px-6 py-4">{{ idx }}</td>

        <td v-for="(col, idx2) in col_names" :key="idx2" class="px-6 py-4">{{ row[col] }}</td>
        <td class="whitespace-nowrap px-6 py-4"><fa-icon icon="fa-solid fa-trash" /></td>
      </tr>
    </tbody>
  </table>
</template>
