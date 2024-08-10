<script setup>

import { useFileStore } from '@/stores/files.js'
import Dialog from 'primevue/dialog'
import { computed } from 'vue'
const store = useFileStore()

const isHttpErr = computed(()=>{
    return store.isErr
})
const httpErrMsg = computed(()=>{
    return store.errMessage
})

</script>
<template>
    <Dialog v-model:visible="isHttpErr" header="Error" :style="{ width: '25rem' }" :modal="true" :closable="false" :draggable="false">
        <span class="text-surface-500 dark:text-surface-400 block mb-8">Something went wrong</span>
        <span>
            {{ httpErrMsg }}
        </span>
    </Dialog>
</template>