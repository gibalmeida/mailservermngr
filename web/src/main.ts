import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import { setApiAccessToken } from './repositories/api_config'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { faBars, faClose, faRightFromBracket, faRightToBracket, faUserSecret, faPlus, faEdit, faTrash, faArrowLeft } from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(faBars, faClose, faRightFromBracket, faRightToBracket, faUserSecret, faPlus, faEdit, faTrash, faArrowLeft)

const app = createApp(App)

app.use(router)
app.use(Toast)

app.component('fa-icon', FontAwesomeIcon)

app.mount('#app')

setApiAccessToken(localStorage.getItem('access_token') ?? '')