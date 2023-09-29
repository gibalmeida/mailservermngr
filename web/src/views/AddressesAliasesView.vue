<script setup lang="ts">
import AppLayout from '@/layouts/AppLayout.vue';
import DefaultTable from '@/components/DefaultTable.vue';
import { ref } from 'vue'
import { getAllAddressesAliases } from '@/repositories/addresses_aliases';

const addressesAliases = ref()
getAllAddressesAliases().then((result) => (addressesAliases.value = result.data))
</script>
<template>
  <AppLayout>
    <DefaultTable title="Domains Aliases" :rows="addressesAliases" :col_names="['alias','addresses']" />
  </AppLayout>
</template>
