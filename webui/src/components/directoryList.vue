<script setup>
import Tree from 'primevue/tree'
const items = defineProps({
    dirs: {
        default: []
    },
    select:{
        default: ""
    }
})
const onNodeSelect = (node) => {
    if (items.select && typeof items.select==="function") {
        items.select(node.label)
    }
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
