
<style>
.fe26-file-list{
    --c-datatable-body-cell-padding:0.5rem;
    --c-datatable-row-color: #7f7f7f;
    margin: 1rem;
}
.fe26-file-list a{
    color: var( --c-datatable-row-color);
    text-decoration: none;
}
.fe26-file-list .p-button{
    opacity: 1;
    padding: 0.5rem 0.5rem;
    width: 2.2rem;
}
.fe26-file-list .p-datatable-header-cell,
.fe26-file-list .p-datatable-sortable-column:not(.p-datatable-column-sorted):hover
{
    border: 0;
    background: none;
}
.fe26-file-list .p-datatable-column-header-content{
    color: var( --c-datatable-row-color);
}
.fe26-file-list .p-datatable-column-sorted .p-datatable-column-header-content,
.fe26-file-list .p-datatable-column-sorted .p-datatable-sort-icon {
    color: #484848;
}


.fe26-file-list.p-datatable.p-datatable-striped .p-datatable-tbody > tr.p-row-odd{
    background: #f5f5f5;
    border: 0;
}
.fe26-file-list.p-datatable.p-datatable-striped .p-datatable-tbody > tr:hover{
    background: #e6e4e4;
}
.fe26-file-list.p-datatable.p-datatable-striped .p-datatable-tbody > tr{
    background: none;
}
.fe26-file-list .p-button,
.fe26-file-list .p-button.p-button-secondary:not(:disabled):hover,
.fe26-file-list .p-button.p-button-danger:not(:disabled):hover
{
    border : 0;
    border-radius: 2px;
    background: transparent;
}

.fe26-file-list .p-button.btn-default,.fe26-file-list .p-button.btn-default:hover{
    background: #dedede;
    color: #636972;
}

.fe26-file-list .p-button.btn-go-code,.fe26-file-list .p-button.btn-go-code:hover{
    background: #2db6da;
    color: #ffffff;
}
.fe26-file-list .p-button.btn-img,.fe26-file-list .p-button.btn-img:hover{
    background: #11b98c;
    color: #ffffff;
}
.fe26-file-list .p-button.btn-js-code,.fe26-file-list .p-button.btn-js-code:hover{
    background: #da892d;
    color: #ffffff;
}

.fe26-file-list .p-button.btn-pdf,.fe26-file-list .p-button.btn-pdf:hover{
    background: #c50000;
    color: #ffffff;
}

</style>

<script setup>
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";
import { humanFileSize } from '../lib/utils.js'
import { DateTime } from 'luxon'
import { ref } from "vue";
import { useFileStore } from '@/stores/files.js'

const store = useFileStore()
const toast = useToast();
const items = defineProps({
    files: {
        default: []
    },
})

function fmtTime(input ){
    return DateTime
        .fromSeconds(input)
        .toFormat('dd.MM.yy HH:mm')
}

function getIcon(name){
    const ext =  name.split('.').pop();
    switch (ext.toUpperCase()) {
        case 'JPG':
        case 'JPEG':
        case 'PNG':
        case 'GIF':
        case 'TIFF':
            return "pi pi-image"
        case 'GO':
        case 'JS':
        case 'PY':
            return "pi pi-code"
        case 'ZIP':
        case 'TAG':
        case 'GZ':
            return "pi pi-box"
        case 'PDF':
            return "pi pi-file-pdf"
        case 'MD':
            return "pi pi-align-justify"
        default:
            return "pi pi-file"
    }
}

function getColor(name){
    const ext =  name.split('.').pop();
    switch (ext.toUpperCase()) {
        case 'JPG':
        case 'JPEG':
        case 'PNG':
        case 'GIF':
        case 'TIFF':
            return "btn-img"
        case 'GO':
            return "btn-go-code"
        case 'JS':
            return "btn-js-code"
        // case 'PY':
        //     return "pi pi-code"
        // case 'ZIP':
        // case 'TAG':
        // case 'GZ':
        //     return "pi pi-box"
        case 'PDF':
            return "btn-pdf"
        default:
            return "btn-default"
    }
}



function copyClipboard(path){
    const URL = `${window.location.protocol}//${window.location.host}${path}`
    navigator.clipboard.writeText(URL);
    toast.add({ severity: 'info', summary: '', detail: 'Link copied to clipboard', life: 3000 });
}

const dialogVisible = ref(false);
const toDelFile = ref("")

function askConfirmation(path){
    toDelFile.value = path
    dialogVisible.value = true
}
// todo remove
console.log("\""+toDelFile.value+"\"")
function deleteFile (){
    store.deleteFile(toDelFile.value)
    toDelFile.value = ""
    dialogVisible.value = false
}
</script>

<template>
    <Toast class="only-message" />
    <Dialog v-model:visible="dialogVisible" :modal="true" :closable="false" :draggable="false" header="Confirm Deletion" :style="{ width: '50rem' }"  >
        <span class="block mb-4">Are you sure you want to delete "{{ toDelFile}}"</span>
        <div class="flex justify-end gap-3">
            <Button type="button" label="Ok"  icon="pi pi-check" @click="deleteFile"></Button>
            <Button type="button" label="Cancel" icon="pi pi-times" severity="secondary" @click="dialogVisible = false"></Button>
        </div>
    </Dialog>
    <DataTable class="fe26-file-list" :value="items.files" stripedRows >
        <Column header="" style="width: 1rem" >
            <template #body="slotProps">
                <Button
                    disabled
                    size="small"
                    :class="getColor(slotProps.data.name)"
                    :icon="getIcon(slotProps.data.name)"
                />
            </template>
        </Column>
        <Column field="name" sortable header="File">
            <template #body="slotProps">
                <a :href="slotProps.data.filePath" target="_blank">{{ slotProps.data.name}}</a>
            </template>
        </Column>
        <Column field="size" sortable header="Size" style="width: 7rem">
            <template #body="slotProps">
                {{ humanFileSize(slotProps.data.size )}}
            </template>
        </Column>
        <Column field="modTime" sortable header="Last Modified" style="width: 9.5rem">
            <template #body="slotProps">
                {{ fmtTime(slotProps.data.modTime )}}
            </template>
        </Column>
        <Column header="" style="width: 1rem" >
            <template #body="slotProps">
                <Button
                    icon="pi pi-link"
                    severity="secondary"
                    text
                    aria-label="Cancel"
                    @click="copyClipboard(slotProps.data.filePath)"
                />
            </template>
        </Column>
        <Column header=""  style="width: 1rem">
            <template #body="slotProps">
                <Button
                    icon="pi pi-trash"
                    severity="danger"
                    text
                    aria-label="Cancel"
                    @click="askConfirmation(slotProps.data.name)"
                />
            </template>
        </Column>
    </DataTable>
</template>
