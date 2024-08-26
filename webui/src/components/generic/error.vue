<script setup>
import Dialog from 'primevue/dialog'
import { computed } from 'vue'
import { useErrorStore } from '@/stores/error.js'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'
import { watch } from 'vue';

const store = useErrorStore()
const toast = useToast()
const istCritical = computed(() => {
    return store.isErr && store.isCritical
})

// Watch the error property for changes
watch(
    () => store.isErr,
    (newError) => {
        if (newError && !store.isCritical){
                toast.add({
                    severity: 'error',
                    summary: `Error`,
                    detail: store.msg,
                    life: 3000,
                });
        }
    },
    { immediate: false } // If you want to trigger on initial load as well
);

</script>
<template>
    <Dialog
        v-model:visible="istCritical"
        header="Critica Error"
        :style="{ width: '25rem' }"
        :modal="true"
        :closable="false"
        :draggable="false"
    >
        <span class="text-surface-500 dark:text-surface-400 block mb-8"> {{ store.msg }}</span>
     </Dialog>
    <Toast />
</template>
