import { authApiConfig, setApiAccessToken } from './api_config'
import { AuthApiFactory } from './clients/mailservermngr'

const authApi = AuthApiFactory(authApiConfig)

class MailServerMngrAuth {
  isLoggedIn() {
    const token = localStorage.getItem('access_token')
    if (!token)
       return false
    
    return true
  }
  login(username: string, password: string) {
    return new Promise<void>((resolve, reject) => {
      authApi
        .getToken({ username, password })
        .then((response) => {
          localStorage.setItem('access_token', response.data.accessToken)
          setApiAccessToken(response.data.accessToken)
          resolve()
        })
        .catch((reason) => {
          if (reason.response) {
            reject(reason.response.data.message)
          }
          reject(reason.message)
        })
    })
  }
  logout() {
    localStorage.setItem('access_token','')
  }
}

export const Authenticator = new MailServerMngrAuth()
