<script setup lang="ts">
import AppLayout from '@/layouts/AppLayout.vue'
import { getAccountsByDomain } from '@/repositories/accounts';
import { getAddressesAliasesByDomain } from '@/repositories/addresses_aliases';
import { type DomainAlias, type Account, type Domain, type AddressAlias } from '@/repositories/clients/mailservermngr/api'
import { getAllDomains } from '@/repositories/domains'
import { getDomainsAliasesByDomain } from '@/repositories/domains_aliases';
import { ref } from 'vue'

const domains = ref<Domain[]>()
const accounts = ref<Account[]>()
const domainsAliases = ref<DomainAlias[]>()
const addressesAliases = ref<AddressAlias[]>()
const selectedDomain = ref<string>()

getAllDomains().then((result) => (domains.value = result.data))

function loadDomainObjects(domain: string) {
  selectedDomain.value = domain
  
  getAccountsByDomain(domain).then((result) => accounts.value=result.data)
  getDomainsAliasesByDomain(domain).then((result) => domainsAliases.value=result.data)
  getAddressesAliasesByDomain(domain).then((result) => addressesAliases.value=result.data)

}
</script>

<template>
  <AppLayout>
    <div class="flex flex-row">
      <div class="flex flex-col px-4">
        <div>
          <h1 class="text-2xl font-semibold">Server Local Domains</h1>
        </div>
        <div class="mt-4">
          <div v-for="domain in domains" :key="domain.domain" class="border-b transition duration-300 ease-in-out hover:bg-neutral-100 dark:border-neutral-500 dark:hover:bg-neutral-600" @click="loadDomainObjects(domain.domain)">
            {{ domain.domain }}
          </div>
        </div>
      </div>
      <div class="flex flex-col ring-1 px-4">
        <div>
          <h1 class="text-2xl font-semibold">Domain Aliases</h1>
        </div>
        <div class="mt-4">
          <div v-for="domainAlias in domainsAliases" :key="domainAlias.alias">
            {{ domainAlias.alias }}
          </div>
        </div>
      </div> 
      <div class="flex flex-col ring-1 px-4">
        <div>
          <h1 class="text-2xl font-semibold">Domain Accounts</h1>
        </div>
        <div class="mt-4">
          <div v-for="account in accounts" :key="account.name+'@'+account.domain">
            {{ account.name+'@'+account.domain }}
          </div>
        </div>
      </div>
      <div class="flex flex-col ring-1 px-4">
        <div>
          <h1 class="text-2xl font-semibold">Addresses Aliases</h1>
        </div>
        <div class="mt-4 flex flex-row">
          <div v-for="addressAlias in addressesAliases" :key="addressAlias.alias">
            <div>{{ addressAlias.alias }}</div>
            <div>{{ addressAlias.addresses }}</div>
          </div>
        </div>
      </div>      
     
    </div>
  </AppLayout>
</template>
