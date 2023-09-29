import { defaultApiConfig } from "./api_config";
import { DomainsApiFactory } from "./clients/mailservermngr";

const domainsApi = DomainsApiFactory(defaultApiConfig)

export async function getAllDomains() {
    return await domainsApi.getDomains()
}