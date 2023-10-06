<script setup lang="ts">
import AppLayout from '@/layouts/AppLayout.vue'
import AccountForm from '@/components/AccountForm.vue'
import PrimaryButton from '@/components/PrimaryButton.vue'

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
const selectedAccount = ref<Account>()
const newAccount = ref<Account>()
const showNewAccountForm = ref(false)
const defaultTab = 'accounts'
const tabActive = ref<string>(defaultTab)

getAllDomains().then((result) => (domains.value = result.data))

function loadDomainObjects(domain: string) {
  selectedDomain.value = domain

  getAccountsByDomain(domain).then((result) => (accounts.value = result.data))
  getDomainsAliasesByDomain(domain).then((result) => (domainsAliases.value = result.data))
  getAddressesAliasesByDomain(domain).then((result) => (addressesAliases.value = result.data))
}

function newAccountButtonPressed() {
  showNewAccountForm.value=true
  newAccount.value = {name: '', domain: selectedDomain.value ?? ''}

}
function editAccount(account: Account) {
  selectedAccount.value = account
  // implement me
}

function cancelDomainSelected() {
  selectedDomain.value = ''
  tabActive.value = defaultTab
  selectedAccount.value = undefined
}

function cancelAccountSelected() {
  selectedAccount.value = undefined
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col w-full">
      <div class="flex flex-row">
        <div class="flex flex-col" v-if="selectedDomain">
          <div class="flex flex-row">
            <div class="text-2xl font-semibold pr-2">Domain: {{ selectedDomain }}</div>
            <div class="">
              <button
                class="bg-gray-500 px-4 rounded-lg font-extralight"
                @click="cancelDomainSelected"
                title="Cancel domain selection"
              >
                <fa-icon icon="fa-solid fa-close"></fa-icon>
              </button>
              <button class="bg-gray-500 px-4 rounded-lg font-extralight">
                <fa-icon icon="fa-solid fa-trash"></fa-icon>
              </button>
            </div>
          </div>
        </div>
        <div class="flex flex-col w-full" v-else>
          <div>
            <h1 class="text-2xl font-semibold">Local Domains</h1>
            <p class="font-light text-xs">Please select one domain</p>
          </div>
          <div class="mt-4 cursor-pointer">
            <div
              v-for="domain in domains"
              :key="domain.domain"
              class="order-b transition duration-300 ease-in-out hover:bg-neutral-600 dark:border-neutral-500 dark:hover:bg-neutral-600"
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
              class="inline-block p-2 border-b-2 rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
              :class="{ ' border-black': tabActive == 'accounts' }"
              type="button"
              @click="tabActive = 'accounts'"
            >
              Accounts
            </button>
            <button class="bg-gray-500 px-2 rounded-lg font-extralight" @click="newAccountButtonPressed" v-if="!showNewAccountForm"><fa-icon icon="fa-solid fa-plus"></fa-icon></button>
          </li>
          <li class="mr-2">
            <button
              class="inline-block p-2 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
              :class="{ 'border-black': tabActive == 'domains-aliases' }"
              type="button"
              @click="tabActive = 'domains-aliases'"
            >
              Domains Aliases
            </button>
            <button class="bg-gray-500 px-2 rounded-lg font-extralight"><fa-icon icon="fa-solid fa-plus"></fa-icon></button>
          </li>
          <li class="mr-2">
            <button
              class="inline-block p-2 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
              :class="{ 'border-black': tabActive == 'addresses-aliases' }"
              type="button"
              @click="tabActive = 'addresses-aliases'"
            >
              Addreses Aliases
            </button>
            <button class="bg-gray-500 px-2 rounded-lg font-extralight"><fa-icon icon="fa-solid fa-plus"></fa-icon></button>
          </li>
        </ul>
      </div>
      <div v-if="selectedDomain" class="cursor-pointer">
        <div class="py-2" v-if="tabActive == 'accounts'">
          <div class="flex flex-col">
            <div>
              <div v-if="!selectedAccount">
                <div v-if="!showNewAccountForm">
                  <div
                    v-for="account in accounts"
                    :key="account.name + '@' + account.domain"
                    class="border-b transition duration-300 ease-in-out hover:bg-neutral-600 dark:border-neutral-500 dark:hover:bg-neutral-600"
                    @click="editAccount(account)"
                  >
                    {{ account.name + '@' + account.domain }}
                  </div>
                </div>
                <div v-else>
                  <AccountForm v-bind:account="newAccount!" :cancel-button="() => showNewAccountForm=false" :new="true"></AccountForm>
                </div>
              </div>
              <div v-else>
                <AccountForm
                  v-bind:account="selectedAccount"
                  :cancel-button="cancelAccountSelected"
                  :new="false"
                ></AccountForm>
              </div>
            </div>
          </div>
        </div>
        <div class="py-2" v-if="tabActive == 'domains-aliases'">
          <div class="flex flex-col">
            <div>
              <div
                v-for="domainAlias in domainsAliases"
                :key="domainAlias.alias"
                class="border-b transition duration-300 ease-in-out hover:bg-neutral-600 dark:border-neutral-500 dark:hover:bg-neutral-600"
              >
                {{ domainAlias.alias }}
              </div>
            </div>
          </div>
        </div>
        <div class="py-2" v-if="tabActive == 'addresses-aliases'">
          <div class="flex flex-col">
            <div class="flex flex-col">
              <div
                v-for="addressAlias in addressesAliases"
                :key="addressAlias.alias"
                class="flex flex-row border-b transition duration-300 ease-in-out hover:bg-neutral-600 dark:border-neutral-500 dark:hover:bg-neutral-600"
              >
                <div class="w-[50%]">{{ addressAlias.alias }}</div>
                <!-- <div>{{ addressAlias.addresses }}</div> -->
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </AppLayout>
</template>
