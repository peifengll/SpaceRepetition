// src\stores\index.ts
import {defineStore} from 'pinia'
// import {UserInfo} from '../api/types/common.ts'
// import {USER} from '../utils/constants.ts'
// import {getItem, setItem} from '../utils/storage.ts'
//
// type User = ({ token: string } & UserInfo) | null

export const useStore = defineStore('main', {
    state: () => ({
        count: 0,
        isCollapse: false,
        token:"" ,
    }),
    getters: {
        doubleCount(state) {
            return state.count * 2
        }
    },
    actions: {
        increment() {
            this.count++
        },

    }
});


export default useStore

