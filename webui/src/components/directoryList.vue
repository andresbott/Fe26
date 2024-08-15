<script setup>
import Tree from 'primevue/tree'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import { computed, ref } from 'vue'
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

const folders = computed(()=>{
    items.dirs.sort((a, b) => {
        if (a.label < b.label) return -1; // a comes before b
        if (a.label > b.label) return 1;  // b comes before a
        return 0; // a and b are equal
    });
    if (!store.isRoot()){
        items.dirs.unshift({
            key: 0,
            label: '..',
            selectable: true,
            icon: 'pi pi-fw pi-chevron-up ',
            type: "levelUp"
        })
    }
    items.dirs.push({
        key: 0,
        label: 'Add folder...',
        selectable: true,
        icon: 'pi pi-fw pi-plus',
        type: "createNew"
    })
    return items.dirs
})

const onNodeSelect = (node) => {
    if (node.type === 'createNew') {
        // ask the create dir dialog
        createDirDialog.value = true
    }else{
        if (items.select && typeof items.select==="function") {
            items.select(node.label)
        }
    }
}

const createDirDialog = ref(false);
const dirToCreate = ref("")
const createDir = (i)=>{
    store.createDir(dirToCreate.value)
    dirToCreate.value = ""
    createDirDialog.value = false
}

const deleteDialogVisible = ref(false);
const dirToDelete = ref("")
function askConfirmation(path){
    dirToDelete.value = path
    deleteDialogVisible.value = true
}
function deleteDir (){
    store.deleteItem(dirToDelete.value)
    dirToDelete.value = ""
    deleteDialogVisible.value = false
}
</script>

<template>
    <Dialog :visible="createDirDialog" :modal="true" :closable="false" :draggable="false" header="Create directory" :style="{ width: '50rem' }"  >
        <div v-focustrap >
            <span class="block mb-4">
                <InputText v-model="dirToCreate"  type="text" placeholder="Name" :style="{ width: '100%' }"  />
            </span>
            <div class="flex justify-end gap-3">
                <Button type="button" label="Ok"  icon="pi pi-check" @click="createDir"></Button>
                <Button type="button" label="Cancel" icon="pi pi-times" severity="secondary" @click="createDirDialog = false"></Button>
            </div>
        </div>

    </Dialog>

    <Dialog v-model:visible="deleteDialogVisible" :modal="true" :closable="false" :draggable="false" header="Confirm Deletion" :style="{ width: '50rem' }"  >
        <span class="block mb-4">Are you sure you want to delete "{{ dirToDelete }}"</span>
        <div class="flex justify-end gap-3">
            <Button type="button" label="Ok"  icon="pi pi-check" @click="deleteDir"></Button>
            <Button type="button" label="Cancel" icon="pi pi-times" severity="secondary" @click="deleteDialogVisible = false"></Button>
        </div>
    </Dialog>


    <Tree
        :value="folders"
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
            <span style="font-style: italic" >{{ slotProps.node.label }}</span>
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
