<script setup lang="ts">
import AppLayout from '@/layouts/AppLayout.vue';
import DefaultTable from '@/components/DefaultTable.vue';
import { getAccountsByDomain, getAllAccounts } from '@/repositories/accounts'
import { ref } from 'vue'

const props = defineProps(['domain'])

const accounts = ref()
if (props.domain) {
  getAccountsByDomain(props.domain).then((result) => (accounts.value = result.data))
} else {
  getAllAccounts().then((result) => (accounts.value = result.data))
}
</script>
<template>
  <AppLayout>
    <DefaultTable title="E-mails Accounts" :rows="accounts" :col_names="['name','domain']" />
  </AppLayout>
</template>
