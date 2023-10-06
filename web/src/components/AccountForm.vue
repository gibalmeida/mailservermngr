<script setup lang="ts">
import PrimaryButton from '@/components/PrimaryButton.vue'
import DangerButton from '@/components/DangerButton.vue'
import CancelButton from '@/components/CancelButton.vue'
import InputText from './InputText.vue';
import type { Account } from '@/repositories/clients/mailservermngr';

interface Props {
    new?: boolean
    account: Account
    cancelButton?: Function 
}

const props = defineProps<Props>()
const account = props.account

function saveAccount() {
    // implement me
    alert('Save clicked')
}

function deleteAccount() {
    // implement me
    alert('Delete clicked')
}
</script>

<template>
    <form  @submit.prevent="saveAccount">
        <div>
            <h1 class="font-semibold">{{ props.new ? 'New' : 'Editing' }} Account</h1>
        </div>
        <InputText id="inputName" label="Name" placeholder="Type the email name" v-model="account.name" :required=true ></InputText>
        <InputText id="inputDomain" label="Domain" placeholder="Type the email domain" v-model=account.domain :required=true :disabled=true></InputText>

          <div class="flex mt-4 justify-center md:justify-end">
                <PrimaryButton type="submit" class="mr-2" >Save</PrimaryButton>
                <DangerButton type="button" class="mr-2" @click="deleteAccount" v-if="!props.new">Delete</DangerButton>
                <CancelButton type="button" @click="props.cancelButton">Cancel</CancelButton>
        </div>
    </form>
</template>