import { defaultApiConfig } from "./api_config";
import { AddressesAliasesApiFactory } from "./clients/mailservermngr";

const addressesAliasesApi = AddressesAliasesApiFactory(defaultApiConfig)

export async function getAllAddressesAliases() {
    return await addressesAliasesApi.getAddressAliases()
}

export async function getAddressesAliasesByDomain(domain: string) {
    return await addressesAliasesApi.getAddressAliasesByDomain(domain)
}