import { defaultApiConfig } from "./api_config";
import { DomainsAliasesApiFactory } from "./clients/mailservermngr";

const domainsAliasesApi = DomainsAliasesApiFactory(defaultApiConfig)

export async function getAllDomainsAliases() {
    return await domainsAliasesApi.getDomainsAliases()
}

export async function getDomainsAliasesByDomain(domain: string) {
    return await domainsAliasesApi.getDomainsAliasesByDomain(domain)
}