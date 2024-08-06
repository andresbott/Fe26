<script setup>
import Tree from 'primevue/tree'
import { useFileStore } from '@/stores/files.js'
const store = useFileStore()
const items = defineProps({
    dirs: {
        default: []
    },
})

const onNodeSelect = (node) => {
    store.goTo(node.label)
}
</script>

<template>
    <Tree
        :value="items.dirs"
        @nodeSelect="onNodeSelect"
        selectionMode="single"
        class="w-full md:w-30rem"
    ></Tree>
</template>
