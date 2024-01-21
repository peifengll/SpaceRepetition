import {createApp} from 'vue'
// import './style.css'
import router from './router'

import App from './App.vue'
import {createPinia} from 'pinia'
// import locale from "element-plus/es/locale/lang/zh-cn";
import elementPlus from './plugins/element-plus'
// import piniaPluginPersist from 'pinia-plugin-persist' //解决pinia数据刷新丢死问题
// import { QuillEditor } from '@vueup/vue-quill';
import VueUeditorWrap from 'vue-ueditor-wrap'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import piniaPlugin from './plugins/piniaPlugin'


const store = createPinia()
store.use(piniaPlugin({ // 第三步
    key: "pinia"
}))
createApp(App)
    // .component('QuillEditor', QuillEditor)
    .use(router)
    .use(store)
    .use(elementPlus)
    .use(VueUeditorWrap)
    // .use(piniaPluginPersist)
    .mount('#app')
