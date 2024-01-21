import { type PiniaPluginContext } from 'pinia';
import {toRaw} from 'vue';
type Options = { key?:string }
const __pinia_key__:string = 'pinia';
const setStorage = (key:string,value:any) => {
    localStorage.setItem(key,JSON.stringify(value))
}
const getStorage = (key:string) => {
    return localStorage.getItem(key) ? JSON.parse(localStorage.getItem(key) as string) : {}
}
const piniaPlugin = (options:Options) => {
    return (context:PiniaPluginContext) => {
        const {store} = context;
        const data = getStorage(`${options?.key ?? __pinia_key__}-${store.$id}`);
        store.$subscribe(()=>{ // 监听store变化
            setStorage(`${options?.key ?? __pinia_key__}-${store.$id}`,toRaw(store.$state))
        });
        console.log(store,'store');
        return {...data}
    }
}
export default piniaPlugin;