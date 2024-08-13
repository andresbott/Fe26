<script setup>
import Tree from 'primevue/tree'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import { ref } from 'vue'
import { useFileStore } from '@/stores/files.js'
const store = useFileStore()
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
const dialogVisible = ref(false);
const toDelFile = ref("")
function askConfirmation(path){
    toDelFile.value = path
    dialogVisible.value = true
}

function deleteDir (){
    store.deleteItem(toDelFile.value)
    toDelFile.value = ""
    dialogVisible.value = false
}

</script>

<template>
    <Dialog v-model:visible="dialogVisible" :modal="true" :closable="false" :draggable="false" header="Confirm Deletion" :style="{ width: '50rem' }"  >
        <span class="block mb-4">Are you sure you want to delete "{{ toDelFile}}"</span>
        <div class="flex justify-end gap-3">
            <Button type="button" label="Ok"  icon="pi pi-check" @click="deleteDir"></Button>
            <Button type="button" label="Cancel" icon="pi pi-times" severity="secondary" @click="dialogVisible = false"></Button>
        </div>
    </Dialog>
    <Tree
        :value="items.dirs"
        @nodeSelect="onNodeSelect"
        selectionMode="single"
        class="w-full md:w-30rem fe26-tree-list "
    >
        <template #default="slotProps">
            <span >{{ slotProps.node.label }}</span>
            <Button
                icon="pi pi-trash"
                severity="danger"
                text
                aria-label="Cancel"
                @click.stop="askConfirmation(slotProps.node.label)"
            />
        </template>
        <template #levelUp="slotProps">
            <span >{{ slotProps.node.label }}</span>
        </template>
        <template #createNew="slotProps">
            <span style="font-style: italic">{{ slotProps.node.label }}</span>
        </template>


    </Tree>
</template>
<style lang="scss">
.fe26-tree-list{
    .p-tree-node-content{
        position: relative;
    }

    .p-tree-node:hover .p-button{
        display: inline;
    }

    .p-tree-node .p-button{
        position: absolute;
        right: 1rem;
    }
    .p-button{
        display: none;
        padding: 0;
        border : 0;
        border-radius: 2px;
        background: transparent;
    }
    .p-button.p-button-secondary:not(:disabled):hover,
    .p-button.p-button-danger:not(:disabled):hover
    {
        border: 0;
        background: transparent;
    }
}
</style>
