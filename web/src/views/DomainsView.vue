<script setup lang="ts">
import DefaultTable from '@/components/DefaultTable.vue';
import AppLayout from '@/layouts/AppLayout.vue'
import { getAllDomains } from '@/repositories/domains'
import { ref } from 'vue'

const domains = ref()

getAllDomains().then((result) => (domains.value = result.data))
</script>

<template>
  <AppLayout>
    <button class="px-4 py-2 rounded-full bg-black text-white">New Domain</button>
    <DefaultTable title="Domains" :rows="domains" :col_names="['domain']" url_prefix="/accounts?domain=" url_colname="domain" />
  </AppLayout>
</template>
