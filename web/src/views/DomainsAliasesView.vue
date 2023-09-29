<script setup lang="ts">
import AppLayout from '@/layouts/AppLayout.vue';
import DefaultTable from '@/components/DefaultTable.vue';
import { ref } from 'vue'
import { getAllDomainsAliases } from '@/repositories/domains_aliases';

const domainsAliases = ref()
getAllDomainsAliases().then((result) => (domainsAliases.value = result.data))
</script>
<template>
  <AppLayout>
    <DefaultTable title="Domains Aliases" :rows="domainsAliases" :col_names="['alias','domain']" />
  </AppLayout>
</template>
