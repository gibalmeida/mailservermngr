import { Configuration } from "./clients/mailservermngr";

const DEFAULT_BASE_PATH = 'http://localhost:8080'
export const defaultApiConfig = new Configuration()
defaultApiConfig.basePath = DEFAULT_BASE_PATH

export const authApiConfig = new Configuration()
authApiConfig.basePath = DEFAULT_BASE_PATH

export function setApiAccessToken(token: string) {
    defaultApiConfig.accessToken = token
}