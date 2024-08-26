import { defineStore } from 'pinia'
import axios from 'axios'
import path from 'path-browserify'
import { ref } from 'vue'
import { useErrorStore } from '@/stores/error.js'

const filesEndpoint = import.meta.env.VITE_SERVER_URL_V0 + '/fs'
const errStore = useErrorStore()

export const useFileStore = defineStore('files', () => {
    const files = ref([])
    const location = ref('/')
    const isLoading = ref(false)
    const filePath = ref('')

    const processData = (payload) => {
        if (payload.Items && Array.isArray(payload.Items)) {
            files.value = payload.Items
        } else {
            // todo error?
            console.log('No items found in the payload.')
        }
    }

    const isRoot = () => {
        return location.value === '/'
    }
    const getLocation = (dir) => {
        return path.normalize(path.join('/', location.value, dir))
    }

    const load = (dest) => {
        isLoading.value = true
        errStore.clear()
        location.value = path.normalize(dest)
        filePath.value = filesEndpoint + location.value
        files.value = []
        axios
            .get(path.join(filesEndpoint, dest))
            .then((res) => {
                if (res.status === 200) {
                    processData(res.data)
                } else {
                    errStore.set(false,'unexpected response code',"files_axios")
                }
            })
            .catch((err) => {
                switch (err.response.status) {
                    case 404:{
                        errStore.set(false,err.message)
                        break
                    }
                    default:{
                        errStore.set(true,err.message)
                    }
                }
            })
            .finally(() => {
                isLoading.value = false
            })
    }
    const deleteItem = (file) => {
        isLoading.value = true
        errStore.clear()
        axios
            .delete(path.join(filesEndpoint, getLocation(file)))
            .then((res) => {
                if (res.status === 200) {
                    // remove the file from the files list
                    files.value = files.value.filter((item) => item.Name !== file)
                } else {
                    errStore.set(false,'unexpected response code',"files_axios")
                }
            })
            .catch((err) => {
                switch (err.response.status) {
                    case 404:{
                        errStore.set(false,err.message)
                        break
                    }
                    default:{
                        errStore.set(true,err.message)
                    }
                }
            })
            .finally(() => {
                isLoading.value = false
            })
    }

    const createDir = (dirName) => {
        isLoading.value = true
        errStore.clear()
        axios
            .put(path.join(filesEndpoint, getLocation(dirName)))
            .then((res) => {
                if (res.status === 200) {
                    // add new dir to file list
                    files.value.push({
                        Name: dirName,
                        IsDir: true
                    })
                } else {
                    errStore.set(false,'unexpected response code',"files_axios")
                }
            })
            .catch((err) => {
                switch (err.response.status) {
                    case 404:{
                        errStore.set(false,err.message)
                        break
                    }
                    default:{
                        errStore.set(true,err.message)
                    }
                }
            })
            .finally(() => {
                isLoading.value = false
            })
    }

    return {
        files, // a list of files in the dir
        filePath, // full location path including api
        load, // trigger navigation
        deleteItem,
        createDir,
        isRoot, // check if the dir is the root
        getLocation, // return the new path after adding a directory
    }
})
