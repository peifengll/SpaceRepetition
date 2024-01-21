import {defineStore} from "pinia";

export const useCounterStore = defineStore('counter', {
    state: () => {
        return {
            count: 0,
            token: ""
        }
    },
    actions: {
        increment() {
            this.count++
        }
    }
})