import { defineStore } from 'pinia'
import axios from 'axios'
import path from 'path-browserify'
import { ref } from 'vue'

const filesEndpoint = import.meta.env.VITE_SERVER_URL_V0 + '/fs'

export const useFileStore = defineStore('files', () => {
    const files = ref([])
    const location = ref('/')
    const isDataLoaded = ref(false)
    const isLoading = ref(false)

    const processData = (payload) => {
        if (payload.Items && Array.isArray(payload.Items)) {
            files.value = payload.Items
        } else {
            console.log('No items found in the payload.')
        }
    }
    const isRoot = () =>{
        if (location.value === "/"){
            return true
        }else {
            return false
        }
    }

    const goTo = (dest) => {
        isDataLoaded.value = true
        isLoading.value = true
        location.value = path.normalize(path.join(location.value,dest))

        axios
            .get(filesEndpoint +location.value)
            .then((res) => {
                if (res.status === 200) {
                    processData(res.data)
                } else {
                    console.log('err')
                    console.log(res)
                    // error?
                }
            })
            .catch((err) => {
                console.log(err)
            })
            .finally(() => {
                isLoading.value = false
            })
    }


    return {
        files,
        goTo,
        isRoot

    }
})

