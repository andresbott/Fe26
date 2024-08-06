<script setup>
import Tree from 'primevue/tree'
import { useFileStore } from '@/stores/files.js'
const store = useFileStore()
const items = defineProps({
    files: {
        default: []
    },
})

console.log(items)

</script>

<template>
    <Tree
        :value="items.files"

        selectionMode="single"
        class="w-full md:w-30rem"
    ></Tree>
</template>
