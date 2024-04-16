// src\utils\request.ts
import axios from 'axios'
import {ElMessage} from 'element-plus'
// 在 plugins/element-plus.ts 引入了全部组件样式，这里不需额外引入
import {useStore} from '../stores'

export const SelfUrl:string= 'http://localhost:5173/'

// import {stringify} from "querystring";
// 创建 axios 实例
const request = axios.create({
    baseURL: 'http://peifeng.site:8003/'
    // baseURL: 'http://127.0.0.1:8003/'
})

// 请求拦截器
request.interceptors.request.use(function (config) {
    // 统一设置用户身份 token
    const lj = useStore()

    if (lj.token != "") {
        config.headers.Authorization = "Bearer " + lj.token
    } else {
        ElMessage.error("你还没有登录")
    }
    return config
}, function (error) {
    return Promise.reject(error)
})
//
// // 响应拦截器
// request.interceptors.response.use(function (response) {
//     // 统一处理接口响应错误，如 token 过期无效、服务端异常等
//     if (response.data.status && response.data.status !== 200) {
//         ElMessage.error(response.data.msg || '请求失败，请稍后重试')
//         return Promise.reject(response.data)
//     }
//     return response
// }, function (error) {
//     return Promise.reject(error)
// })

export default request

