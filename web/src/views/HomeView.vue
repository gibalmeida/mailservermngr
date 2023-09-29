<script setup lang="ts">
import AppLayout from '@/layouts/AppLayout.vue'
import { getAccountsByDomain } from '@/repositories/accounts'
import { getAddressesAliasesByDomain } from '@/repositories/addresses_aliases'
import {
  type DomainAlias,
  type Account,
  type Domain,
  type AddressAlias
} from '@/repositories/clients/mailservermngr/api'
import { getAllDomains } from '@/repositories/domains'
import { getDomainsAliasesByDomain } from '@/repositories/domains_aliases'
import { ref } from 'vue'

const domains = ref<Domain[]>()
const accounts = ref<Account[]>()
const domainsAliases = ref<DomainAlias[]>()
const addressesAliases = ref<AddressAlias[]>()
const selectedDomain = ref<string>()
const selectedAccount = ref<string>()
const defaultTab = 'accounts'
const tabActive = ref<string>(defaultTab)

getAllDomains().then((result) => (domains.value = result.data))

function loadDomainObjects(domain: string) {
  selectedDomain.value = domain

  getAccountsByDomain(domain).then((result) => (accounts.value = result.data))
  getDomainsAliasesByDomain(domain).then((result) => (domainsAliases.value = result.data))
  getAddressesAliasesByDomain(domain).then((result) => (addressesAliases.value = result.data))
}

function editAccount(account: Account) {
  // implement me
}

function cancelDomainSelected() {
  selectedDomain.value = ''
  tabActive.value = defaultTab
  selectedAccount.value = ''
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col">
      <div class="flex flex-row">
        <div class="flex flex-col" v-if="selectedDomain">
          <div class="flex flex-row">
            <div class="text-2xl font-semibold pr-2">{{ selectedDomain }}</div>
            <div class="">
              <button
                class="bg-gray-500 px-4 rounded-lg font-extralight"
                @click="cancelDomainSelected"
              >
                <fa-icon icon="fa-solid fa-close"></fa-icon>
              </button>
              <button class="bg-gray-500 px-4 rounded-lg font-extralight">
                <fa-icon icon="fa-solid fa-trash"></fa-icon>
              </button>
            </div>
          </div>
        </div>
        <div class="flex flex-col px-4" v-else>
          <div>
            <h1 class="text-2xl font-semibold">Local Domains</h1>
            <p class="font-light text-xs">Please select one domain</p>
          </div>
          <div class="mt-4">
            <div
              v-for="domain in domains"
              :key="domain.domain"
              class="border-b transition duration-300 ease-in-out hover:bg-neutral-100 dark:border-neutral-500 dark:hover:bg-neutral-600"
              @click="loadDomainObjects(domain.domain)"
            >
              {{ domain.domain }}
            </div>
          </div>
        </div>
      </div>

      <div class="mb-4 border-b border-gray-200 dark:border-gray-700" v-if="selectedDomain">
        <ul class="flex flex-wrap -mb-px text-sm font-medium text-center">
          <li class="mr-2">
            <button
              class="inline-block p-4 border-b-2 rounded-t-lg"
              type="button"
              @click="tabActive = 'accounts'"
            >
              Accounts
            </button>
          </li>
          <li class="mr-2">
            <button
              class="inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
              type="button"
              @click="tabActive = 'domains-aliases'"
            >
              Domains Aliases
            </button>
          </li>
          <li class="mr-2">
            <button
              class="inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
              type="button"
              @click="tabActive = 'addresses-aliases'"
            >
              Addreses Aliases
            </button>
          </li>
        </ul>
      </div>
      <div v-if="selectedDomain">
        <div class="p-4 rounded-lg bg-gray-50 dark:bg-gray-800" v-if="tabActive == 'accounts'">
          <div class="flex flex-col ring-1 px-4" >
            <div>
              <h1 class="text-2xl font-semibold">Domain Accounts</h1>
            </div>
            <div class="mt-4">
              <div
                v-for="account in accounts"
                :key="account.name + '@' + account.domain"
                class="border-b transition duration-300 ease-in-out hover:bg-neutral-100 dark:border-neutral-500 dark:hover:bg-neutral-600"
                @click="editAccount(account)"
              >
                {{ account.name + '@' + account.domain }}
              </div>
            </div>
          </div>
        </div>
        <div class=" p-4 rounded-lg bg-gray-50 dark:bg-gray-800" v-if="tabActive == 'domains-aliases'">
          <div class="flex flex-col ring-1 px-4" >
            <div>
              <h1 class="text-2xl font-semibold">Domain Aliases</h1>
            </div>
            <div class="mt-4">
              <div v-for="domainAlias in domainsAliases" :key="domainAlias.alias">
                {{ domainAlias.alias }}
              </div>
            </div>
          </div>
        </div>
        <div class="p-4 rounded-lg bg-gray-50 dark:bg-gray-800" v-if="tabActive == 'addresses-aliases'">
          <div class="flex flex-col ring-1 px-4" >
            <div>
              <h1 class="text-2xl font-semibold">Addresses Aliases</h1>
            </div>
            <div class="mt-4 flex flex-col">
              <div
                v-for="addressAlias in addressesAliases"
                :key="addressAlias.alias"
                class="flex flex-row justify-between"
              >
                <div>{{ addressAlias.alias }}</div>
                <div>{{ addressAlias.addresses }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </AppLayout>
</template>
