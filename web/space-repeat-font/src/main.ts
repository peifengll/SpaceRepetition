import {createApp} from 'vue'
import router from './router'

import App from './App.vue'
import {createPinia} from 'pinia'
import elementPlus from './plugins/element-plus'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import piniaPlugin from './plugins/piniaPlugin'
import './styles/global.css'
import naive from 'naive-ui'

const store = createPinia()
store.use(piniaPlugin({ // 第三步
    key: "pinia"
}))
createApp(App)
    // .component('QuillEditor', QuillEditor)
    .use(router)
    .use(store)
    .use(elementPlus)
    .use(naive)
    .mount('#app')
