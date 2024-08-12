import { defineStore } from 'pinia'
import axios from 'axios'
import path from 'path-browserify'
import { ref } from 'vue'

const filesEndpoint = import.meta.env.VITE_SERVER_URL_V0 + '/fs'

export const useFileStore = defineStore('files', () => {
    const files = ref([])
    const location = ref('/')
    const isLoading = ref(false)
    const isErr = ref(false)
    const errMessage = ref("")
    const filePath = ref("")

    const processData = (payload) => {
        if (payload.Items && Array.isArray(payload.Items)) {
            files.value = payload.Items
        } else {
            // todo error?
            console.log('No items found in the payload.')
        }
    }

    const isRoot = () =>{
        return location.value === "/";
    }
    const getLocation = (dir) =>{
        return path.normalize(path.join("/",location.value,dir))
    }

    const load = (dest) => {
        isLoading.value = true
        isErr.value =false
        errMessage.value=""

        location.value = path.normalize(dest)
        filePath.value = filesEndpoint +location.value

        axios
            .get(path.join(filesEndpoint, dest))
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
    const deleteFile = (file) =>{
        isLoading.value = true
        isErr.value =false
        errMessage.value=""
        axios
            .delete(path.join(filesEndpoint, getLocation(file)))
            .then((res) => {
                if (res.status === 200) {
                    // remove the file from the files list
                    files.value = files.value.filter(item => item.Name !== file);
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
        filePath, // full location path including api
        load, // trigger navigation
        deleteFile,
        isRoot, // check if the dir is the root
        getLocation, // return the new path after adding a directory
        isErr, // true if there was an error
        errMessage // the error message
    }
})

