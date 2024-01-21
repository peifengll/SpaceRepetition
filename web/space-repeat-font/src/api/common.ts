// src\api\common.ts
// 公共基础接口封装
import request from '../utils/request'
import {LoginResponse} from './types/common.ts'
//  封装的泛型
interface ResponseData<T = any> {
    status: number
    msg: string
    data: T
}

export const demo = () => {
    return request.get<ResponseData<{}>>('/')
}

export const login = (data: {
    email:string
    password:string
}) => {
    return request<LoginResponse>({
        method:'POST',
        url:'v1/login',
        data:JSON.stringify(data),
    })
}





