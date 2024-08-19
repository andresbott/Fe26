import { defineStore } from 'pinia'
import axios from 'axios'
import path from 'path-browserify'
import { ref } from 'vue'

const metaInfoEndpoint = import.meta.env.VITE_SERVER_URL_V0 + '/info'

export const userInfoStore = defineStore('info', () => {
    const version = ref()
    const buildTime = ref()
    const commit = ref()

    const isLoading = ref(false)
    const loadErr = ref("")

    const processData = (payload) => {
        console.log(payload)
        if (payload.meta.version ) {
            version.value = payload.meta.version
        }
        if (payload.meta.buildtime ) {
            buildTime.value = payload.meta.buildtime
        }
        if (payload.meta.commit ) {
            commit.value = payload.meta.commit
        }
    }

    const load = () => {
        isLoading.value = true
        loadErr.value = ''


        axios
            .get(metaInfoEndpoint)
            .then((res) => {
                if (res.status === 200) {
                    processData(res.data)
                } else {
                    console.log('err')
                    console.log(res)
                    loadErr.value= err.message
                    // error?
                }
            })
            .catch((err) => {
                loadErr.value= err.message

            })
            .finally(() => {
                isLoading.value = false
            })
    }

    return {
        load,
        version,
        buildTime,
        commit
    }
})
