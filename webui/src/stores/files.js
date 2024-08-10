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
    const isErr = ref(false)
    const errMessage = ref("")
    const filePath = ref("")

    const processData = (payload) => {
        if (payload.Items && Array.isArray(payload.Items)) {
            files.value = payload.Items
        } else {
            console.log('No items found in the payload.')
        }
    }
    const isRoot = () =>{
        return location.value === "/";
    }

    const goTo = (dest) => {
        isDataLoaded.value = true
        isLoading.value = true
        isErr.value =false
        errMessage.value=""
        location.value = path.normalize(path.join(location.value,dest))
        filePath.value = filesEndpoint +location.value

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
                isErr.value =true
                errMessage.value=err.message
            })
            .finally(() => {
                isLoading.value = false
            })
    }


    return {
        files, // a list of files in the dir
        filePath,
        goTo, // trigger navigation
        isRoot, // check if the dir is the root
        isErr, // true if there was an error
        errMessage // the error message

    }
})

