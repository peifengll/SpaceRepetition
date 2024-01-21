// // src\api\common.ts
// // 公共基础接口封装
// import request from '@/utils/request'
// import { DemoData, LoginResponse } from '@/api/types/common'
//
// export const demo = () => {
//     return request<DemoData>({
//         method: 'GET',
//         url: '/demo'
//     })
// }
//
// export const login = (data: {
//     account: string
//     pwd: string
//     imgcode: string
// }) => {
//     return request<LoginResponse>({
//         method: 'POST',
//         url: '/login',
//         data
//     })
// }
//
// // src\api\types\common.ts
// export interface DemoData {
//     title: string
//     date: number
// }
//
// export interface Menu {
//     id: number,
//     routePath: string
//     routeName: string
//     name: string
//     icon: string
//     children?: Menu[]
// }
//
// export interface UserInfo {
//     id: number
//     account: string
//     realName: string
// }
//
// export interface LoginResponse {
//     token: string
//     userInfo: UserInfo
//     menus: Menu[]
//     uniqueAuth: string[]
// }
//
