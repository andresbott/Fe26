import { defineStore } from 'pinia'
import axios from 'axios'
import path from 'path-browserify'
import { ref } from 'vue'


export const useErrorStore = defineStore('error', () => {
    const isErr = ref(false)
    const isCritical = ref(false)
    const msg = ref("")
    const origin = ref("")


    const clear = () => {
        isErr.value = false
    }

    const set = (critical = false,input ="", ori = "") =>{
        isErr.value = true
        isCritical.value = critical
        msg.value = input
        origin.value = ori
    }

    return {
        isErr,
        isCritical,
        msg,
        origin,
        clear,
        set
    }
})
