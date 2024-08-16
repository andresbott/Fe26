<script setup>
import FileUpload from 'primevue/fileupload'
import Toast from 'primevue/toast'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import { computed, ref } from 'vue'
import { useFileStore } from '@/stores/files.js'
const toast = useToast()
const store = useFileStore()

const onAdvancedUpload = ($event) => {
    isDragging.value = false
    const overlay = document.querySelector('.fe26-upload .p-fileupload-content')
    overlay.style.display = 'none' // Hide overlay
    toast.add({ severity: 'info', summary: 'Success', detail: 'File Uploaded', life: 3000 })

    store.load(store.getLocation(''))
}
// TODO upload progress...

const isDragging = ref(false)
document.addEventListener('dragover', (event) => {
    event.preventDefault()
    if (!isDragging.value) {
        isDragging.value = true
        const overlay = document.querySelector('.fe26-upload .p-fileupload-content')
        overlay.style.display = 'flex'
    }
})

document.addEventListener('dragleave', (event) => {
    // Check if we're actually leaving the entire window/document
    if (event.clientX === 0 && event.clientY === 0) {
        isDragging.value = false
        const overlay = document.querySelector('.fe26-upload .p-fileupload-content')
        overlay.style.display = 'none' // Hide overlay
    }
})
</script>
<template>
    <div class="">
        <Toast />
        <div class="fe26-upload">
            <FileUpload
                name="demo[]"
                :url="store.filePath"
                accept="*"
                :multiple="true"
                @upload="onAdvancedUpload($event)"
                :auto="true"
            >
                <template #header="{ chooseCallback, uploadCallback, clearCallback, files }">
                    <div class="flex gap-2">
                        <Button
                            @click="chooseCallback()"
                            icon="pi pi-upload"
                            rounded
                            severity="primary"
                        ></Button>
                    </div>
                </template>
                <template #content>
                    <div>
                        <div class="dots"></div>
                        <div class="msg">Drop files here to upload</div>
                    </div>
                </template>
                <template #empty>
                    <div>
                        <!--                    <div class="dots"> </div>-->
                        <!--                    <div class="msg">Drop files here to upload</div>-->
                    </div>
                </template>
            </FileUpload>
        </div>
    </div>
</template>

<style lang="scss">
.fe26-upload {
    .p-fileupload-advanced {
        border: 0;
    }
    .p-fileupload-header {
        background: none;
        > div {
            margin-left: auto;
        }
    }
    .p-fileupload-content {
        display: none;
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;

        background: rgba(0, 0, 0, 0.5); /* Semi-transparent background */
        align-items: center;
        justify-content: center;
        z-index: 9999; /* High z-index to overlay everything else */
        > * {
            pointer-events: none;
        }
        //pointer-events: none; /* Allows clicks to pass through */
        .dots {
            position: fixed;
            border: 2px dashed var(--accent-color);
            top: 5%;
            left: 5%;
            width: 90%;
            height: 90%;
        }
        .msg {
            color: var(--accent-color);
            font-size: 24px;
            text-align: center;
        }
    }
}
</style>
