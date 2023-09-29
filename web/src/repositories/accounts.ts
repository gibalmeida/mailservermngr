import { defaultApiConfig } from "./api_config";
import { AccountsApiFactory } from "./clients/mailservermngr";

const accountsApi = AccountsApiFactory(defaultApiConfig)

export async function getAllAccounts() {
    return await accountsApi.getAccounts()
}

export async function getAccountsByDomain(domain: string) {
    return await accountsApi.getAccountsByDomain(domain)
}