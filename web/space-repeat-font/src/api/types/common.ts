// src\api\types\common.ts
export interface DemoData {
    title: string
    date: number
}

export interface Menu {
    id: number,
    routePath: string
    routeName: string
    name: string
    icon: string
    children?: Menu[]
}

export interface UserInfo {
    id: number
    account: string
    realName: string
}

export interface LoginResponse {
    // any
    // data: JSON
    // userInfo: UserInfo
    // menus: Menu[]
    // uniqueAuth: string[]
}

